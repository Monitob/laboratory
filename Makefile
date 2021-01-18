GO111MODULES=on
APP?=laboratory

.PHONY: build
## build: build the application
build:	clean
	@echo "Building..."
	@go build -o ${APP} main.go

.PHONY: run
## run: runs go run main.go
run:
	go run -race main.go

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@go clean

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

.PHONY: bench
## bench: runs go bench with default values
bench:
	go test -bench=. -benchmem ./...
.PHONY: fmt
# to be formatted
fmt:
	go fmt ./...

.PHONY: setup
## setup: setup go modules
setup:
	@go mod init \
		&& go mod tidy \
		&& go mod vendor

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

