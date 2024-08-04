build:
	GOOS=darwin GOARCH=arm64 go build -o bin/exchange-tool-arm .
	GOOS=linux GOARCH=amd64 go build -o bin/exchange-tool-amd .

build-opti:
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/exchange-tool-arm .
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"  -o bin/exchange-tool-amd .

clean:
	rm -rf bin

run:
	go run .

test:
	go test .