# Configuration variables for pathing to system variable,
# module directory, and go file. 
$WorkDir = "C:\Users\david\go\src\github.com\davidhintelmann\godb"
$SystemVariable = "go"
$File = "main.go"
$Binary = ($File.Replace(".go", ".exe"))
$BinaryPath = ($WorkDir + "\" + $Binary)
$ArgumentBuild = ("build -o " + $Binary + " " + $File)
$RunBinary = (".\" + $Binary)
$FileCSV = "sample.csv"

Write-Host "PowerShell running..."

# Effectively running 'go build -o main.exe main.go' in terminal/PowerShell
# but instead using PowerShell functions to measure execution duration.
$BuildTime = Measure-Command {
    Start-Process -FilePath $SystemVariable -WorkingDirectory $WorkDir -ArgumentList $ArgumentBuild -NoNewWindow -wait
}
Write-Host ("Time taken to build go binary -> " + $BuildTime)

# Now run the go binary
$RunBinaryTime = Measure-Command {
    Start-Process -FilePath $RunBinary -WorkingDirectory $WorkDir -NoNewWindow -wait
}
Write-Host ("Time taken to run go binary -> " + $RunBinaryTime)

# delete sample.csv file that is generated
# since the file is around 1.00 GB
# and delete go binary
Remove-Item ($WorkDir + "\sample.csv")
Remove-Item $BinaryPath
Write-Host ("Deleted " + $FileCSV)
Write-Host ("Deleted " + $Binary)