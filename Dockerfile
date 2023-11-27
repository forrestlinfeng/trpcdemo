## build
FROM golang:1.18-buster AS build-env

COPY . /go/src/build

WORKDIR /go/src/build

RUN make build

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
