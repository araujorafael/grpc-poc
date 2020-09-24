FROM rust:latest as builder

ENV PROTOC_VERSION=3.11.0
RUN wget -q https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -O protoc.zip \
    && unzip protoc.zip -d /usr/local \
    && rm protoc.zip 

COPY . /mnt

WORKDIR /mnt/services/discount

RUN rustup target add x86_64-unknown-linux-musl \
    && rustup component add rustfmt --toolchain 1.39.0-x86_64-unknown-linux-gnu

RUN cargo build --release --target x86_64-unknown-linux-musl

####### service image #######

FROM alpine:latest
COPY --from=builder /mnt/services/discount/target/x86_64-unknown-linux-musl/release/discount  /usr/bin/discount-server
ENTRYPOINT ./usr/bin/discount-server

EXPOSE 3003
