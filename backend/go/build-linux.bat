@echo off
chcp 65001 >nul
echo Building Go API for Linux (amd64)...
cd /d %~dp0
if not exist bin mkdir bin
if exist bin\server.exe del /f /q bin\server.exe >nul 2>&1
if exist bin\server del /f /q bin\server >nul 2>&1

setlocal
set "GOOS=linux"
set "GOARCH=amd64"
set "CGO_ENABLED=0"
echo.
echo Environment variables:
echo   GOOS=%GOOS%
echo   GOARCH=%GOARCH%
echo   CGO_ENABLED=%CGO_ENABLED%
echo.
echo Building...
go build -o bin/server ./cmd/server
if errorlevel 1 (
    echo.
    echo Build failed! Please check errors above.
    endlocal
    pause
    exit /b 1
)
endlocal

if exist bin\server (
    for %%i in (bin\server) do (
        echo.
        echo Build successful!
        echo File: %%~fi
        echo Size: %%~zi bytes
    )
) else (
    echo.
    echo Build finished but bin\server not found!
    pause
    exit /b 1
)

echo.
echo Backend build done -^> backend\go\bin\server (linux/amd64)
echo Upload this file to your server and run with WATCH_* env vars.
pause
