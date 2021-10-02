# syntax = docker/dockerfile:experimental

# Base
FROM debian:buster-slim as base
RUN apt-get -y update && \
    apt-get -y install wget
ARG ARCH=amd64
ARG GO_VERSION=1.17
ENV GOPATH /go
# GOROOTのデフォルトは/usr/local/go
RUN wget https://golang.org/dl/go$GO_VERSION.linux-$ARCH.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go$GO_VERSION.linux-$ARCH.tar.gz && \
    rm -f go$GO_VERSION.linux-$ARCH.tar.gz 
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin
WORKDIR /workspace

# Workspace
FROM base as workspace
RUN apt -y update && \
    apt -y install unzip curl ssh git wget 
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf awscliv2.zip ./aws
RUN go install golang.org/x/tools/gopls@latest
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
WORKDIR /workspace

# Build
FROM base as build
COPY . /root/app
WORKDIR /root/app
RUN go build -o /go/bin/app cmd/api/main.go

# Release 
FROM gcr.io/distroless/base-debian10 as release
COPY --from=build /go/bin/app /
COPY --from=build /workspace/setup /setup
CMD ["/app"]
EXPOSE 8080

