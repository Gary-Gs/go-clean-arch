# go-clean-arch

This repo is a backend service boilerplate code, a clean architecture framework implementation for Go (with echo).

Initially forked from bxcodec/go-clean-arch, I have done some optimization and refactoring to make it more efficient and easier to use for greenfield projects.

Just simply pull and get started.

### Description
This is an example of implementation of Clean Architecture in Go (Golang) projects.

Rule of Clean Architecture by Uncle Bob
 * Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
 * Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
 * Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
 * Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
 * Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

This project has  4 Domain layer :
 * Model Layer
 * Repository Layer
 * Usecase Layer  
 * Delivery Layer

#### The diagram:

![golang clean architecture](https://github.com/Gary-Gs/go-clean-arch/raw/master/resources/clean-arch.png)

### How to run this project in containerized environment
pre-requisite: golang v1.6+ and docker

```bash
# move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/Gary-Gs/go-clean-arch.git

# move to project
$ cd go-clean-arch

# install make (macOS) or `apt-get install make` (Linux)
$ brew install make

# install tools
$ make tools

# Bring up all dependencies and services
$ docker-compose up

# check if the containers are running
$ docker ps

# Execute the call
$ curl localhost:9090/api/v1/articles

# Stop
$ make stop

# Generate API Documentation
# http://localhost:9090/swagger/index.html
$ make swag
```

## Tools used
Detailed list can be found in `go.mod` and `Makefile`

## Features out of the box
- Example CRUD APIs 
- Containerized with Docker
- API documentation with Swagger
- DB migration support
- Example unit test with GoMock & GoMonkey

test