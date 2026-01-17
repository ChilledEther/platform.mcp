<#
.SYNOPSIS
    Builds Docker image for the Platform MCP server.

.DESCRIPTION
    This script builds the Docker image for the Platform MCP implementation
    located at the root of the repository.

.PARAMETER ImageBaseName
    Base name for the Docker image. Defaults to "platform-mcp".

.EXAMPLE
    ./Invoke-DockerBuild.ps1
    # Builds the platform-mcp image
#>
[CmdletBinding()]
param(
    [string]$ImageBaseName = "platform-mcp"
)

$ErrorActionPreference = "Stop"

$root = Split-Path $PSScriptRoot -Parent
$context = $root

Write-Host " Docker Build - Platform MCP" -ForegroundColor Cyan
Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor DarkGray

$imageName = $ImageBaseName

Write-Host ""
Write-Host "📦 Building: $imageName" -ForegroundColor Yellow
Write-Host "   Context:  $context" -ForegroundColor DarkGray

# Verify Dockerfile exists
$dockerfile = Join-Path $context "build/package/platform-mcp/Dockerfile"
if (-not (Test-Path $dockerfile)) {
    Write-Host "   ❌ Dockerfile not found at $dockerfile" -ForegroundColor Red
    exit 1
}

# Build arguments
$buildArgs = @(
    "build",
    "--tag", "$($imageName):latest",
    "--file", "$dockerfile",
    $context
)

# Execute docker build
$startTime = Get-Date
docker @buildArgs
$duration = (Get-Date) - $startTime

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "✅ Build complete! ($([math]::Round($duration.TotalSeconds, 1))s)" -ForegroundColor Green
} else {
    Write-Host ""
    Write-Host "❌ Build failed with exit code $LASTEXITCODE" -ForegroundColor Red
    exit $LASTEXITCODE
}
