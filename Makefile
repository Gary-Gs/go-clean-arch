BINARY=engine

tools:
	@echo "installing air for hot reloading"
	go get -u github.com/cosmtrek/air
	@echo "installing sql-migrate to handle schema migrations"
	go get -v github.com/rubenv/sql-migrate/...
	@echo "installing swagger documentation"
	go get -u github.com/swaggo/swag/cmd/swag

test:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out

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