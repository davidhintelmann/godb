# Configuration variables for pathing to system variable,
# module directory, and go file. 
$WorkDir = "C:\Users\david\go\src\github.com\davidhintelmann\godb"
$SystemVariable = "go"
$File = "main.go"
$ArgumentRun = ("run " + $File)
$FileCSV = "sample.csv"

Write-Host "PowerShell running..."

# Effectively running 'go run main.go' in terminal/PowerShell
# but instead using PowerShell functions to measure execution duration.
$GoRunTime = Measure-Command {
    Start-Process -FilePath $SystemVariable -WorkingDirectory $WorkDir -ArgumentList $ArgumentRun -NoNewWindow -wait
}
Write-Host ("Time taken to run go -> " + $GoRunTime)

# delete sample.csv file that is generated
# since the file is around 1.00 GB
Remove-Item ($WorkDir + "\" + $FileCSV)
Write-Host ("Deleted " + $FileCSV)