package cmd

import (
	"testing"

	"exchange.ledger.fr/crypto"
)

func TestCal(t *testing.T) {
	calInfo := generateCal(crypto.R1Curve{}, "../../../../coinify-pubkey.pem", "Coinify", 1, "swap")

	if calInfo.Name != "Coinify" {
		t.Fatalf("Wrong CAL info")
	}
	if calInfo.PublicKey.Data != "044f22668f5f321d3784266c932a2a3141c3ec196ddd51f42cf975267eda23d3a8b02170e4c5c70536e7d03ba4e66ee3e1f9d65e772d3217871a830a7cf60da366" {
		t.Fatalf("Wrong pubkey conversion")
	}
	const expectedApdu = "07436f696e696679044f22668f5f321d3784266c932a2a3141c3ec196ddd51f42cf975267eda23d3a8b02170e4c5c70536e7d03ba4e66ee3e1f9d65e772d3217871a830a7cf60da366"
	if calInfo.Apdu != expectedApdu {
		t.Fatalf("Wrong APDU.\nGot:\n%s\nExpected:\n%s\n", calInfo.Apdu, expectedApdu)
	}
	const expectedSignature = "30450221008e8b2172ddd48e196dbff81ebe8aebc4ec0988f72de1b02da202f3a8d8f33f9c02205273c1d426aeb460fd36a27696aafda68bdff4139886f29e558895ea85527749"
	if calInfo.Signature != expectedSignature {
		t.Fatalf("Wrong signature.\nGot:\n%s\nExpected:\n%s\n", calInfo.Signature, expectedSignature)
	}
}
