build:
	GOOS=darwin GOARCH=arm64 go build -o bin/exchange-tool-mac .
	GOOS=linux GOARCH=amd64 go build -o bin/exchange-tool-linux .

build-opti:
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/exchange-tool-mac .
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"  -o bin/exchange-tool-linux .

clean:
	rm -rf bin

run:
	go run .
