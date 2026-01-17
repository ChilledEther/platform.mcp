<#
.SYNOPSIS
    Builds Docker images for both TypeScript and Go MCP server implementations.

.DESCRIPTION
    This script builds Docker images for the TypeScript and Go implementations
    located in the implementations/ directory. By default, it builds both images
    in parallel. You can specify a single implementation using the -Implementation parameter.

.PARAMETER ImageBaseName
    Base name for the Docker images. Defaults to "mcp-github-agentic".
    Final image names will be "{ImageBaseName}-ts" and "{ImageBaseName}-go".

.PARAMETER Implementation
    Which implementation to build. Valid values: "all", "typescript", "go".
    Defaults to "all" (builds both).

.EXAMPLE
    ./Invoke-DockerBuild.ps1
    # Builds both TypeScript and Go images

.EXAMPLE
    ./Invoke-DockerBuild.ps1 -Implementation typescript
    # Builds only the TypeScript image

.EXAMPLE
    ./Invoke-DockerBuild.ps1 -ImageBaseName "my-mcp-server"
    # Builds images named "my-mcp-server-ts" and "my-mcp-server-go"
#>
[CmdletBinding()]
param(
    [string]$ImageBaseName = "mcp-github-agentic",

    [ValidateSet("all", "typescript", "go")]
    [string]$Implementation = "all"
)

$ErrorActionPreference = "Stop"

$root = Split-Path $PSScriptRoot -Parent

# Define implementations configuration
$implementations = @(
    @{
        Name    = "typescript"
        Suffix  = "-ts"
        Context = Join-Path $root "implementations/typescript"
    },
    @{
        Name    = "go"
        Suffix  = "-go"
        Context = Join-Path $root "implementations/go"
    }
)

# Filter based on parameter
if ($Implementation -ne "all") {
    $implementations = $implementations | Where-Object { $_.Name -eq $Implementation }
}

Write-Host "🐳 Docker Build - MCP GitHub Agentic" -ForegroundColor Cyan
Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor DarkGray

$results = @()

foreach ($impl in $implementations) {
    $imageName = "$ImageBaseName$($impl.Suffix)"
    $context = $impl.Context

    Write-Host ""
    Write-Host "📦 Building: $imageName" -ForegroundColor Yellow
    Write-Host "   Context:  $context" -ForegroundColor DarkGray

    # Verify Dockerfile exists
    $dockerfile = Join-Path $context "Dockerfile"
    if (-not (Test-Path $dockerfile)) {
        Write-Host "   ❌ Dockerfile not found at $dockerfile" -ForegroundColor Red
        $results += @{ Name = $imageName; Success = $false; Error = "Dockerfile not found" }
        continue
    }

    # Build arguments
    $buildArgs = @(
        "build",
        "--tag", "$($imageName):latest",
        $context
    )

    # Execute docker build
    $startTime = Get-Date
    docker @buildArgs
    $duration = (Get-Date) - $startTime

    if ($LASTEXITCODE -eq 0) {
        Write-Host "   ✅ Build complete! ($([math]::Round($duration.TotalSeconds, 1))s)" -ForegroundColor Green
        $results += @{ Name = $imageName; Success = $true; Duration = $duration }
    } else {
        Write-Host "   ❌ Build failed with exit code $LASTEXITCODE" -ForegroundColor Red
        $results += @{ Name = $imageName; Success = $false; Error = "Exit code $LASTEXITCODE" }
    }
}

# Summary
Write-Host ""
Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor DarkGray
Write-Host "📊 Build Summary" -ForegroundColor Cyan

$successful = ($results | Where-Object { $_.Success }).Count
$failed = ($results | Where-Object { -not $_.Success }).Count

foreach ($result in $results) {
    if ($result.Success) {
        Write-Host "   ✅ $($result.Name) ($([math]::Round($result.Duration.TotalSeconds, 1))s)" -ForegroundColor Green
    } else {
        Write-Host "   ❌ $($result.Name) - $($result.Error)" -ForegroundColor Red
    }
}

Write-Host ""
if ($failed -gt 0) {
    throw "❌ $failed of $($results.Count) builds failed"
} else {
    Write-Host "✅ All $successful builds completed successfully!" -ForegroundColor Green
}
