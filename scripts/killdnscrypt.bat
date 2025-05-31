@echo off
color 0A
title Kill dnscrypt-proxy.exe (admin required)

:: Check for admin rights
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo [!] Requesting administrator privileges...
    powershell -Command "Start-Process '%~f0' -Verb RunAs"
    exit /b
)

:: Header
echo -----------------------------
echo   DNSCrypt Killer Script
echo -----------------------------
echo.

:: Kill dnscrypt-proxy.exe process (forcefully)
echo [*] Attempting to kill dnscrypt-proxy.exe...
taskkill /IM dnscrypt-proxy.exe /F >nul 2>&1

:: Check result
if %errorlevel% == 0 (
    echo [OK] dnscrypt-proxy.exe terminated successfully.
) else (
    echo [INFO] dnscrypt-proxy.exe is not running or could not be killed.
)

echo.
pause
