// SPDX-FileCopyrightText: 2024 Ledger
//
// SPDX-License-Identifier: Apache-2.0

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

type SubConfig struct {
	Ticker    string
	Magnitude uint8
	ChainId   uint16
}
type CoinConfig struct {
	Ticker    string
	AppName   string
	SubConfig *SubConfig
}

// Generate CAL info for currency
// Returns an hash and its signature
func GenerateCoinConfig(coin CoinConfig) (string, string) {
	// SubConfig
	// if coin.appName == "Ethereum" && coin.subConfig.chainId == 0 {
	// 	log.Fatalln("ChainId required with Ethereum app")
	// }

	subConfig := make([]byte, 0)
	if coin.SubConfig != nil {
		if coin.SubConfig.ChainId == 0 {
			subConfig, _ = hex.DecodeString(
				fmt.Sprintf("%02x", len(coin.SubConfig.Ticker)) + hex.EncodeToString([]byte(coin.SubConfig.Ticker)) + fmt.Sprintf("%02x", coin.SubConfig.Magnitude),
			)
		} else {
			subConfig, _ = hex.DecodeString(
				fmt.Sprintf("%02x", len(coin.SubConfig.Ticker)) + hex.EncodeToString([]byte(coin.SubConfig.Ticker)) + fmt.Sprintf("%02x", coin.SubConfig.Magnitude) + fmt.Sprintf("%016x", coin.SubConfig.ChainId),
			)
		}
	}

	// serialized_config
	serialized := fmt.Sprintf("%02x", len(coin.Ticker)) + hex.EncodeToString([]byte(coin.Ticker)) +
		fmt.Sprintf("%02x", len(coin.AppName)) + hex.EncodeToString([]byte(coin.AppName)) +
		fmt.Sprintf("%02x", len(subConfig)) + hex.EncodeToString(subConfig)

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
