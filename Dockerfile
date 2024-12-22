FROM golang:1.23.4-alpine

RUN apk add --no-cache \
    protobuf \
    protobuf-dev \
    grpc \
    grpc-dev \
    bash \
    curl \
    git \
    build-base \
    libc-dev

RUN apk add --no-cache git bash && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="$PATH:/root/go/bin"

WORKDIR /app

CMD ["sh", "-c", "set -x && rm -rf ./generated && mkdir -p generated && mkdir -p vendor && find . -name '*.proto' | xargs protoc \
    -I=. \
    -I=vendor \
    --descriptor_set_out=generated/websub.protoset \
    --include_imports \
    --include_source_info \
    --go_out=paths=source_relative:generated \
    --go-grpc_out=paths=source_relative:generated \
    && go mod tidy && go mod vendor"]