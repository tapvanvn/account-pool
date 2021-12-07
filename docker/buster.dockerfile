FROM golang:1.16-buster  AS build

WORKDIR /

RUN apt update && apt install -y git  gcc g++ musl-dev curl openssh-client

RUN mkdir -p /src

COPY ./ /src/

RUN cd /src && go get ./... && go build -o accpool

FROM debian:buster-slim
RUN apt update \
        && apt upgrade \
        && apt install --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
        
COPY --from=build               /src/accpool / 
COPY static/                    /static
COPY route.jsonc                /route.jsonc 

ENV PORT=80

ENTRYPOINT ["/accpool"]