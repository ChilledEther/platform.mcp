[cmdletBinding()]
param()

$ErrorActionPreference = "Stop"

# 1. Base Metadata
# Detect repo name and owner from git
$owner = "chilledether"
$repoName = "platform.mcp"
$description = "A powerful Model Context Protocol (MCP) platform built with Go."

try {
    $remoteUrl = (git remote get-url origin).Trim()
    if ($remoteUrl -match 'github\.com[:/]([^/]+)/([^/]+)') {
        $owner = $Matches[1].ToLower()
        $repoName = $Matches[2] -replace '\.git$', ''
    }
} catch {
    Write-Host "⚠️ Git remote not found, using defaults." -ForegroundColor Gray
}

# Keep the ID constant so it updates the same entry
$internalId = $repoName
$dateAdded = Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ"

$upstream = "https://github.com/$owner/$repoName"

# Determine source (current branch/commit)
$branch = "main"
try { $branch = (git rev-parse --abbrev-ref HEAD).Trim() } catch { }
$source = "$upstream/tree/$branch"

# Icon from GitHub
$icon = "https://github.com/$owner.png"

Write-Host "🧬 Processing catalog for: $internalId (Owner: $owner)" -ForegroundColor Cyan

# 3. Path Setup (Handle WSL2 -> Windows Docker Path)
$customCatalogName = "custom-mcps"
$catalogPath = Join-Path $HOME ".docker/mcp/catalogs"
if ($env:WSL_DISTRO_NAME) {
    try {
        $winHomeRaw = powershell.exe -NoProfile -Command "Write-Host -NoNewline `$env:USERPROFILE"
        $catalogPath = Join-Path (powershell.exe -NoProfile -Command "wslpath '$winHomeRaw'").Trim() ".docker/mcp/catalogs"
        $catalogPathWindows = "$winHomeRaw\.docker\mcp\catalogs"
        Write-Host "🪟 WSL Detected: Targeting Windows Docker Path..." -ForegroundColor Gray
    } catch { }
}
$catalogId = $internalId -replace '\.', '-'
$catalogFile = Join-Path $catalogPath "$catalogId.yaml"

if (!(Test-Path $catalogFile)) {
    # Ensure directory exists
    if (!(Test-Path $catalogPath)) { New-Item -ItemType Directory -Force -Path $catalogPath | Out-Null }
    Write-Host "✨ Creating catalog file: $catalogId" -ForegroundColor Yellow
}

# 4. Construct Server Entry
$catalogEntry = [ordered]@{
    name = $catalogId
    displayName = "$owner/$internalId"
    registry = [ordered]@{
        $catalogId = [ordered]@{
            description = $description
            title       = ($internalId -replace '-', ' ').ToUpper()
            type        = "server"
            dateAdded   = $dateAdded
            image       = "ghcr.io/$owner/${internalId}:latest"
            ref         = ""
            source      = $source
            upstream    = $upstream
            icon        = $icon
        }
    }
}

# 5. Merge Metadata using yq
$jsonContent = $catalogEntry | ConvertTo-Json -Depth 10

try {
    Write-Host "📝 Updating $catalogFile with container-style metadata..." -ForegroundColor Cyan

    $jsonContent | yq --input-format json --output-format yaml | Out-File -Path $catalogFile -Force

    $importPath = if ($env:WSL_DISTRO_NAME) { $catalogFileWindows } else { $catalogFile }
    
    docker mcp catalog import $importPath
    
    docker mcp catalog add $customCatalogName $catalogId $importPath --force

    Write-Host "✅ Catalog updated successfully!" -ForegroundColor Green
} catch {
    Write-Error "❌ Error updating catalog: $_"
} finally {
    docker mcp catalog update
}
