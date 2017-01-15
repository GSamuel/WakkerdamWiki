@echo off

goto Restart

:Start
git fetch
for /f %%i in ('git rev-parse HEAD') do set VAR1=%%i
for /f %%i in ('git rev-parse @{u}') do set VAR2=%%i

if %VAR1% NEQ %VAR2% goto Restart

:End
timeout /t 60
goto Start

:Restart
git pull
taskkill /f /im weerwolvenwiki.exe
go build src/main/main.go
copy main.exe src\main\weerwolvenwiki.exe
del main.exe
cd src\
start main\weerwolvenwiki.exe
cd..
goto End

pause