NAME=golang-daemon-starter
VERSION=`git rev-parse HEAD`
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.PHONY: help
help:
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
.PHONY: clean
clean: ## - Cleanup docker
	@docker stop $(NAME) > /dev/null 2>&1 | true
	@docker rm $(NAME) > /dev/null 2>&1 | true
	@docker rmi -f $(NAME) > /dev/null 2>&1 | true
.PHONY: build
build: ## - Build image
	docker build -t $(NAME) .
.PHONY: start
start: ## - Run container
	docker run --net=host --name $(NAME) -d $(NAME)
.PHONY: stop
stop: ## - Stop running container
	@docker stop $(NAME)
	@docker rm $(NAME)
.PHONY: info
info: ## - Show project information
	@echo name: $(NAME)
	@echo build: $(BUILD)
	@echo version: $(VERSION)
