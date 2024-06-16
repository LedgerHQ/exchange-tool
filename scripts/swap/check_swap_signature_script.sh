#!/bin/bash

# This script is used to check the signature of the swap response from the swap service.

# Prerequisites:
# 1. Install jq: https://stedolan.github.io/jq/download/
# 2. Install go: https://golang.org/doc/install
# 3. Install make: https://www.gnu.org/software/make/

# Usage:
# 1. Put the public key file in keys directory
# 2. Run the script with the following parameters:
#   $1: <public_key_file>: The public key file name
#   $2: <os>: The operating system of the machine where the script is running: mac or linux 
#   ./check_swap_signature_script.sh <public_key_file> <os>

# Script:
PROJECT_DIR="$(cd ../.. && pwd)"
PUBLIC_KEY_FILE=$1
OS=$2

if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go."
    exit
fi

if ! command -v make o &> /dev/null
then
    echo "Make could not be found. Please install make."
    exit
fi

if ! command -v jq &> /dev/null
then
    echo "jq could not be found. Please install jq."
    exit
fi

#Clean
rm payload.json

# Step 1: Build the exchange-tool
cd $PROJECT_DIR/go
make -f $PROJECT_DIR/go/Makefile build

# Step 2: Fetch Swap Response and put it in payload.json
cd $PROJECT_DIR/scripts/swap
# Replace the below curl command with the actual one to fetch the swap response and put it in payload.json
curl -A "sample-user-agent" --location 'https://exchange-s.exodus.io/v3/ledger/swap' \
--header 'Content-Type: application/json' \
--data '{
    "amountToWallet": "51800",
    "payloadCurrencyFrom": "BTC",
    "payloadCurrencyTo": "USDT",
    "nonce": "1",
    "payoutAddress": "0x9D84548a1454f4D60725EB7f12631F2973423995",
    "refundAddress": "bc1q9mz6pklngy8ze9m2n9gtm297jaed5alvakhkw6"
}' > payload.json
# curl --request POST YOUR_ENDPOINT ... > payload.json

# Step 3: Parse the Response and extract the binaryPayload and signature
PAYLOAD_CONTENT=$(cat payload.json)
BINARY_PAYLOAD_BASE64=$(echo $PAYLOAD_CONTENT | jq -r '.providerSig.payload')
SIGNATURE_BASE64=$(echo $PAYLOAD_CONTENT | jq -r '.providerSig.signature')

echo "binary payload base64: "
echo $BINARY_PAYLOAD_BASE64
echo "signature:" 
echo $SIGNATURE_BASE64

# Step 4: Validate the Response
$PROJECT_DIR/go/bin/exchange-tool-$OS check -c k1 -p $PROJECT_DIR/scripts/swap/keys/$PUBLIC_KEY_FILE $BINARY_PAYLOAD_BASE64 $SIGNATURE_BASE64