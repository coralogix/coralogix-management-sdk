# Copyright 2026 Coralogix Ltd.
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
FROM golang:1.25-alpine AS interim
FROM interim AS interim-amd64
ENV ARCH=x86_64
FROM interim AS interim-arm64
ENV ARCH=aarch_64

FROM interim-${TARGETARCH}


RUN apk add --no-cache make && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Set working directory
WORKDIR /app

# Copy Makefile and required files
COPY ../ .
# Set environment variables
ENV CORALOGIX_REGION=us
ENV CORALOGIX_API_KEY=test_key

# Run tests using Makefile target
CMD ["make", "test-openapi-ro"]