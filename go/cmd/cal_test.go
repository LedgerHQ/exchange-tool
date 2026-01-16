// SPDX-FileCopyrightText: Ledger SAS 2024
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"testing"

	"exchange.ledger.fr/crypto"
)

func TestCal(t *testing.T) {
	calInfo := generateCal(crypto.R1Curve{}, "../../samples/sample-pub-key-secp256r1.pem", "Coinify", 1, "swap")

	if calInfo.Name != "Coinify" {
		t.Fatalf("Wrong CAL info")
	}
	if calInfo.PublicKey.Data != "047c13debdb9e1afac5a82bbe78da6dca98a2e59af8a9a3f827acb8ed325c79a6c5749eff65d4e3470bd2995def771f45426c58eada7227d536feeffb58fa71c24" {
		t.Fatalf("Wrong pubkey conversion. Got: %s", calInfo.PublicKey.Data)
	}
	const expectedApdu = "07436f696e696679047c13debdb9e1afac5a82bbe78da6dca98a2e59af8a9a3f827acb8ed325c79a6c5749eff65d4e3470bd2995def771f45426c58eada7227d536feeffb58fa71c24"
	if calInfo.Apdu != expectedApdu {
		t.Fatalf("Wrong APDU.\nGot:\n%s\nExpected:\n%s\n", calInfo.Apdu, expectedApdu)
	}
}
