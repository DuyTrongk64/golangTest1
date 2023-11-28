go mod init name_prj
mkdir -p cmd/web
mkdir views
mkdir models
-install touch: 
npm install touch-cli -g
touch Makefile
-install make:
-step1: install choco:
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
-step2: install make:
choco install make
-install chi routing:
go get -u github.com/go-chi/chi/v5
-install jet teamplate engine:
go get github.com/CloudyKit/jet/v6
