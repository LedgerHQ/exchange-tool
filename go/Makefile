build:
	GOOS=darwin go build -o bin/exchange-tool-mac .
	GOOS=linux go build -o bin/exchange-tool-linux .

build-opti:
	GOOS=darwin go build -ldflags="-s -w" -o bin/exchange-tool-mac .
	GOOS=linux go build -ldflags="-s -w"  -o bin/exchange-tool-linux .

clean:
	rm -rf bin

run:
	go run .
