# base-client-go

## Installation
```
go get -u github.com/heaven-chp/base-client-go
```

## How to add config
 - json type config file add
   - see config/socket_client.config
 - test add
   - see config_test.go
 - struct add
   - see config/config.go
 - example of use
   - socketClientConfig of socket_client/main.go

## How to use grpc client
 - install
   - go install github.com/heaven-chp/base-client-go/grpc_client
 - run
   - server must be running
     - go get -u github.com/heaven-chp/base-server-go
     - go install github.com/heaven-chp/base-server-go/grpc_server && ./bin/grpc_server -config_file src/github.com/heaven-chp/base-server-go/config/grpc_server.config
   - ./bin/grpc_client -config_file src/github.com/heaven-chp/base-client-go/config/grpc_client.config
 - log
   - ./log/grpc_client_YYYYMMDD.log

## How to use socket client
 - install
   - go install github.com/heaven-chp/base-client-go/socket_client
 - run
   - server must be running
     - go get -u github.com/heaven-chp/base-server-go
     - go install github.com/heaven-chp/base-server-go/socket_server && ./bin/socket_server -config_file src/github.com/heaven-chp/base-server-go/config/socket_server.config
   - ./bin/socket_client -config_file src/github.com/heaven-chp/base-client-go/config/socket_client.config 
 - log
   - ./log/socket_client_YYYYMMDD.log
