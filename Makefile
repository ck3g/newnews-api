BINARY_NAME=newnewsapi

build:
	@echo "Building NewNews API..."
	@go build -o tmp/${BINARY_NAME}

run: build
	@echo "Starting NewNews API..."
	@./tmp/${BINARY_NAME} &

stop:
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "NewNews API stopped!"
