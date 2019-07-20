# golang-daemon-starter

    This project builds a docker container which runs a simple http server in a Go daemon.

## Requirements

* Docker

## Install

```/bin/bash
# 1) clone git repository 
git clone https://github.com/StephenPerson/golang-daemon-starter.git # using git
go get github.com/StephenPerson/golang-daemon-starter # using go
# 2) build docker image
make build
# 3) start docker container
make start
# 4) visit local http server @ localhost:8080
## To see all available commands
make
```

## Reference

* https://socketloop.com/tutorials/golang-daemonizing-a-simple-web-server-process-example
