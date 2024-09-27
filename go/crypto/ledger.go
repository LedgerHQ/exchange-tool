package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
)

const LEDGER_FAKE_PRIVATE_KEY = "b1ed47ef58f782e2bc4d5abe70ef66d9009c2957967017054470e0f3e10f5833"

var LEDGER_FAKE_PUBLIC_KEY = hex.EncodeToString(
	[]byte{0x4,
		0x20, 0xDA, 0x62, 0x00, 0x3C, 0x0C, 0xE0, 0x97,
		0xE3, 0x36, 0x44, 0xA1, 0x0F, 0xE4, 0xC3, 0x04,
		0x54, 0x06, 0x9A, 0x44, 0x54, 0xF0, 0xFA, 0x9D,
		0x4E, 0x84, 0xF4, 0x50, 0x91, 0x42, 0x9B, 0x52,
		0x20, 0xAF, 0x9E, 0x35, 0xC0, 0xB2, 0xD9, 0x28,
		0x93, 0x80, 0x13, 0x73, 0x07, 0xDE, 0x4D, 0xD1,
		0xD4, 0x18, 0x42, 0x8C, 0xF2, 0x1A, 0x93, 0xB3,
		0x35, 0x61, 0xBB, 0x09, 0xD8, 0x8F, 0xE5, 0x79},
)

func SignProviderInfo(providerName string, pubKey string, curve Curve, version uint) (string, string) {
	var msg string
	if version > 1 {
		msg = fmt.Sprintf("%02x", len(providerName)) + hex.EncodeToString([]byte(providerName)) + hex.EncodeToString(curve.Code()) + pubKey

	} else {
		msg = fmt.Sprintf("%02x", len(providerName)) + hex.EncodeToString([]byte(providerName)) + pubKey
	}
	log.Println("APDU to sign:", msg)

	privKey := ledgerPrivateKey()
	msgBytes, _ := hex.DecodeString(msg)
	signature := SignMessageInDER(msgBytes, privKey)
	return msg, hex.EncodeToString(signature)
}

type CoinConfig struct {
	ticker  string
	appName string
	config  []byte
}

// Generate CAL info for currency
// Returns an hash and its signature
func GenerateCoinConfig(coin CoinConfig) (string, string) {
	/*
			const payload = Buffer.concat([Buffer.from([ticker.length]), Buffer.from(ticker),
		                                   Buffer.from([applicationName.length]), Buffer.from(applicationName),
		                                   Buffer.from([coinConfig.length]), coinConfig]);
			const hash = Buffer.from(sha256.sha256.array(payload));
			const signature = secp256k1.sign(hash, ledgerPrivateKey).signature;
			const der = secp256k1.signatureExport(signature);
			return { "coinConfig": payload, "signature": der };
	*/
	/*
		"bitcoin",
		"0342544307426974636f696e00",
		"3045022100cb174382302219dca359c0a4d457b2569e31a06b2c25c0088a2bd3fd6c04386a02202c6d0a5b924a414621067e316f021aa13aa5b2eee2bf36ea3cfddebc053b201b"
	*/

	// serialized_config
	serialized := fmt.Sprintf("%02x", len(coin.ticker)) + hex.EncodeToString([]byte(coin.ticker)) +
		fmt.Sprintf("%02x", len(coin.appName)) + hex.EncodeToString([]byte(coin.appName)) +
		fmt.Sprintf("%02x", len(coin.config)) + hex.EncodeToString(coin.config)

	// Sign process
	configByte, _ := hex.DecodeString(serialized)
	signature := SignMessageInDER(configByte, ledgerPrivateKey())

	return serialized, hex.EncodeToString(signature)
}

func ledgerPrivateKey() *ecdsa.PrivateKey {
	curve := K1Curve{}
	privKey, err := curve.ReadPrivateKeyFromHex(LEDGER_FAKE_PRIVATE_KEY)
	if err != nil {
		log.Fatalln("Error while reading Ledger's key:", privKey)
	}

	return privKey
}

func ledgerPublicKey() *ecdsa.PublicKey {
	curve := K1Curve{}
	pubKey := curve.ReadHexPublicKey(LEDGER_FAKE_PUBLIC_KEY)

	return pubKey
}
