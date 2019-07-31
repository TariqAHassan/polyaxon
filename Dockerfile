FROM polyaxon/polyaxon-npm-base

# This dockerfile is intended for dev only purposes
LABEL maintainer="mourad mourafiq <mourad@polyaxon.com>"

RUN apt-get -y update && \
    apt-get -y install curl unzip && \
    apt-get autoremove && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install Go
RUN mkdir -p /goroot && \
  curl https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /gopath
RUN mkdir -p $GOPATH/src
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# Install `protoc` v3.9.0
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.9.0/protoc-3.9.0-linux-x86_64.zip
RUN unzip protoc-3.9.0-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/bin/
RUN mv protoc3/include/* /usr/local/include/

# Go Deps
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger
# Hhhm???
RUN go install github.com/golang/protobuf/protoc-gen-go
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go install github.com/go-swagger/go-swagger/cmd/swagger

# Install Python gRPC tools.
RUN python -m pip install grpcio grpcio-tools

WORKDIR /polyaxon
