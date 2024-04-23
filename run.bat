@echo off
cls
chcp 866

echo Start process ...
rem set CGO_CFLAGS="-g -O2 -Wno-return-local-addr"
rem set CGO_ENABLED=0

rem go env
rem go mod init wrk
rem go version
set Keyprogram=Keyprogram-02020
set User=AdministartorSuperRight
REM SET GOCACHE=off 


:: 1 - показ 0 - нет
set Devoption=1

:: Показать запросы в строке
set Mode=dev2

:: Показ в отчетах обработку
set showlog=N

echo Go mod ...
go mod tidy

set FLAGS="-w -s"
set NAMEAPP=gogems.exe

::Ne version GO
go version
rem go install golang.org/dl/go1.22.1@latest


rem Все для проверки кода
Rem https://sparkbox.com/foundry/go_vet_gofmt_golint_to_code_check_in_Go
rem https://golangci-lint.run/usage/install/


rem go clean -cache
rem go clean -fuzzcache


rem LINTER 
rem Check code
rem go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

rem Запуск и проверка
rem golangci-lint run  

rem Запуск и проверка с опциями
rem golangci-lint run --disable-all -E errcheck

rem Check logical code
rem go vet 

rem В два раза больше размер файла
rem go build -ldflags "-linkmode external -extldflags -static" -o kpinfo.exe

cls
echo Compile...
rem go build -pgo=auto -ldflags %FLAGS% -o %NAMEAPP%

rem -x, которая печатает все, что Go делает для создания исполняемого файла.
rem -n, который печатает шаги, которые он обычно выполняет, без фактического создания двоичного файла.
rem -work помимо запуска программы печатает путь к рабочему каталогу,
rem -gcflags="-S" показывает инструкции по сборке.
rem go build -ldflags="-s -w" - уменьшить файл
rem -race - гонку включить обнаружение гонки данных.
rem -v выводить имена пакетов по мере их компиляции.

go build  -pgo=auto -ldflags %FLAGS% -o %NAMEAPP%
rem go build  & strip



rem %NAMEAPP% -port 8000
%NAMEAPP% -port 7004


rem Developer mode
rem kpinfo.exe --action=check

pause