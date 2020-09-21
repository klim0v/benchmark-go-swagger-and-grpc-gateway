#!/usr/bin/env bash

cd "$(dirname "$0")" || exit

protoc -I. \
		--grpc-gateway_out=logtostderr=true:./ \
		--swagger_out=allow_merge=true,merge_file_name=api:. \
		--go_out=plugins=grpc:./shop ./shop/*.proto