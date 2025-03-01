#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

for dir in $proto_dirs; do
	protoc \
		-I "./proto" \
  		-I "$dir" \
		-I third_party/grpc-gateway/ \
		-I third_party/googleapis/ \
		-I third_party/proto/ \
		--go_out=plugins=grpc,paths=source_relative:./proto-gen \
		--grpc-gateway_out=paths=source_relative:./proto-gen \
		--openapiv2_out=third_party/OpenAPI/ \
		$(find "${dir}" -maxdepth 1 -name '*.proto')

# hack auto-generated proto -> remove omitempty tag
# find ./proto-gen -path -prune -o -name '*.pb.go' | xargs -n1 -IX bash -c "sed -e 's/,omitempty//' X > X.tmp && mv X{.tmp,}"

done
