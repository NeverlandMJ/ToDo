#!/bin/bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   v1/todopb/todo.proto v1/userpb/user.proto