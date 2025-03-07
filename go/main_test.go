// SPDX-FileCopyrightText: 2024 Ledger
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os/exec"
	"testing"

	"exchange.ledger.fr/crypto"
	ethereum "github.com/ethereum/go-ethereum/crypto"
)

// const payloadBase64 = "CipiYzFxYXIwc3Jycjd4Zmt2eTVsNjQzbHlkbnc5cmU1OWd0enp3ZjVtZHEaKmJjMXFhcjBzcnJyN3hma3Z5NWw2NDNseWRudzlyZTU5Z3R6endmNHRlcSoqMHhiNzk0ZjVlYTBiYTM5NDk0Y2U4Mzk2MTNmZmZiYTc0Mjc5NTc5MjY4OgNCVENCA0JBVEoCBH5SBgV0-95gAGIgNQrqDJf3R_HQ92CBRhSkdSOAGxrrfQvLuqKk9Gv4GEs="

type consistencyTest struct {
	name, secretFilename, publicFilename string
	curve                                crypto.Curve
	signFormat                           crypto.SignFormat
}

func TestConsistency(t *testing.T) {
	// Given
	inputsTest := []consistencyTest{
		{"K1", "../samples/sample-priv-key-secp256k1.pem", "../samples/sample-pub-key-secp256k1.pem", crypto.K1Curve{}, crypto.Raw},
		{"K1", "../samples/sample-priv-key-secp256k1.pem", "../samples/sample-pub-key-secp256k1.pem", crypto.K1Curve{}, crypto.Jwt},
		{"R1", "../samples/sample-priv-key-secp256r1.pem", "../samples/sample-pub-key-secp256r1.pem", crypto.R1Curve{}, crypto.Raw},
		{"R1", "../samples/sample-priv-key-secp256r1.pem", "../samples/sample-pub-key-secp256r1.pem", crypto.R1Curve{}, crypto.Jwt},
	}
	samplePayloadFilename := "../samples/payload-example.json"

	for _, input := range inputsTest {
		t.Run(input.name, func(t *testing.T) {
			// When
			actual := new(bytes.Buffer)
			RootCmd.SetOut(actual)
			RootCmd.SetErr(actual)
			RootCmd.Flags().Set("curve", input.curve.Flag())
			RootCmd.Flags().Set("format", string(input.signFormat))
			RootCmd.Flags().Set("private", input.secretFilename)
			RootCmd.SetArgs([]string{"generate", samplePayloadFilename})
			RootCmd.Execute()
			fmt.Println("OUTPUT:\n", actual)
			fmt.Println("*********")
			// payload64, sign64 := GenerateProtoAndSign(input.curve, input.signFormat, input.secretFilename, samplePayloadFilename)
			// if payload64 != payloadBase64 {
			// 	t.Fatal("Base64 payload has diverged!")
			// }
			// status := CheckPayload(input.curve, input.signFormat, input.publicFilename, payloadBase64, sign64)

			// // Then
			// if !status.isOk {
			// 	t.Fatal("Unable to check self signature")
			// }
		})
	}
}

func TestCrossCheck(t *testing.T) {
	// Given
	inputsTest := []consistencyTest{
		{"K1", "../example/sample-priv-key-secp256k1.pem", "../example/sample-pub-key-secp256k1.pem", crypto.K1Curve{}, crypto.Raw},
		{"R1", "../example/sample-priv-key-secp256r1.pem", "../example/sample-pub-key-secp256r1.pem", crypto.R1Curve{}, crypto.Raw},
	}
	// samplePayloadFilename := "../example/payload-example.json"

	for _, input := range inputsTest {
		t.Run(input.name, func(t *testing.T) {
			// When
			// privateKey, _ := input.curve.ReadPrivateKey(input.secretFilename)
			// payload64 := convertJsonToBase64(samplePayloadFilename)
			// signature := crypto.SignMessageInDER(payload64, privateKey, crypto.Raw)

			// os.WriteFile("example/encoded_payload.txt", []byte(payload64), 0644)
			// os.WriteFile("example/signature.txt", signature, 0644)

			// opensslCommand(t, input.publicFilename)
		})
	}
}

func TestConvertHexToPub(t *testing.T) {
	pubHex := "044989cad389020fadfb9d7a85d29338a450beec571347d2989fb57b99ecddbc8907cf8c229deee30fb8ac139e978cab8f6efad76bde2a9c6d6710ceda1fe0a4d8"
	k1Curve := crypto.K1Curve{}
	pubKey := k1Curve.ReadHexPublicKey(pubHex)

	der := ethereum.FromECDSAPub(pubKey)
	log.Println("DER format:", hex.EncodeToString(der))

	log.Println("Try PEM")
	// pubHex = "04b2779a60948b55963f86e62cd018d131a02f40d843baeadf356dbc7fe8294bc6a0127c6684693e83c8221cdee13d05fd078d9b68f3f4816e6274f1d5a9ead70e"
	// k1Curve.ConvertHexToPem(pubHex)
}

func opensslCommand(t *testing.T, pubFilename string) {
	cmd := exec.Command("openssl", "dgst", "-sha256", "-verify", pubFilename, "-signature", "../example/signature.txt", ".../example/encoded_payload.txt")
	output, err := cmd.Output()
	if err != nil {
		t.Error("OpenSSL error:", err)
	}
	t.Log("OpenSSL output:", string(output))
}
