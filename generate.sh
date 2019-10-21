#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

go run calculator/calculate_server/server.go 
go run calculator/calculate_client/client.go 