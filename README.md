# Base Client for Go

[![CI](https://github.com/base-client/go/workflows/CI/badge.svg)](https://github.com/base-client/go/actions)
[![Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/heaven-chp/8d4d04368033a33c6220507686f78072/raw/coverage.json)](https://github.com/base-client/go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/base-client/go)](https://goreportcard.com/report/github.com/base-client/go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/base-client/go?logo=go)](https://github.com/base-client/go)
[![Reference](https://pkg.go.dev/badge/github.com/base-client/go.svg)](https://pkg.go.dev/github.com/base-client/go)
[![License](https://img.shields.io/github/license/base-client/go)](https://github.com/base-client/go/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/base-client/go)](https://github.com/base-client/go/stargazers)

<br/>

## Features
 - cloudevents
 - grpc
 - http
 - long polling
 - socket

<br/>

## How to add config
 - config file is yaml format
   - see [common/config/config.yaml](https://github.com/base-client/go/blob/main/common/config/config.yaml)
 - struct add
   - see [common/config/config.go](https://github.com/base-client/go/blob/main/common/config/config.go)
 - test add
   - see [common/config/config_test.go](https://github.com/base-client/go/blob/main/common/config/config_test.go)
 - example of use
   - see client main.go files ([grpc-client/main.go](https://github.com/base-client/go/blob/main/grpc-client/main.go), [socket-client/main.go](https://github.com/base-client/go/blob/main/socket-client/main.go), etc.)

<br/>

## How to use client
 - prepare
   - run
     - server must be running
       - [How to use server](https://github.com/base-server/go#how-to-use-server)
 - cloudevents
   - build
     - `go build -o ./bin/cloudevents-client ./cloudevents-client/`
   - run
     - `./bin/cloudevents-client -config-file ./common/config/config.yaml`
 - grpc
   - build
     - `go build -o ./bin/grpc-client ./grpc-client/`
   - run
     - `./bin/grpc-client -config-file ./common/config/config.yaml`
   - log
     - `./common/log/grpc-client_YYYYMMDD.log`
 - http
   - build
     - `go build -o ./bin/http-client ./http-client/`
   - run
     - `./bin/http-client -config-file ./common/config/config.yaml`
   - log
     - `./common/log/http-client_YYYYMMDD.log`
 - long-polling
   - build
     - `go build -o ./bin/long-polling-client ./long-polling-client/`
   - run
     - `./bin/long-polling-client -config-file ./common/config/config.yaml`
   - log
     - `./common/log/long-polling-client_YYYYMMDD.log`
 - socket
   - build
     - `go build -o ./bin/socket-client ./socket-client/`
   - run
     - `./bin/socket-client -config-file ./common/config/config.yaml`
   - log
     - `./common/log/socket-client_YYYYMMDD.log`

<br/>

## Test and Coverage
 - Test
   - `go clean -testcache && go test -cover ./...`
 - Coverage
   - make coverage file
     - `go clean -testcache && go test -coverprofile=coverage.out -cover ./...`
   - convert coverage file to html file
     - `go tool cover -html=./coverage.out -o ./coverage.html`
