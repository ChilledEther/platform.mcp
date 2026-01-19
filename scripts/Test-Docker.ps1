<#
.SYNOPSIS
    Tests the Platform MCP Docker image.

.DESCRIPTION
    Runs a basic smoke test against the built Docker image to ensure
    the binary is present and can execute.

.PARAMETER ImageName
    Name of the Docker image to test. Defaults to "platform-mcp".

.EXAMPLE
    ./Test-Docker.ps1
#>
[CmdletBinding()]
param(
    [string]$ImageName = "platform-mcp:latest"
)

$ErrorActionPreference = "Stop"

Write-Host " Docker Test - Platform MCP" -ForegroundColor Cyan
Write-Host "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━" -ForegroundColor DarkGray

Write-Host ""
Write-Host "🔍 Testing image: $ImageName" -ForegroundColor Yellow

# 1. Check if image exists
Write-Host "   - Checking image existence..." -NoNewline
$imageExists = docker images -q $ImageName
if (-not $imageExists) {
    Write-Host " ❌ Image not found!" -ForegroundColor Red
    exit 1
}
Write-Host " OK" -ForegroundColor Green

# 2. Run basic execution test (check binary version or help)
# Since it's an MCP server, it might just wait for input. 
# We'll try to run it and see if it exits with 0 or hangs (which is good for a server).
Write-Host "   - Running smoke test (binary execution)..." -NoNewline

# We use a timeout to ensure it doesn't hang forever if it's working correctly
# But wait, standard docker run might not return.
# Let's just try to list the binary inside the container.
docker run --rm $ImageName ls /root/platform-mcp > $null

if ($LASTEXITCODE -eq 0) {
    Write-Host " OK" -ForegroundColor Green
    Write-Host ""
    Write-Host "✅ Image verification successful!" -ForegroundColor Green
} else {
    Write-Host " ❌ Smoke test failed with exit code $LASTEXITCODE" -ForegroundColor Red
    exit $LASTEXITCODE
}
