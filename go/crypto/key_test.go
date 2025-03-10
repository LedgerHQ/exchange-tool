// SPDX-FileCopyrightText: Ledger SAS 2024
//
// SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"encoding/base64"
	"testing"
)

func TestReadPrivateFileKey(t *testing.T) {
	filename := "../../samples/sample-priv-key-secp256k1.pem"

	curve := K1Curve{}

	_, err := curve.ReadPrivateKey(filename)
	if err != nil {
		t.Error("Unable to read secret file:", err)
	}
}

func TestReadPublicFileKey(t *testing.T) {
	filename := "../../samples/sample-pub-key-secp256k1.pem"

	curve := K1Curve{}

	curve.ReadPublicKeyFile(filename)
}

func TestK1PartnerKey(t *testing.T) {
	pubHex := "044989cad389020fadfb9d7a85d29338a450beec571347d2989fb57b99ecddbc8907cf8c229deee30fb8ac139e978cab8f6efad76bde2a9c6d6710ceda1fe0a4d8"
	payload := "CioweEM5RUY3MDYxMjcxQWMyNGM1MzI5MzkxNGM4QjU5MkU4RDYxMDNmQmUaKjB4OTQ3RTU3NjY5ZjhDOGMyYjE2M0I0MjA1QjhEOUVCMzE4M2EwNzhBYioqMHg5NDdFNTc2NjlmOEM4YzJiMTYzQjQyMDVCOEQ5RUIzMTgzYTA3OEFiOgNidGNCA2J0Y0oQAAAAAAAYkRAAAAAAAAAAAFIQyPNLSBuoAAAAAAAAAAAAAFoBMWIQAAAAAAAAAAEAAAAAAAAAAA"
	signature, _ := base64.RawURLEncoding.DecodeString("C_wzKkIF9VSGsyU181_F1sZysPO_sQJbMXX8gzaN9ibgkA2eeAuJBXz-SwoMKU-ddOBSoiTX7_Yv3XKArzUYjg")

	curve := K1Curve{}
	key := curve.ReadHexPublicKey(pubHex)

	result := VerifyRSSignature(key, payload, signature, Jwt)

	if result == false {
		t.Fatal("Unable to verify signature")
	}
}

func TestK1PartnerFileKey(t *testing.T) {
	pubFile := "../../samples/sample-pub-key-secp256k1.pem"
	payload := "CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHESA0JUQxoCBH4iKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoDRVVSMgwKCAAAAAAAAAAMEAI6IDUK6gyX90fx0PdggUYUpHUjgBsa630Ly7qipPRr-BhL"
	signature, _ := base64.URLEncoding.DecodeString("j2kXZWAAIk6rIQbccYbZ0BgFaGrqedcWKoAyekqLDSTdej2FEfF5GYDWpzxV1cn8Y4bJiL1xM-geWVZkqok2eA==")

	curve := K1Curve{}
	key := curve.ReadPublicKeyFile(pubFile)

	result := VerifyRSSignature(key, payload, signature, Raw)

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

	result := VerifyRSSignature(key, payload, signature, Jwt)

	if result == false {
		t.Fatal("Unable to verify signature")
	}
}

func TestR1PartnerFileKey_Raw(t *testing.T) {
	pubFile := "../../samples/sample-pub-key-secp256r1.pem"
	payload := "CjBVUUFidnMydENuc1RXeENaWDdKVy1kcWxrMHZNOHhfbThhSnFGNHd3UldHdFRFWkQaMFVRQWJ2czJ0Q25zVFd4Q1pYN0pXLWRxbGswdk04eF9tOGFKcUY0d3dSV0d0VEVaRCoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNUT05CA0JBVEoCBH5SBgV0-95gAGIgNQrqDJf3R_HQ92CBRhSkdSOAGxrrfQvLuqKk9Gv4GEs="
	signature, _ := base64.URLEncoding.DecodeString("w90hVJjj0Pquqq0CLmtq6sFNAGrMRWMmNKn3OqRkZzbj6fpIQMO1pN7d70sSL4DrLEXO9Hacvi3tib5D-p5uEA==")

	curve := R1Curve{}
	key := curve.ReadPublicKeyFile(pubFile)

	result := VerifyRSSignature(key, payload, signature, Raw)

	if result == false {
		t.Fatal("Unable to verify signature")
	}
}

func TestR1PartnerFileKey_Jwt(t *testing.T) {
	pubFile := "../../samples/sample-pub-key-secp256r1.pem"
	payload := "CjBVUUFidnMydENuc1RXeENaWDdKVy1kcWxrMHZNOHhfbThhSnFGNHd3UldHdFRFWkQaMFVRQWJ2czJ0Q25zVFd4Q1pYN0pXLWRxbGswdk04eF9tOGFKcUY0d3dSV0d0VEVaRCoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNUT05CA0JBVEoCBH5SBgV0-95gAGIgNQrqDJf3R_HQ92CBRhSkdSOAGxrrfQvLuqKk9Gv4GEs="
	signature, _ := base64.URLEncoding.DecodeString("s86hbFOoRDJsA7l62q9k10dfToyJPZSuuZZtOQbZU07LoqEVjCklASXqbFMJeU3pWXtQQiNu96KoH7xjhhRTqg==")

	curve := R1Curve{}
	key := curve.ReadPublicKeyFile(pubFile)

	result := VerifyRSSignature(key, payload, signature, Jwt)

	if result == false {
		t.Fatal("Unable to verify signature")
	}
}
