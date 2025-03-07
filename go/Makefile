# SPDX-FileCopyrightText: 2024 Ledger
# SPDX-License-Identifier: Apache-2.0

build:
	go mod tidy
	GOOS=darwin GOARCH=arm64 go build -o bin/exchange-tool-mac-arm64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/exchange-tool-mac-amd64 .
	GOOS=linux GOARCH=amd64 go build -o bin/exchange-tool-linux .

build-opti:
	go mod tidy
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/exchange-tool-mac-arm64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/exchange-tool-mac-amd64 .
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/exchange-tool-linux .

clean:
	rm -rf bin

run:
	go run .

test:
	go test ./...
