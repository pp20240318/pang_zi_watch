@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

set "BACKEND_PORT=8080"
set "FRONTEND_PORT=5173"
set "ADMIN_PORT=5180"
set "NPM_CMD="
set "NODEJS_DIR=C:\Program Files\nodejs"

if exist "%NODEJS_DIR%\node.exe" (
    echo %PATH% | find /I "%NODEJS_DIR%" >nul
    if errorlevel 1 set "PATH=%NODEJS_DIR%;%PATH%"
)

for /f "delims=" %%i in ('where npm 2^>nul') do (
    if not defined NPM_CMD set "NPM_CMD=%%i"
)
if not defined NPM_CMD if exist "C:\Program Files\nodejs\npm.cmd" set "NPM_CMD=C:\Program Files\nodejs\npm.cmd"

:menu
cls
echo ========================================
echo   胖子腕表 Watch
echo ========================================
echo   1. 启动 C 端商城     (port %FRONTEND_PORT%)
echo   2. 启动管理后台     (port %ADMIN_PORT%)
echo   3. 启动 Go API      (port %BACKEND_PORT%)
echo   4. 启动全部三个
echo   7. 发布 Go API     (Linux amd64)
echo   8. 构建 C 端商城
echo   9. 构建管理后台
echo   0. 退出
echo.
set "choice="
set /p "choice=请选择 (1-4, 7, 8-9, 0): "

if "%choice%"=="1" goto start_frontend
if "%choice%"=="2" goto start_admin
if "%choice%"=="3" goto start_backend
if "%choice%"=="4" goto start_all
if "%choice%"=="7" goto build_backend
if "%choice%"=="8" goto build_frontend
if "%choice%"=="9" goto build_admin
if "%choice%"=="0" goto end
goto menu

:kill_port
setlocal
set "PORT=%1"
for /f "tokens=5" %%a in ('netstat -ano ^| findstr :%PORT% ^| findstr LISTENING') do (
    set "PID=%%a"
    if not "!PID!"=="" (
        taskkill /PID !PID! /F >nul 2>&1
        timeout /t 1 /nobreak >nul
    )
)
endlocal
goto :eof

:ensure_npm
if defined NPM_CMD exit /b 0
echo npm 未找到，请安装 Node.js LTS 或重启终端。
timeout /t 2 /nobreak >nul
exit /b 1

:npm_install_if_needed
if exist node_modules exit /b 0
call "%NPM_CMD%" install
exit /b 0

:start_frontend
call :kill_port %FRONTEND_PORT%
call :ensure_npm
if errorlevel 1 goto menu
echo 正在启动 C 端商城...
pushd frontend
call :npm_install_if_needed
start "watch-frontend" cmd /k ""%NPM_CMD%" run dev"
popd
timeout /t 2 /nobreak >nul
goto menu

:start_admin
call :kill_port %ADMIN_PORT%
call :ensure_npm
if errorlevel 1 goto menu
echo 正在启动管理后台...
pushd backend\vue
call :npm_install_if_needed
start "watch-admin" cmd /k ""%NPM_CMD%" run dev"
popd
timeout /t 2 /nobreak >nul
goto menu

:start_backend
call :kill_port %BACKEND_PORT%
echo 正在启动 Go API...
pushd backend\go
start "watch-api" cmd /k "go run ./cmd/server"
popd
timeout /t 2 /nobreak >nul
goto menu

:start_all
call :kill_port %BACKEND_PORT%
call :kill_port %FRONTEND_PORT%
call :kill_port %ADMIN_PORT%
call :ensure_npm
if errorlevel 1 goto menu
echo 正在启动 Go API、C 端商城、管理后台...
pushd backend\go
start "watch-api" cmd /k "go run ./cmd/server"
popd
timeout /t 2 /nobreak >nul
pushd frontend
call :npm_install_if_needed
start "watch-frontend" cmd /k ""%NPM_CMD%" run dev"
popd
timeout /t 1 /nobreak >nul
pushd backend\vue
call :npm_install_if_needed
start "watch-admin" cmd /k ""%NPM_CMD%" run dev"
popd
timeout /t 2 /nobreak >nul
goto menu

:build_backend
echo 正在发布 Go API（Linux amd64）...
pushd backend\go
start "" "build-linux.bat"
popd
timeout /t 1 /nobreak >nul
goto menu

:build_frontend
call :ensure_npm
if errorlevel 1 goto menu
echo 正在构建 C 端商城...
pushd frontend
call :npm_install_if_needed
call "%NPM_CMD%" run build
popd
echo 构建完成 -^> frontend\dist
timeout /t 2 /nobreak >nul
goto menu

:build_admin
call :ensure_npm
if errorlevel 1 goto menu
echo 正在构建管理后台...
pushd backend\vue
call :npm_install_if_needed
call "%NPM_CMD%" run build
popd
echo 构建完成 -^> backend\vue\dist
timeout /t 2 /nobreak >nul
goto menu

:end
echo 再见。
pause
exit /b
