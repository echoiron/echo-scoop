@echo off
chcp 65001 1>nul 2>nul
net.exe session 1>NUL 2>NUL && (
    goto as_admin
) || (
    goto not_admin
)

:as_admin
cd /d "%~dp0"
reg add "HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Image File Execution Options\notepad.exe" /v "Debugger" /d "\"%~dp0Notepad3.exe\" /z" /f
cls
echo.
echo Replaced successfully by echoiron
echo.

goto end

:not_admin
echo Please run as administrator

:end
echo Press any key to end
pause 1>nul 2>nul