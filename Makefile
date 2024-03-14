PATH := $(PATH):$(shell go env GOPATH)/bin

.PHONY: postgresinit
postgresinit:
	@echo "Create and start postgres database in docker"
	docker run --name film-library -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

.PHONY: dropdb
dropdb:
	@echo "Drop postgres database from docker"
	docker exec -it postgres dropdb film-library

.PHONY: test
test:
	@echo "Start tests"
	go clean -testcache &&  go test ./server/internal/service/ ./server/internal/handler/

.PHONY: deps
deps:
	go mod tidy

.PHONY: swagger_init
swagger_init:
	@echo "Generate swagger API"
	swag init --generalInfo  server/cmd/main.go -o server/docs

.PHONY: swag_ui
swag_ui:
	@echo "Open swagger index.html"
	open http://localhost:8000/swagger/index.html

.PHONY: build
build: deps swagger_init
	@echo "Running docker-compose"
	docker-compose up