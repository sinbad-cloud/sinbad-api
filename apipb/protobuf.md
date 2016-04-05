# protobuf

## How to generate protobuf code?

Generate protobuf classes, this will generate the Go file for messages in the `apipb/` folder.

    protoc apipb/deployment.proto --go_out=./ 

But what we want is to generate the client and server code as well. In order to do this, run

    protoc -I ./apipb apipb/deployment.proto --go_out=plugins=grpc:apipb