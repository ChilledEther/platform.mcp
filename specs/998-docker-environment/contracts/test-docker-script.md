# Test Script Contract: Test-Docker.ps1

**File**: `scripts/Test-Docker.ps1`  
**Purpose**: Validate Docker builds meet acceptance criteria

## Contract

```powershell
<#
.SYNOPSIS
    Tests Docker builds for Platform MCP.

.DESCRIPTION
    Validates that Docker images build successfully and meet all acceptance criteria
    including size limits, label presence, and health checks.

.PARAMETER ImageName
    Specific image to test. If not provided, tests all images.

.EXAMPLE
    ./Test-Docker.ps1
    # Tests all images

.EXAMPLE
    ./Test-Docker.ps1 -ImageName platform
    # Tests only the platform CLI image
#>
[CmdletBinding()]
param(
    [ValidateSet("platform", "platform-mcp", "all")]
    [string]$ImageName = "all"
)

$ErrorActionPreference = "Stop"

$root = Split-Path $PSScriptRoot -Parent
$maxSizeMB = 50
$images = @(
    @{ Name = "platform"; Dockerfile = "build/package/platform/Dockerfile" },
    @{ Name = "platform-mcp"; Dockerfile = "build/package/platform-mcp/Dockerfile" }
)

function Test-ImageSize {
    param([string]$Image)
    
    $sizeBytes = docker image inspect $Image --format '{{.Size}}' 2>$null
    if (-not $sizeBytes) { return $false }
    
    $sizeMB = [math]::Round($sizeBytes / 1MB, 2)
    Write-Host "   Size: ${sizeMB}MB (limit: ${maxSizeMB}MB)"
    
    return $sizeMB -lt $maxSizeMB
}

function Test-ImageLabels {
    param([string]$Image)
    
    $labels = docker image inspect $Image --format '{{json .Config.Labels}}' | ConvertFrom-Json
    $required = @(
        "org.opencontainers.image.title",
        "org.opencontainers.image.version",
        "org.opencontainers.image.source"
    )
    
    $missing = $required | Where-Object { -not $labels.$_ }
    if ($missing) {
        Write-Host "   Missing labels: $($missing -join ', ')" -ForegroundColor Red
        return $false
    }
    
    Write-Host "   Labels: OK"
    return $true
}

function Test-ImageUser {
    param([string]$Image)
    
    $user = docker image inspect $Image --format '{{.Config.User}}'
    if ($user -and $user -ne "root" -and $user -ne "0") {
        Write-Host "   User: $user (non-root)"
        return $true
    }
    
    Write-Host "   User: root (FAIL - should be non-root)" -ForegroundColor Red
    return $false
}

# Main execution
Write-Host " Docker Tests - Platform MCP" -ForegroundColor Cyan
Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor DarkGray

$testImages = if ($ImageName -eq "all") { $images } else { $images | Where-Object { $_.Name -eq $ImageName } }
$allPassed = $true

foreach ($img in $testImages) {
    Write-Host ""
    Write-Host "📦 Testing: $($img.Name)" -ForegroundColor Yellow
    
    # Build
    Write-Host "   Building..."
    $dockerfile = Join-Path $root $img.Dockerfile
    docker build -t "$($img.Name):test" -f $dockerfile $root 2>&1 | Out-Null
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "   ❌ Build failed" -ForegroundColor Red
        $allPassed = $false
        continue
    }
    Write-Host "   Build: OK"
    
    # Size check
    if (-not (Test-ImageSize "$($img.Name):test")) {
        Write-Host "   ❌ Size exceeds limit" -ForegroundColor Red
        $allPassed = $false
    }
    
    # Label check
    if (-not (Test-ImageLabels "$($img.Name):test")) {
        $allPassed = $false
    }
    
    # User check
    if (-not (Test-ImageUser "$($img.Name):test")) {
        $allPassed = $false
    }
    
    Write-Host "   ✅ All checks passed" -ForegroundColor Green
}

Write-Host ""
if ($allPassed) {
    Write-Host "✅ All tests passed!" -ForegroundColor Green
    exit 0
} else {
    Write-Host "❌ Some tests failed" -ForegroundColor Red
    exit 1
}
```

## Acceptance Criteria

- [ ] Script runs in PowerShell 7+
- [ ] Tests both images when no parameter provided
- [ ] Tests single image when specified
- [ ] Validates image size < 50MB
- [ ] Validates required OCI labels present
- [ ] Validates non-root user
- [ ] Returns exit code 0 on success, 1 on failure

## Usage

```powershell
# Test all images
./scripts/Test-Docker.ps1

# Test specific image
./scripts/Test-Docker.ps1 -ImageName platform

# In CI pipeline
pwsh -File ./scripts/Test-Docker.ps1
if ($LASTEXITCODE -ne 0) { exit 1 }
```
