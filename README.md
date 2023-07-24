# base-client-go

## How to add config
 - json type config file add
   - see [config/socket_client.config](https://github.com/heaven-chp/base-client-go/blob/main/config/socket_client.config)
 - test add
   - see [config_test.go](https://github.com/heaven-chp/base-client-go/blob/main/config/config_test.go)
 - struct add
   - see [config/config.go](https://github.com/heaven-chp/base-client-go/blob/main/config/config.go)
 - example of use
   - socketClientConfig of [socket-client/main.go](https://github.com/heaven-chp/base-client-go/blob/main/socket-client/main.go)

## How to use client
 - prepare
   - run
     - server must be running
       - [How to use server](https://github.com/heaven-chp/base-server-go#how-to-use-server)
 - grpc
   - build
     - `go build -o ./bin/grpc-client ./grpc-client/`
   - run
     - `./bin/grpc-client -config_file config/grpc_client.config`
   - log
     - `./log/grpc-client_YYYYMMDD.log`
 - socket
   - build
     - `go build -o ./bin/socket-client ./socket-client/`
   - run
     - `./bin/socket-client -config_file config/socket_client.config`
   - log
     - `./log/socket-client_YYYYMMDD.log`
