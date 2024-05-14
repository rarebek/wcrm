CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go


proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"


.PHONY: migrate-up migrate-down migrate-force migrate-file

CURRENT_DIR := $(shell pwd)
DB_URL := "postgres://postgres:asadbek@localhost:5432/productdb?sslmode=disable"

migrate-up:
	migrate -path migrations -database $(DB_URL) -verbose up

migrate-down:
	migrate -path migrations -database $(DB_URL) -verbose down

migrate-force:
	migrate -path migrations -database $(DB_URL) -verbose force 1

migrate-file:
	migrate create -ext sql -dir migrations/ -seq product
	

# build-image:
#   docker build --rm -t ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} .
#   docker tag ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG} ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

# push-image:
#   docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${TAG}
#   docker push ${REGISTRY}/${PROJECT_NAME}/${APP}:${ENV_TAG}

# .PHONY: pull-proto-module
# pull-proto-module:
#   git submodule update --init --recursive

# .PHONY: update-proto-module
# update-proto-module:
#   git submodule update --remote --merge