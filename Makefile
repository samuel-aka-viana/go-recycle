.PHONY: default run build test docs clean

APP_NAME := go-project
GO       := go
SWAG     := swag

default: run

run:
	@$(GO) run main.go

build:
	@$(GO) build -o $(APP_NAME) main.go

test:
	@$(GO) test ./...

docs:
	@$(SWAG) init

clean:
	@rm -f $(APP_NAME)
	@rm -rf docs
