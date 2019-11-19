OS ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
APP_ROOT ?= $(shell 'pwd')

include $(APP_ROOT)/Makefile.helpers

## Run sanskriti-vm
run:
	@go run main.go
