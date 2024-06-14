package crypto

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"log"
	"strconv"
)

const LEDGER_FAKE_PRIVATE_KEY = "b1ed47ef58f782e2bc4d5abe70ef66d9009c2957967017054470e0f3e10f5833"

var LEDGER_FAKE_PUBLIC_KEY = []byte{0x4,
	0x20, 0xDA, 0x62, 0x00, 0x3C, 0x0C, 0xE0, 0x97,
	0xE3, 0x36, 0x44, 0xA1, 0x0F, 0xE4, 0xC3, 0x04,
	0x54, 0x06, 0x9A, 0x44, 0x54, 0xF0, 0xFA, 0x9D,
	0x4E, 0x84, 0xF4, 0x50, 0x91, 0x42, 0x9B, 0x52,
	0x20, 0xAF, 0x9E, 0x35, 0xC0, 0xB2, 0xD9, 0x28,
	0x93, 0x80, 0x13, 0x73, 0x07, 0xDE, 0x4D, 0xD1,
	0xD4, 0x18, 0x42, 0x8C, 0xF2, 0x1A, 0x93, 0xB3,
	0x35, 0x61, 0xBB, 0x09, 0xD8, 0x8F, 0xE5, 0x79}

func SignProviderInfo(providerName string, pubKey string, curve Curve, version uint) (string, string) {
	privKey := ledgerPrivateKey()
	log.Println("Length:",
		len(providerName), " - ",
		strconv.Itoa(len(providerName)), " - ",
		[]byte(strconv.Itoa(len(providerName))), " - ",
		binary.BigEndian.AppendUint16(make([]byte, 0), uint16(len(providerName))), " - ",
	)

	var msg string
	if version > 1 {
		// msg = hex.EncodeToString([]byte(strconv.Itoa(len(providerName)))) + hex.EncodeToString([]byte(providerName)) + "00" + hex.EncodeToString(pubKey.X.Bytes())
		msg = hex.EncodeToString(
			append(
				append(
					append([]byte(strconv.Itoa(len(providerName))), []byte(providerName)...),
					curve.Code()...,
				),
				[]byte(pubKey)...,
			),
		)
	} else {
		// Remove the first byte, due to conversion of `providerName` length to 16bits int.
		// This is a constaint to the APDU format, therefore `providerName` length can't be > 255 chars.
		msg = hex.EncodeToString(
			append(binary.BigEndian.AppendUint16(make([]byte, 0), uint16(len(providerName))), []byte(providerName)...),
		)[2:] + pubKey // As pubkey is expected to already be in hex format, we just concat the 2 strings.
	}
	// log.Println("APDU signed:", msg)
	// log.Println("APDU signed:", msg[2:])
	signature := SignMessageInDER(msg, privKey, Raw)
	return msg, hex.EncodeToString(signature)
}

func ledgerPrivateKey() *ecdsa.PrivateKey {
	curve := K1Curve{}
	privKey, err := curve.ReadPrivateKeyFromHex(LEDGER_FAKE_PRIVATE_KEY)
	if err != nil {
		log.Fatalln("Error while reading Ledger's key:", privKey)
	}

	return privKey
}
