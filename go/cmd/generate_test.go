// SPDX-FileCopyrightText: Ledger SAS 2024
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"exchange.ledger.fr/crypto"
	"exchange.ledger.fr/encode"
)

type signTest struct {
	name, secretFilename string
	curve                crypto.Curve
}

func TestSign(t *testing.T) {
	// Given
	inputsTest := []signTest{
		{"K1", "../../samples/sample-priv-key-secp256k1.pem", crypto.K1Curve{}},
		{"R1", "../../samples/sample-priv-key-secp256r1.pem", crypto.R1Curve{}},
	}
	samplePayloadFilename := "../../samples/swap-payload-example.json"

	for _, input := range inputsTest {
		t.Run(input.name, func(t *testing.T) {
			// When
			payload64, sign64 := generateProtoAndSign[encode.SwapDevicePayload](input.curve, crypto.Raw, input.secretFilename, samplePayloadFilename)

			// Then
			if payload64 != payloadBase64 {
				t.Errorf("Payload doesn't have the expected Base64 encoding.\nExpected: %v\nGet: %v", payloadBase64, payload64)
			}
			t.Log("Base64Url signature:", sign64)
		})
	}
}
