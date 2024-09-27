package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestVerifyRSSignature_FromJS_SellPayloadSigned(t *testing.T) {
	// Given
	curve := K1Curve{}
	pubKey := curve.ReadHexPublicKey("0478d5facdae2305f48795d3ce7d9244f5060d2f800901da5746d1f4177ae8d7bbe63f3870efc0d36af8f91962811e1d8d9df91ce3b3ea2cd9f550c7d465f8b7b3")
	encodedPayload, _ := hex.DecodeString("2e436735305a584e305147786c5a47646c6369356d636849445256524947674959616949714d48686b4e6a6b79513249784d7a51324d6a5979526a55344e4551784e304930516a51334d446b314e4455774d5759324e7a4531595467794b674e465656497942416f41454145364944554b366779583930667830506467675559557048556a674273613633304c79377169705052722d42684c")
	payloadSignature, _ := hex.DecodeString("348e938ce75c06da9a27652c2af6146b4581f948c600b4f7093743e8af256d172324131dcc852f068d0bd20afd0c8444bc635e9520d3a29cab7ff7c27dc8b782")

	// When
	isValid := VerifyRSSignature2(pubKey, encodedPayload, payloadSignature)

	// Then
	if !isValid {
		t.Fatalf("Mismatch signature")
	}
}

func TestVerifyDERSignature_FromJS_ApduSigned(t *testing.T) {
	// Given
	pubKey := ledgerPublicKey()
	payload := "0953454c4c5f54455354000478d5facdae2305f48795d3ce7d9244f5060d2f800901da5746d1f4177ae8d7bbe63f3870efc0d36af8f91962811e1d8d9df91ce3b3ea2cd9f550c7d465f8b7b3"
	payloadSignature, _ := hex.DecodeString("30440220471b035b40dafa095d615998c82202b2bd00fb45670b828f1dda3b68e5b24cc3022022a1c64d02b8c14e1e4cc2d05b00234642c11db3d4461ff5366f5af337cf0ced")

	// When
	isValid := VerifyDERSignature(pubKey, payload, payloadSignature)

	// Then
	if !isValid {
		t.Fatalf("Mismatch signature")
	}
}

func TestVerifyRSSignature_FromJS_SimpleExample(t *testing.T) {
	// Given
	curve := K1Curve{}
	pubKey := curve.ReadHexPublicKey("0478d5facdae2305f48795d3ce7d9244f5060d2f800901da5746d1f4177ae8d7bbe63f3870efc0d36af8f91962811e1d8d9df91ce3b3ea2cd9f550c7d465f8b7b3")
	payload := "Something important to cipher"
	payloadSignature, _ := hex.DecodeString("c47cfb2ed7f040f7701e62a422d2fa1b7eb4d6285c4082688ec69313ff2f0890252aba494b865173154b574ebc81b887a0234ae8cd4939f34def777ae94680dc")

	// When
	isValid := VerifyRSSignature(pubKey, payload, payloadSignature, Raw)

	// Then
	if !isValid {
		t.Fatalf("Mismatch signature")
	}
}

func TestVerifyRSSignature_PayloadSigned(t *testing.T) {
	type signEncoding string
	const (
		hexEncode    signEncoding = "hex"
		base64Encode signEncoding = "base64"
	)
	testCases := []struct {
		payload      string
		signature    string
		signEncoding signEncoding
		pubkeyFile   string
		curve        Curve
		signFormat   SignFormat
	}{
		{
			"CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHESA0JUQxoCBH4iKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoDRVVSMgwKCAAAAAAAAAAMEAI6IDUK6gyX90fx0PdggUYUpHUjgBsa630Ly7qipPRr-BhL",
			"accbd06d0ea964ce104c5921a81bfd162a491eb09f6cb4730c6b0cc0cf4206db2ed76b84b6fdefae9a3818b39bbbb2ae45d609ffb471a58b6d8bbb727a12a216",
			hexEncode,
			"../../samples/sample-pub-key-secp256k1.pem",
			K1Curve{},
			Raw,
		},
		{
			"CjBVUUFidnMydENuc1RXeENaWDdKVy1kcWxrMHZNOHhfbThhSnFGNHd3UldHdFRFWkQaMFVRQWJ2czJ0Q25zVFd4Q1pYN0pXLWRxbGswdk04eF9tOGFKcUY0d3dSV0d0VEVaRCoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNUT05CA0JBVEoCBH5SBgV0-95gAGIgNQrqDJf3R_HQ92CBRhSkdSOAGxrrfQvLuqKk9Gv4GEs=",
			"1I1Tml0c9zluE0lBwuEDonMQ4KREF8WeynlkYQv_cyVMj_svljqOfkCs6EXLEauPM9y3X1-PTryk8YaDZ0H3cw==",
			base64Encode,
			"../../samples/sample-pub-key-secp256r1.pem",
			R1Curve{},
			Jwt,
		},
		{
			"CjBVUUFidnMydENuc1RXeENaWDdKVy1kcWxrMHZNOHhfbThhSnFGNHd3UldHdFRFWkQaMFVRQWJ2czJ0Q25zVFd4Q1pYN0pXLWRxbGswdk04eF9tOGFKcUY0d3dSV0d0VEVaRCoqMHg2NmM0MzcxYUU4RkZlRDJlYzFjMkVCYmJjQ2ZiN0U0OTQxODFFMUUzOgNUT05CA0VUSEoCBH5SBgV0-95gAGIgNQrqDJf3R_HQ92CBRhSkdSOAGxrrfQvLuqKk9Gv4GEs=",
			"cF85mRKNOECR5iVFSRTi-YQYGSdv34d3aVZabUqVNbJ_qo4SKaAWkyWgMCdU5MkZOtPczRmXIIMJeVvwe3yTUw==",
			base64Encode,
			"../../samples/sample-pub-key-secp256r1.pem",
			R1Curve{},
			Jwt,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprint("", idx), func(t *testing.T) {
			// Given
			pubKey := tc.curve.ReadPublicKeyFile(tc.pubkeyFile)
			encodedPayload := tc.payload
			var payloadSignature []byte
			if tc.signEncoding == hexEncode {
				payloadSignature, _ = hex.DecodeString(tc.signature)
			} else {
				payloadSignature, _ = base64.URLEncoding.DecodeString(tc.signature)
			}

			// When
			isValid := VerifyRSSignature(pubKey, encodedPayload, payloadSignature, tc.signFormat)

			// Then
			if !isValid {
				t.Fatalf("Mismatch signature")
			}
		})
	}
}
