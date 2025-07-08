$COGNITOOLS_HOME = "$env:USERPROFILE\.cognitools"

function Install-Cognitools {
    $cognitoolsBinary = Join-Path $COGNITOOLS_HOME "bin\cognitools.exe"

    if (Test-Path $cognitoolsBinary) {
        Write-Host "Cognitools binary already exists at $cognitoolsBinary"
        return
    }

    # Detect architecture
    $arch = if ([System.Environment]::Is64BitOperatingSystem) {
        if ($env:PROCESSOR_ARCHITECTURE -match "ARM64") { "arm64" } else { "amd64" }
    } else {
        Write-Host "Unsupported architecture"
        return
    }

    $os = "windows"
    Write-Host "Installing cognitools for $os $arch"

    $binDir = Join-Path $COGNITOOLS_HOME "bin"
    New-Item -ItemType Directory -Force -Path $binDir | Out-Null

    $zipFile = Join-Path $COGNITOOLS_HOME "cognitools_${os}_${arch}.zip"
    $downloadUrl = "https://github.com/gdegiorgio/cognitools/releases/latest/download/cognitools_${os}_${arch}.zip"

    Invoke-WebRequest -Uri $downloadUrl -OutFile $zipFile
    Expand-Archive -Path $zipFile -DestinationPath $COGNITOOLS_HOME -Force
    Remove-Item $zipFile

    $extractedPath = Get-ChildItem -Path $COGNITOOLS_HOME -Recurse -Filter "cognitools.exe" | Select-Object -First 1

    if (-not $extractedPath) {
        Write-Host "Failed to find cognitools.exe after extraction."
        return
    }

    Move-Item -Path $extractedPath.FullName -Destination $cognitoolsBinary -Force

    # Add to PATH if not already there
    $envPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
    if ($envPath -notlike "*$($binDir)*") {
        [System.Environment]::SetEnvironmentVariable("COGNITOOLS_HOME", $COGNITOOLS_HOME, "User")
        [System.Environment]::SetEnvironmentVariable("Path", "$envPath;$binDir", "User")
        Write-Host "Added $binDir to PATH"
    }

    Write-Host "Cognitools installed successfully. Restart your terminal to use it now."
}

Install-Cognitools
