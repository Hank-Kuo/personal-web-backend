include .env

.PHONY: migrate migrate_down migrate_up migrate_version docker prod docker_delve local swaggo test
VERSION ?= $(shell git describe --tags --always)
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
DOCKER_LOGIN ?= 
IMAGE_NAME ?= hank-backend-server
TAG ?=
MIG_VER ?=

# mysql 
MYSQL_HOST ?=
MYSQL_USER ?=
MYSQL_PASSWORD ?=
MYSQL_ROOT_PASSWORD ?=
MYSQL_DATABASE ?=

# postgre
PG_HOST ?=
PG_USER ?=
PG_PASSWORD ?=
PG_ROOT_PASSWORD ?=
PG_DATABASE ?=


# Go command
run:
	go run ./cmd/main.go
build:
	go build -o bin/ ./cmd/main.go
test:
	go test -cover ./...

### Go modules support
deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

### Go migrate postgresql
pg-force:
	migrate -database postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):5432/$(PG_DATABASE)?sslmode=disable -path migrations/pg force $(MIG_VER)
pg-version:
	migrate -database postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):5432/$(PG_DATABASE)?sslmode=disable -path migrations/pg version
pg-up:
	migrate -database postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):5432/$(PG_DATABASE)?sslmode=disable -path migrations/pg up $(MIG_VER)
pg-down:
	migrate -database postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):5432/$(PG_DATABASE)?sslmode=disable -path migrations/pg down $(MIG_VER)

### Go migrate mysql
mysql-force:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):3306)/$(MYSQL_DATABASE)" -path migrations/mysql force $(MIG_VER)
mysql-version:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):3306)/$(MYSQL_DATABASE)" -path migrations/mysql version
mysql-up:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):3306)/$(MYSQL_DATABASE)" -path migrations/mysql up $(MIG_VER)
mysql-down:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):3306)/$(MYSQL_DATABASE)" -path migrations/mysql down $(MIG_VER)


# Docker compose commands
dev:
	@go run ./cmd/main.go

# Docker commands
DOCKER_CMD := $(shell docker ps -aq)
kill:
	docker kill $(DOCKER_CMD)
stop:
	docker stop $(DOCKER_CMD)
restart:
	docker restart $(DOCKER_CMD)
remove:
	docker rm $(DOCKER_CMD)
down:
	docker-compose down
up:
	docker-compose -f docker-compose.yml up -d --build --remove-orphans

# Tools commands
linter:
	echo "Starting linters"
	golangci-lint run ./...

# Swagger
swagger:
	echo "Starting swagger generating"
	swag init -g **/**/*.go