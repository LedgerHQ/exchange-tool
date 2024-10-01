# Go exchange-tool

## Compilation
If you want to build an executable, simply: `make build` (you can use the alternate command `make build-opti`)

## Usage in dev
In root directory, replace README command example `./exchange-tool` by `go run .`

## Check test
`make test`

## Generate go file from protobuf
```sh
  # Without adding `option go_package = "swap.ledger.fr/proto";` in protocol.proto file
  protoc --go_out=Mprotocol.proto=swap.ledger.fr/proto protocol.proto

  protoc --go_out=. proto/protocol.proto
```

## Testing and consistency
To check if the code is still consistent with what is expected from Ledger:
```sh
  go test .

  # For log info
  go test . -v
```

## Check OSS dependency
This part relies on third party tool: [OSS Review Toolkit](https://oss-review-toolkit.org/ort/).

You must use Docker (Podman seems not to work without heavy settings).

On Mac (M1 and M2 processor), you have to compile yourself the [tool](https://github.com/oss-review-toolkit/ort).

Launch analysis:
```sh
#!/bin/bash

docker run \
  -v $PWD/:/project \
  ort --info \
  analyze -i /project \
  -o /project

```

Then launch report:
```sh
docker run \
  -v $PWD/:/project \
  ort evaluate \
  --license-classifications-file /project/.ort/config/license-classifications.yml \
  --package-curations-file /project/.ort/config/curations.yml \
  --rules-file /project/.ort/config/evaluator.rules.kts \
  -i /project/analyzer-result.yml \
  -o /project
```

## Source
[Protobuf official](https://protobuf.dev/)

[Protobuf Go package](https://pkg.go.dev/google.golang.org/protobuf)

[Ethereum Crypto package](https://pkg.go.dev/github.com/ethereum/go-ethereum/crypto)
