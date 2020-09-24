FROM golang:alpine as build-env

ENV GOBIN="/go/bin"
ENV PATH="$GOBIN/bin:$PATH"

ENV GLIBC_VERSION=2.30-r0
RUN apk --no-cache add wget \
    && wget -q https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub -O /etc/apk/keys/sgerrand.rsa.pub \
    && wget -q https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk -O glibc.apk \
    && apk add glibc.apk \
    && rm /etc/apk/keys/sgerrand.rsa.pub glibc.apk

ENV PROTOC_VERSION=3.11.0
RUN wget -q https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -O protoc.zip \
    && unzip protoc.zip -d /usr/local \
    && rm protoc.zip \
    && apk del wget

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ADD . /usr/bin/server
WORKDIR /usr/bin/server/services/product

RUN go build -ldflags "-s -w" -o product-server

# ------ Server Image -------

FROM alpine
WORKDIR /usr/bin
COPY --from=build-env /usr/bin/server/services/product/product-server ./
ENTRYPOINT /usr/bin/product-server

EXPOSE 8080