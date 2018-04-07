@echo off

cd /d "%~dp0"

setlocal

REM ----- create build information

set GOBINPATH=

if not "%GOBIN%" == "" (
    set GOBINPATH=%GOBIN%
) else (
    if not "%GOPATH%" == "" (
        set GOBINPATH=%GOPATH%\bin
    ) else (
        set GOBINPATH=%USERPROFILE%\go\bin
    )
)

if not exist "%GOBINPATH%\ktn-build-info.exe" (
    echo ERROR - Failed to find konsorten build tool; run "go get -u github.com/konsorten/ktn-build-info" first
    exit /b 1
)

"%GOBINPATH%\ktn-build-info.exe"

if %ERRORLEVEL% NEQ 0 (
    exit /b 1
)

REM ----- rebuild the tool

go build

if %ERRORLEVEL% NEQ 0 (
    exit /b 1
)

REM ----- create help file

echo # Commandline help text: > HELP.md
echo ``` >> HELP.md
.\consul-versioning.exe -h >> HELP.md 2>NUL
echo ``` >> HELP.md

endlocal

exit /b 0
