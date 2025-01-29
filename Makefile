run:
	@echo "Running..."
	@go mod tidy
	go run .
	@echo "Done!"

build:
	@echo "Building..."
	@go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bin/gith -ldflags="-s -w" .
	GOOS=windows GOARCH=amd64 go build -o bin/gith.exe -ldflags="-s -w" .
	GOOS=darwin GOARCH=amd64 go build -o bin/gith.darwin -ldflags="-s -w" .
	@echo "Done!"

clean:
	@echo "Cleaning..."
	rm -rf bin/
	@echo "Done!"

test:
	@echo "Testing..."
	@go mod tidy
	go test -v ./...
	@echo "Done!"

# @reload:
# 	@echo "Reloading..."
# 	@echo "Done!"

@default:
	@echo "Usage: make [run|build|clean|test]"

.PHONY: run build clean test