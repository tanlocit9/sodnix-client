@echo off
echo Running Go app...
cd src
go mod tidy
go run cmd/app/main.go
