BINARY=engine

tools:
	@echo "installing air for hot reloading"
	go install github.com/cosmtrek/air@latest
	@echo "installing sql-migrate to handle schema migrations"
	go install github.com/rubenv/sql-migrate/...@latest
	@echo "installing swagger documentation"
	go install github.com/swaggo/swag/cmd/swag@latest
	@echo "installing GoMock for unit testing"
	go install github.com/golang/mock/mockgen@latest

test:
	make mock
	go test -coverprofile=resources/coverage.out ./...
	go tool cover -html=resources/coverage.out -o resources/coverage.html
	go tool cover -func resources/coverage.out

engine:
	go build -o ${BINARY} app/*.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-clean-arch .

run:
	docker-compose up --build -d

stop:
	docker-compose down

mock:
	go generate ./...

swag:
	swag init --parseDependency --output resources/webapps/swagger --outputTypes go,yaml

migrate:
	sql-migrate up --config=resources/db/dbconfig.yml --env=development

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint