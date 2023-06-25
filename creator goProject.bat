@echo off

:: Ввод данных:
set /p projectname="Название проекта:"
cd D:\Projects\golang
mkdir %projectname%
cd %projectname%
mkdir cmd
mkdir cmd/app
mkdir internal
mkdir pkg
mkdir pkg/service
mkdir pkg/handler
mkdir pkg/repository
mkdir utills
mkdir utills/logging
mkdir model


pause>nul
