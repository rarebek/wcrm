# -include .env
# export

# CURRENT_DIR=$(shell pwd)
# APP=content_service
# CMD_DIR=./cmd

# .DEFAULT_GOAL = build

# # build for current os
# .PHONY: build
# build:
# 	go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# # build for linux amd64
# .PHONY: build-linux
# build-linux:
# 	CGO_ENABLED=0 GOARCH="amd64" GOOS=linux go build -ldflags="-s -w" -o ./bin/${APP} ${CMD_DIR}/app/main.go

# # run service
# .PHONY: run
# run:
# 	go run ${CMD_DIR}/app/main.go

# # migrate
# .PHONY: migrate
# migrate:
# 	migrate -source file://migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable up

# # proto
# .PHONY: proto-gen
# proto-gen:
# 	./scripts/gen-proto.sh

# # git submodule init 	
# .PHONY: pull-proto
# pull-proto:
# 	git submodule update --init --recursive

# # go generate	
# .PHONY: go-gen
# go-gen:
# 	go generate ./...

# # run test
# test:
# 	go test -v -cover -race ./internal/...

# # -------------- for deploy --------------
# build-image:
# 	docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
# 	docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

# push-image:
# 	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
# 	docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go
proto-gen:
	./scripts/gen-proto.sh

run:
	go run cmd/app/main.go

migrate-up:
	migrate -path migrations -database "postgresql://postgres:123@localhost:5432/orders?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgres://postgres:123@localhost:5432/orders?sslmode=disable" -verbose down

migrate_file:
	migrate create -ext sql -dir migrations/ -seq orders

migrate-dirty:
	migrate -path ./migrations/ -database "postgresql://postgres:123@localhost:5432/orders?sslmode=disable" force 1