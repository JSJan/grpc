
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

Run the protoc command to generate Go code from your .proto file
protoc --go_out=. example.proto


**Error encountered **

protocol_buffer % protoc --go_out=. building.proto
building.proto:3:1: Expected ";".

protocol_buffer % protoc --go_out=. building.proto
building.proto:4:16: Missing field number.
building.proto:5:16: Missing field number.
building.proto:6:20: Missing field number.

protocol_buffer % protoc --go_out=. building.proto
building.proto:6:14: "string" is already defined in "Building".
building.proto:4:5: "name" is not defined.
building.proto:5:5: "area" is not defined.
building.proto:6:5: "timezone" is not defined.

Wrong declaration of fields
message Building {
    name string =1;
    area double =2;
    timezone string=3;
}

protocol_buffer % protoc --go_out=. building.proto
protoc-gen-go: unable to determine Go import path for "building.proto"

Please specify either:
        • a "go_package" option in the .proto source file, or
        • a "M" argument on the command line.

See https://protobuf.dev/reference/go/go-generated#package for more information.


protoc --go_out=. --go_opt=paths=source_relative --go_opt=Mbuilding.proto=github.com/jsjan/yourproject/building building.proto

go generate - take the input commands from gen.go