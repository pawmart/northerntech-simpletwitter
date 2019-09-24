.PHONY: help

APP_NAME=northerntech-simpletwitter-srv
MONGO_NAME=northerntech-mongo
PACKAGE_NAME=github.com/pawmart/northerntech-simpletwitter
VERSION=latest

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

localtest: ## Run godog locally.
	NORTHTECH_DB_DATABASE='northerntech-simpletwitter-test' \
	NORTHTECH_DB_HOST='localhost' \
	NORTHTECH_DB_USER='root' \
    NORTHTECH_DB_PASSWORD='example' \
    NORTHTECH_DB_AUTH='admin' \
    go test ./...

localbuild: localtest ## Run build manually.
	dep ensure && go build -o app

localrun: localbuild ## Run locally.
	NORTHTECH_DB_DATABASE='northerntech-simpletwitter-app' \
	NORTHTECH_DB_HOST='localhost' \
	NORTHTECH_DB_USER='root' \
	NORTHTECH_DB_PASSWORD='example' \
	NORTHTECH_DB_AUTH='admin' \
	./app

test: mongo ## Test.
	docker run --rm \
		-v ${PWD}:/go/src/${PACKAGE_NAME} \
		-v ${HOME}/.ssh/id_rsa:/root/.ssh/id_rsa \
		-w /go/src/${PACKAGE_NAME} \
		-e NORTHTECH_DB_DATABASE='northerntech-simpletwitter-test' \
		-e NORTHTECH_DB_HOST='host.docker.internal' \
		-e NORTHTECH_DB_USER='root' \
		-e NORTHTECH_DB_PASSWORD='example' \
		-e NORTHTECH_DB_AUTH='admin' \
		rat4m3n/go-builder:latest /bin/sh -c "dep ensure && godog"

build: test ## Build.
	docker run --rm \
		-v ${PWD}:/go/src/${PACKAGE_NAME} \
		-v ${HOME}/.ssh/id_rsa:/root/.ssh/id_rsa \
		-w /go/src/${PACKAGE_NAME} \
		-e GOOS=linux \
		-e GOARCH=386 \
		rat4m3n/go-builder:latest /bin/sh -c "dep ensure && go build -o app" && \
	docker build -t ${APP_NAME}:${VERSION} .

up: build ## Build and Run.
	docker run -d \
		--name=${APP_NAME} \
		-p 6543:6543 \
		-e NORTHTECH_DB_DATABASE='northerntech-simpletwitter-app' \
		-e NORTHTECH_DB_HOST='host.docker.internal' \
		-e NORTHTECH_DB_USER='root' \
		-e NORTHTECH_DB_PASSWORD='example' \
		-e NORTHTECH_DB_AUTH='admin' \
		${APP_NAME}

mongo: ## Start storage.
	docker run -d \
		--name ${MONGO_NAME} \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME='root' \
		-e MONGO_INITDB_ROOT_PASSWORD='example' \
		mongo:3.4-jessie

down: ## Clear.
	docker stop ${APP_NAME} && docker rm ${APP_NAME}
	docker stop ${MONGO_NAME} && docker rm ${MONGO_NAME}

validate: ## Validate swagger.
	${GOPATH}/bin/swagger validate ./api/swagger/swagger.yml

gen: validate ## Generate server from swagger.
	${GOPATH}/bin/swagger generate server \
		--target=./internal/ \
		--spec=./api/swagger/swagger.yml \
		--exclude-main \
		--name=northerntech-simpletwitter

deploy:
	docker build -t rat4m3n/northerntech-simpletwitter .
