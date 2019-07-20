# golang-daemon-starter

    This project builds a docker container which runs a simple http server in a Go daemon.

## Requirements

* Docker
* Golang

## Install
```bash
git clone git@github.com:personjs/golang-daemon-starter.git ## using git
go get github.com/personjs/golang-daemon-starter ## using go
cd <project-directory>
```

## Local
```bash
# BUILD
go build
# RUN
./golang-daemon-starter [console|start|stop|status]
```

## Container

```bash
# BUILD
make build
# RUN
make [run|start|stop]
# REMOVE
make clean
```

## Reference

* https://socketloop.com/tutorials/golang-daemonizing-a-simple-web-server-process-example
