package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"strconv"
)

const LEDGER_FAKE_PRIVATE_KEY = "b1ed47ef58f782e2bc4d5abe70ef66d9009c2957967017054470e0f3e10f5833"

func SignProviderInfo(providerName string, pubKey *ecdsa.PublicKey) string {
	privKey := ledgerPrivateKey()
	log.Println("Length:", len(providerName), " - ", strconv.Itoa(len(providerName)), " - ", []byte(strconv.Itoa(len(providerName))))
	// msg := hex.EncodeToString([]byte(strconv.Itoa(len(providerName)))) + hex.EncodeToString([]byte(providerName)) + "00" + hex.EncodeToString(pubKey.X.Bytes())
	msg := hex.EncodeToString(
		append(
			append(
				append([]byte(strconv.Itoa(len(providerName))), []byte(providerName)...),
				"00"...,
			),
			pubKey.X.Bytes()...,
		),
	)
	log.Println("APDU signed:", msg)
	signature := SignMessageInDER(msg, privKey, Raw)
	return hex.EncodeToString(signature)
}

func ledgerPrivateKey() *ecdsa.PrivateKey {
	curve := K1Curve{}
	privKey, err := curve.ReadPrivateKeyFromHex(LEDGER_FAKE_PRIVATE_KEY)
	if err != nil {
		log.Fatalln("Error while reading Ledger's key:", privKey)
	}

	return privKey
}
