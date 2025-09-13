@echo off
echo [1/2] Cleaning up Go modules...
cd src
go mod tidy

if %errorlevel% neq 0 (
    echo Failed to tidy go.mod.
    exit /b %errorlevel%
)

echo [2/2] Generating Swagger docs...
swag init -g cmd/app/main.go -o docs --parseDependency --parseInternal

if %errorlevel% neq 0 (
    echo Failed to generate Swagger docs.
    exit /b %errorlevel%
)