FROM golang:1.18.3

RUN apt-get update && apt-get install -y unzip  
  
RUN set -eux; \
  arch="$(dpkg --print-architecture)"; \
  url=; \
  case "$arch" in \
    'amd64') url='https://github.com/protocolbuffers/protobuf/releases/download/v21.1/protoc-21.1-linux-x86_64.zip' ;; \
    'arm64') url='https://github.com/protocolbuffers/protobuf/releases/download/v21.1/protoc-21.1-linux-aarch_64.zip' ;; \
    *) echo >&2 "error: unsupported architecture '$arch' (likely packaging update needed)"; exit 1 ;; \
  esac; \
  mkdir -p /tmp/protoc && \
  curl -L "$url" > /tmp/protoc/protoc.zip && \
  cd /tmp/protoc && \
  unzip protoc.zip && \
  mv include/* /usr/local/include/ && \
  cp /tmp/protoc/bin/protoc /usr/local/bin && \  
  chmod go+rx /usr/local/bin/protoc && \  
  cd /tmp && \  
  rm -r /tmp/protoc  
  
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
