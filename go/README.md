# Go exchange-tool

## Compilation
If you want to build an exec, simply: `go build -o SwapTool .`

If you need to compile for different platform:
 * Mac : `GOOS=darwin GOARCH=amd64 go build -o SwapTool .`
 * Linux : `GOOS=linux GOARCH=amd64 go build -o SwapTool .`

(both example target 64 bits architecture ; for 32 bits, use `386` value for `GOARCH` var).

[Go official doc](https://go.dev/doc/install/source#environment)

## Usage
### Generate
`go run . generate <CURVE> <PRIVATE_KEY> <PAYLOAD_JSON_FORMAT>`

Example:

`go run . generate k1 ../samples/sample-priv-key-secp256k1.pem payload-example.json`

### Check
`go run . check <CURVE> <PUBLIC_KEY> <BINARY_PAYLOAD_BASE64> <SIGNATURE_BASE64>`

Example:

`go run . check k1 ../samples/sample-pub-key-secp256k1.pem CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoIMTIwMDAwMDBSCDExNTAwMDAwWgpBQkNERUZHSElK 5-J8C2lb9bZj2yGWaNCjKyW15mDx3zaYc3u59Bag7t-G0-vjzpadZzWTHMGUJeY2IJMr5NxQV5RqdFemOvbaWQ==`

### Read
`go run . read <BINARY_PAYLOAD_BASE64>`

Example:

`go run . read CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoIMTIwMDAwMDBSCDExNTAwMDAwWgpBQkNERUZHSElK`

### Convert asn1 pubkey to hex value
`go run . hex <PROVIDER_PUBLIC_KEY>`

Example:

`go run . hex ../samples/sample-pub-key-secp256k1.pem`

### Sign Provider
`go run . sign <PROVIDER_NAME> <PROVIDER_PUBLIC_KEY>`

Example:

`go run . sign SELL_TEST ../samples/sample-pub-key-secp256k1.pem`

### Generate CAL format for Provider info
`go run . cal <PROVIDER_NAME> <CURVE> <PROVIDER_PUBLIC_KEY>`

Example:

`go run . cal SELL_TEST k1 ../samples/sample-pub-key-secp256k1.pem`

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
