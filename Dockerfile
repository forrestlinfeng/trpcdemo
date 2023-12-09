FROM ccr.ccs.tencentyun.com/dops/alpine:3.9

LABEL maintainer="linfengchen"

RUN apk add tcpdump net-tools busybox-extras && mkdir -p /app

WORKDIR /app

COPY trpcdemo /app/trpcdemo
COPY trpc_go.yaml /app/trpc_go.yaml

ENV PATH $PATH:/app

EXPOSE 8000
CMD ["./trpcdemo"]
