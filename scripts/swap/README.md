# Check Swap Signature Script

This script verifies the swap service's response signature.

## Prerequisites

1. Install `jq`: [Download jq](https://stedolan.github.io/jq/download/)
2. Install `go`: [Downlaod go](https://golang.org/doc/install)
3. Install `make`: [Downlaod make](https://www.gnu.org/software/make/)

## Usage

1. Put the public key file in the `keys` directory.
2. Replace the `curl` command in the script with the actual one to fetch the swap response and save it in `payload.json`.
3. Fill `BINARY_PAYLOAD_BASE64` and `SIGNATURE_BASE64` with the values from `payload.json` using `jq`.
4. Run the script with the following parameters:
   - `$1`: `<public_key_file>`: The public key file name.
   - `$2`: `<os>`: The operating system of the machine where the script is running: `mac` or `linux`.

```sh
./check_swap_signature_script.sh <public_key_file> <os>
```

## Example

Assuming you have a public key file named public_key.pem and you're running the script on a Mac, you would run the script like this:
```sh
./check_swap_signature_script.sh public_key.pem mac