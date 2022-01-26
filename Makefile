BINARY_NAME=newnewsapi

build:
	@echo "Building NewNews API..."
	@go build -o tmp/${BINARY_NAME}

start: build
	@echo "Starting NewNews API..."
	@./tmp/${BINARY_NAME} &

stop:
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "NewNews API stopped!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

db_connect:
	@psql -U postgres -W --port=5439 --host=localhost --dbname=newnews_development

migration:
	@migrate create -ext sql -dir db/migrations -seq $(name)

migrate:
	@migrate -database $(database_url) -path db/migrations up

rollback:
	@migrate -database $(database_url) -path db/migrations down
