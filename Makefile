port ?= 8080

NAME 	= golang-daemon-starter
VERSION = `git rev-parse HEAD`
BUILD	= `date +%FT%T%z`
LDFLAGS = -ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

all: build run
help:
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
build: ## - Compile and build image
	@go test
	@GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"
	docker build --pull --build-arg PORT=$(port) -t $(NAME) .
	@go clean
run: ## - Run container
	docker run --rm -p $(port):$(port) -it $(NAME)
start: ## - Start detatched container
	docker run -p $(port):$(port) -e PORT=$(port) --name $(NAME) -d $(NAME)
stop: ## - Stop detatched container
	docker stop $(NAME)
	docker rm $(NAME)
clean: ## - Cleanup go and docker
	@go clean
	@docker container prune --force
	@docker image prune --force
	@docker image prune --filter "dangling=true" -f
	@docker rmi --force $(NAME)
info: ## - Show project information
	@echo name: $(NAME)
	@echo build: $(BUILD)
	@echo version: $(VERSION)