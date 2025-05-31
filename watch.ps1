$watcher = New-Object System.IO.FileSystemWatcher
$watcher.Path = "."
$watcher.Filter = "*.go"
$watcher.IncludeSubdirectories = $true
$watcher.EnableRaisingEvents = $true

$action = {
    $path = $Event.SourceEventArgs.FullPath
    $changeType = $Event.SourceEventArgs.ChangeType
    $timeStamp = (Get-Date).ToString("yyyy-MM-dd HH:mm:ss")
    Write-Host "[$timeStamp] Detected $changeType in $path"
    
    # Kill existing process if any
    Get-Process -Name "main" -ErrorAction SilentlyContinue | Stop-Process -Force
    
    # Build and run
    Write-Host "Building and starting server..."
    Start-Process -NoNewWindow -FilePath "go" -ArgumentList "run", "main.go"
}

Register-ObjectEvent $watcher "Created" -Action $action
Register-ObjectEvent $watcher "Changed" -Action $action
Register-ObjectEvent $watcher "Deleted" -Action $action

Write-Host "Watching for changes in .go files..."
Write-Host "Press Ctrl+C to stop watching"

# Start the server initially
Start-Process -NoNewWindow -FilePath "go" -ArgumentList "run", "main.go"

# Keep the script running
while ($true) { Start-Sleep 1 } 