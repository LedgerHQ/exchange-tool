package crypto

import (
	"encoding/base64"
	"testing"
)

func TestRead(t *testing.T) {
	filename := "../../samples/sample-priv-key-secp256k1.pem"

	curve := K1Curve{}

	_, err := curve.ReadPrivateKey(filename)
	if err != nil {
		t.Error("Unable to read secret file:", err)
	}
}

func TestK1PartnerKey(t *testing.T) {
	pubHex := "044989cad389020fadfb9d7a85d29338a450beec571347d2989fb57b99ecddbc8907cf8c229deee30fb8ac139e978cab8f6efad76bde2a9c6d6710ceda1fe0a4d8"
	payload := "CioweEM5RUY3MDYxMjcxQWMyNGM1MzI5MzkxNGM4QjU5MkU4RDYxMDNmQmUaKjB4OTQ3RTU3NjY5ZjhDOGMyYjE2M0I0MjA1QjhEOUVCMzE4M2EwNzhBYioqMHg5NDdFNTc2NjlmOEM4YzJiMTYzQjQyMDVCOEQ5RUIzMTgzYTA3OEFiOgNidGNCA2J0Y0oQAAAAAAAYkRAAAAAAAAAAAFIQyPNLSBuoAAAAAAAAAAAAAFoBMWIQAAAAAAAAAAEAAAAAAAAAAA"
	signature, _ := base64.RawURLEncoding.DecodeString("C_wzKkIF9VSGsyU181_F1sZysPO_sQJbMXX8gzaN9ibgkA2eeAuJBXz-SwoMKU-ddOBSoiTX7_Yv3XKArzUYjg")

	curve := K1Curve{}
	key := curve.ReadHexPublicKey(pubHex)

	result := VerifySignature(key, payload, signature, Jwt)

	if result == false {
		t.Fatal("Unable to verify signature")
	}
}

func TestR1PartnerKey(t *testing.T) {
	pubHex := "044f22668f5f321d3784266c932a2a3141c3ec196ddd51f42cf975267eda23d3a8b02170e4c5c70536e7d03ba4e66ee3e1f9d65e772d3217871a830a7cf60da366"
	payload := "CiRzdGVwaGFuZS5wcm9oYXN6a2ErY29pbmlmeUBsZWRnZXIuZnISA0JUQxoDB6EgIiMyTXhKdUw2ZVByck5IOFdYWlRZN2RCSnNZWmduSFM3UWQ1NSoDRVVSMgYKAnQsEAI6IDUK6gyX90fx0PdggUYUpHUjgBsa630Ly7qipPRr-BhL"
	signature, _ := base64.RawURLEncoding.DecodeString("u63xyhIAlgIKj0bfatpXGfoCxG0OfkeplLX9tVPia65mxBubSzj31MS-ohJvexi990b4gjgkUF1fORbUe9UdmA")

	curve := R1Curve{}
	key := curve.ReadHexPublicKey(pubHex)

	result := VerifySignature(key, payload, signature, Jwt)

	if result == false {
		t.Fatal("Unable to verify signature")
	}
}
