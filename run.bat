@echo off
cls
chcp 866
go mod tidy

set FLAGS="-w -s"
set NAMEAPP=gogems.exe

rem Check logical code
rem go vet 

rem В два раза больше размер файла
rem go build -ldflags "-linkmode external -extldflags -static" -o gogems.exe

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

pause