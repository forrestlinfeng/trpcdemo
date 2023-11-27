## build
FROM golang:alpine AS build-env

COPY . /go/src/build

WORKDIR /go/src/build

RUN env GOOS=linux go build -o trpcdemo

## run
FROM alpine:3.9

LABEL maintainer="linfengchen"

RUN mkdir -p /app

WORKDIR /app

COPY --from=build-env /go/src/build/trpcdemo /app/trpcdemo
COPY --from=build-env /go/src/build/trpc_go.yaml /app/trpc_go.yaml

ENV PATH $PATH:/app

EXPOSE 8000
CMD ["./trpcdemo"]
