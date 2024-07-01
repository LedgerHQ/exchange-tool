# Exchange Tool
This repo contains a suit of tools for testing, generating data and cross checking against different tech env.

## Foreword
The swap payload used by the Nano is in Protobuf format.

The private and public key are expected to be in ES256 format, using a secp256k1 or secp256r1 curve.

## Usage
### Generate
`./exchange-tool generate <CURVE> <PRIVATE_KEY> <PAYLOAD_JSON_FORMAT>`

Example:

`./exchange-tool generate k1 ./samples/sample-priv-key-secp256k1.pem payload-example.json`

### Check
`./exchange-tool check <CURVE> <PUBLIC_KEY> <BINARY_PAYLOAD_BASE64> <SIGNATURE_BASE64>`

Example:

`./exchange-tool check k1 ./samples/sample-pub-key-secp256k1.pem CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoIMTIwMDAwMDBSCDExNTAwMDAwWgpBQkNERUZHSElK 5-J8C2lb9bZj2yGWaNCjKyW15mDx3zaYc3u59Bag7t-G0-vjzpadZzWTHMGUJeY2IJMr5NxQV5RqdFemOvbaWQ==`

### Read
`./exchange-tool read <BINARY_PAYLOAD_BASE64>`

Example:

`./exchange-tool read CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoIMTIwMDAwMDBSCDExNTAwMDAwWgpBQkNERUZHSElK`

### Convert asn1 pubkey to hex value
`./exchange-tool hex <PROVIDER_PUBLIC_KEY>`

Example:

`./exchange-tool hex ./samples/sample-pub-key-secp256k1.pem`

### Sign Provider
`./exchange-tool sign <PROVIDER_NAME> <PROVIDER_PUBLIC_KEY>`

Example:

`./exchange-tool sign SELL_TEST ./samples/sample-pub-key-secp256k1.pem`

### Generate CAL format for Provider info
`./exchange-tool cal -n <PROVIDER_NAME> -c <CURVE> -p <PROVIDER_PUBLIC_KEY> -v <VERSON> -a <APPLICATION_NAME>`


Example:

`./exchange-tool cal -c k1 -p ../sample-public-key.pem -n moonpay -v 2 -a swap`

**Disclaimer**
DO NOT USE private key provided in this repository. Their goal is for test purpose only.

The code use to generate a payload has an amount limitation. To ease the JSON serialization, we are relying on `uint64` type, which is not the best to suit crypto units.

## Generate private/public key
Private key:
```sh
  openssl ecparam -name secp256k1 -genkey -noout -out sample-priv-key-secp256k1.pem
```
Public key (from private key):
```sh
  openssl ec -in sample-priv-key.pem -pubout > sample-pub-key-secp256k1.pem
```

Read file info
```sh
  openssl ec -text -in sample-priv-key-secp256k1.pem
```

(ecparam is for Eliptic Curves params)

## Cross testing and consistency
To check if the code is still consistent with what is expected from Ledger:
```sh
  go test .

  # For log info
  go test . -v
```

### Python
```sh
  pipenv install
  pipenv install cryptography
  pipenv run python swap_test.py
```

### Bash
```sh
  # Sign 'encoded_payload' content
  openssl dgst -sha256 -sign sample-priv-key-secp256r1.pem encoded_payload.txt > signature.der
  # Display R and S from DER format signature
  openssl asn1parse -inform DER -in signature.der
  # Verify signature
  openssl dgst -sha256 -verify sample-pub-key-secp256r1.pem -signature signature.der encoded_payload.txt
```

With current example value
```sh
  openssl dgst -sha256 -verify example/sample-pub-key-secp256k1.pem -signature example/signature-example-k1.der example/encoded_payload.txt
```

#### Generate DER file from R, S value
1. Copy `asn1rs.template.cnf` file (ex: into `asn1rs-example.cnf`)
2. Edit this file and replace `{rvalue_in_hex}` and `{svalue_in_hex}` value with R and S value
2. Call `openssl asn1parse -genconf asn1rs-example.cnf -out signature-example.der`

#### Extract R, S value from base64url signature
```sh
  var=$(basenc --base64url -d signature.txt | hexdump -ve '1/1 "%.2x"')
  printf '%s\n' "${var:0:${#var}/2}" "${var:${#var}/2}"
  unset var
```

## Source
[secp256k1 vs secp256r1](https://www.johndcook.com/blog/2018/08/21/a-tale-of-two-elliptic-curves/)

[OID secp256r1](http://www.oid-info.com/get/1.2.840.10045.3.1.7)

[OID secp256k1](http://www.oid-info.com/get/1.3.132.0.10)

[sec1](https://www.secg.org/sec1-v2.pdf)

[sec2](https://www.secg.org/sec2-v2.pdf)

[JWS RFC](https://www.rfc-editor.org/rfc/rfc7515#section-5)

[JWA ES256](https://www.rfc-editor.org/rfc/rfc7518#section-3.4)

[DER file from RS value](https://superuser.com/questions/1653062/how-can-i-convert-my-plain-text-r-s-signature-to-a-format-that-openssl-can-ver)
