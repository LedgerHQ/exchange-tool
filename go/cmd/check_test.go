// SPDX-FileCopyrightText: 2024 Ledger
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"exchange.ledger.fr/crypto"
)

type checkTest struct {
	name, publicFilename, expectedSignBase64 string
	curve                                    crypto.Curve
}

const payloadBase64 = "CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoCBH5SBgV0-95gAGIgNQrqDJf3R_HQ92CBRhSkdSOAGxrrfQvLuqKk9Gv4GEs="

func TestCheck_WithPubkeyFilename(t *testing.T) {
	// Given
	inputsTest := []checkTest{
		{"K1", "../../samples/sample-pub-key-secp256k1.pem", "HPqmIgJSF3sOIzyo2RKv_MpjLCm4gE9h4Xg6zoyS80TF8GfOxu47EtYMTRHktD8J00-VwznaM7NERuHiEZv8sg==", crypto.K1Curve{}},
		{"R1", "../../samples/sample-pub-key-secp256r1.pem", "PQ1upD7qyswX-GK5Om4nfH4toiaDvFA0fWizIoDIRpMXfQm4si8-62tvShlCSHEQ3nndsuVGpf090e3YssUvDw==", crypto.R1Curve{}},
	}
	for _, input := range inputsTest {
		t.Run(input.name, func(t *testing.T) {
			status := checkPayloadWithKeyFile(input.curve, crypto.Raw, input.publicFilename, payloadBase64, input.expectedSignBase64)

			if !status.isOk {
				t.Fatal("Signature is not the one expected (see Python tests)")
			}
		})
	}
}

func TestCheck_WithPubkeyHexValue(t *testing.T) {
	// Given
	inputsTest := []checkTest{
		{"K1", "0478d5facdae2305f48795d3ce7d9244f5060d2f800901da5746d1f4177ae8d7bbe63f3870efc0d36af8f91962811e1d8d9df91ce3b3ea2cd9f550c7d465f8b7b3", "HPqmIgJSF3sOIzyo2RKv_MpjLCm4gE9h4Xg6zoyS80TF8GfOxu47EtYMTRHktD8J00-VwznaM7NERuHiEZv8sg==", crypto.K1Curve{}},
		{"R1", "047c13debdb9e1afac5a82bbe78da6dca98a2e59af8a9a3f827acb8ed325c79a6c5749eff65d4e3470bd2995def771f45426c58eada7227d536feeffb58fa71c24", "PQ1upD7qyswX-GK5Om4nfH4toiaDvFA0fWizIoDIRpMXfQm4si8-62tvShlCSHEQ3nndsuVGpf090e3YssUvDw==", crypto.R1Curve{}},
	}
	for _, input := range inputsTest {
		t.Run(input.name, func(t *testing.T) {
			status := checkPayloadWithKeyHex(input.curve, crypto.Raw, input.publicFilename, payloadBase64, input.expectedSignBase64)

			if !status.isOk {
				t.Fatal("Signature is not the one expected (see Python tests)")
			}
		})
	}
}
