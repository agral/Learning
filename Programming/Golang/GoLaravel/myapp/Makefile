BINARY_NAME=celeritas

build:
	@echo "Building..."
	@go mod vendor
	@go build -o "tmp/${BINARY_NAME}"
	@echo "Successfully built."

run: build
	@echo "Running..."
	@./tmp/${BINARY_NAME} &
	@echo "Started."

clean:
	@echo "Cleaning..."
	@go clean
	@rm "tmp/${BINARY_NAME}"
	@echo "Cleaned."

test:
	@echo "Testing..."
	@go test ./...

start: run

stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped."

restart: stop start
