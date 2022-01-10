include .env

COMPOSE_FILE?=docker-compose.dev.yml

migrate.up:
	echo "Starting migrate up"
	migrate -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path migrations up 1

migrate.down:
	echo "Starting migrate down"
	migrate -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path migrations down 1

dev:
	echo "Starting docker environment"
	docker-compose -f ${COMPOSE_FILE} up --build

stop:
	echo "Stopping docker environment"
	docker-compose -f ${COMPOSE_FILE} down

swag:
	echo "Starting swagger generating"
	swag init -g router/router.go

lint:
	echo "Starting linters"
	golangci-lint run ./...

test:
	echo "Starting tests"
	go test -cover ./...

# todo move repeated command to make command

docker.migrate.up:
	docker-compose -f ${COMPOSE_FILE}  run --rm --entrypoint sh --no-deps app -c "make migrate.up"

docker.migrate.down:
	docker-compose -f ${COMPOSE_FILE}  run --rm --entrypoint sh --no-deps app -c "make migrate.down"

docker.swag:
	docker-compose -f ${COMPOSE_FILE}  run --rm --entrypoint sh --no-deps app -c "make swag"

docker.lint:
	docker-compose -f ${COMPOSE_FILE}  run --rm --entrypoint sh --no-deps app -c "make lint"

docker.test:
	docker-compose -f ${COMPOSE_FILE}  run --rm --entrypoint sh --no-deps app -c "make test"

sh:
	docker-compose -f ${COMPOSE_FILE} run --rm --entrypoint sh --no-deps app -c "/bin/bash"

# todo connect to mysql inside the container to be independent
db:
	mysql -h ${DB_HOST} -P ${DB_PORT} --protocol=tcp -u ${DB_USER} -p${DB_PASSWORD}

init: swag