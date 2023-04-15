# base-client-go

## How to add config
 - json type config file add
   - see [config/socket_client.config](https://github.com/heaven-chp/base-client-go/blob/main/config/socket_client.config)
 - test add
   - see [config_test.go](https://github.com/heaven-chp/base-client-go/blob/main/config/config_test.go)
 - struct add
   - see [config/config.go](https://github.com/heaven-chp/base-client-go/blob/main/config/config.go)
 - example of use
   - socketClientConfig of [socket_client/main.go](https://github.com/heaven-chp/base-client-go/blob/main/socket_client/main.go)

## How to use grpc client
 - build
   - `go build -o grpc_client ./grpc_client/`
 - run
   - server must be running
     - [How to use grpc server](https://github.com/heaven-chp/base-server-go#how-to-use-grpc-server)
   - `./grpc_client/grpc_client -config_file config/grpc_client.config`
 - log
   - `./log/grpc_client_YYYYMMDD.log`

## How to use socket client
 - build
   - `go build -o socket_client ./socket_client/`
 - run
   - server must be running
     - [How to use socket server](https://github.com/heaven-chp/base-server-go#how-to-use-socket-server)
   - `./socket_client/socket_client -config_file config/socket_client.config`
 - log
   - `./log/socket_client_YYYYMMDD.log`
