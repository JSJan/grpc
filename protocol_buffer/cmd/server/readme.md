
Building Service
-

Run the script to generate grpc related code in pb folder 
include main.go to up the server

- Add a listener
- Create a grpc server 
- Register your building service 
- Serve
- Implement the Start method from Building interface 
- Add reflection


Running the server
-

```go
grpcurl localhost:9292 list
```

Error 
-

The error tls: first record does not look like a TLS handshake typically occurs when grpcurl attempts to establish a TLS (secure) connection with a gRPC server that is running without TLS (insecure). By default, grpcurl assumes the connection should be secure.

Solution
-

To connect to a non-TLS gRPC server with grpcurl, you need to use the -plaintext option to indicate that the connection is insecure:

```go
- grpcurl -plaintext localhost:9292 list

- grpcurl -plaintext localhost:9292 list building.BuildingService

- grpcurl -plaintext localhost:9292 describe building.BuildingService
```