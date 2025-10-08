# Copyright 2024 Coralogix Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


ARG TARGETARCH
FROM rust:1.85.0-slim-bullseye AS interim
FROM interim AS interim-amd64
ENV ARCH=x86_64
FROM interim AS interim-arm64
ENV ARCH=aarch_64

FROM interim-${TARGETARCH}
ENV PB_REL="https://github.com/protocolbuffers/protobuf/releases"
ENV RUST_BACKTRACE=1

WORKDIR /app
ADD . /app

ADD $PB_REL/download/v25.1/protoc-25.1-linux-${ARCH}.zip /tmp

WORKDIR /app/tools/release-management
RUN cargo build --bin cleanup && mv target/debug/cleanup /app

WORKDIR /app/rust/examples
RUN apt update && \
    apt install -y unzip pkg-config libssl-dev curl && \
    unzip /tmp/protoc-25.1-linux-${ARCH}.zip -d /usr && \
    curl -LsSf https://get.nexte.st/latest/linux | tar zxf - -C ${CARGO_HOME:-~/.cargo}/bin && \
    cargo build --tests
