.PHONY: install
## install: installs the check-gen to go/bin
install:
	go install -i .

PHONY: build
## build: builds the check-gen
build:
	go build -o check-gen .

.PHONY: test
## test: tests generated examples
test:
	go test ./examples/...

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
