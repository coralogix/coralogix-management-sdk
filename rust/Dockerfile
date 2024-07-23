FROM rust:1.79.0-slim-bullseye
ENV PB_REL="https://github.com/protocolbuffers/protobuf/releases"
ENV RUST_BACKTRACE=1

WORKDIR /app
ADD . /app
WORKDIR /app/rust
RUN rustup update nightly && rustup default nightly
ADD $PB_REL/download/v25.1/protoc-25.1-linux-x86_64.zip /tmp
RUN apt update && apt install unzip && unzip /tmp/protoc-25.1-linux-x86_64.zip -d /usr && cargo build