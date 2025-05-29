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

default: build

build: 
	cd ./go; make build

release:
	cd ./go; make release

test:
	cd ./go; make test TESTARGS=${TESTARGS}

proto-renew:
	protofetch clear-cache
	protofetch clean
	protofetch fetch
	cp -a proto/src/** proto/
	rm -rf proto/src
	rm -rf proto/deps

proto-go-generate: proto-renew
	cd ./go; make proto-clean; make proto-compile

pull-and-build: proto-renew build
