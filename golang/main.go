/*
 *  Licensed to the Apache Software Foundation (ASF) under one
 *  or more contributor license agreements.  See the NOTICE file
 *  distributed with this work for additional information
 *  regarding copyright ownership.  The ASF licenses this file
 *  to you under the Apache License, Version 2.0 (the
 *  "License"); you may not use this file except in compliance
 *  with the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing,
 *  software distributed under the License is distributed on an
 *  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 *  KIND, either express or implied.  See the License for the
 *  specific language governing permissions and limitations
 *  under the License.
 */

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"swap.ledger.fr/crypto"
	"swap.ledger.fr/encode"
)

type checkStatus struct {
	isOk         bool
	base64Format encode.Base64Format
}

func main() {
	fmt.Println("=== Ledger's Swap Protobuf Utils ===")
	// log.SetOutput(io.Discard)

	params := ReadParams(os.Args)

	switch params.command {
	case GenerateCmd:
		fmt.Println("*** Generate Swap Proto ***")
		marshalledPayload := func() []byte {
			payloadJson := encode.ConvertFileToDevicePayload[encode.SwapDevicePayload](params.payloadFilename)
			return encode.EncodeDevicePaylod(payloadJson)
		}
		payload64, sign64 := GenerateProtoAndSign(params.curve, params.signFormat, params.pemFile, marshalledPayload)

		fmt.Println("--> Result payload base64:", payload64)
		fmt.Println("--> Result signature base64:", sign64)

		hash := sha256.New()
		hash.Write([]byte(payload64))
		fmt.Println("--> Sha256 of base64 payload with no prefix:", hex.EncodeToString(hash.Sum(nil)))
		hash = sha256.New()
		hash.Write([]byte("." + payload64))
		fmt.Println("--> Sha256 of base64 payload with prefix:", hex.EncodeToString(hash.Sum(nil)))
		fmt.Println("--> base64 payload size:", len(payload64))

	case GenerateSellCmd:
		fmt.Println("*** Generate Sell Proto ***")
		marshalledPayload := func() []byte {
			payloadJson := encode.ConvertFileToDevicePayload[encode.SellDevicePayload](params.payloadFilename)
			return encode.EncodeDevicePaylod(payloadJson)
		}
		payload64, sign64 := GenerateProtoAndSign(params.curve, params.signFormat, params.pemFile, marshalledPayload)

		fmt.Println("--> Result payload base64:", payload64)
		fmt.Println("--> Result signature base64:", sign64)

		hash := sha256.New()
		hash.Write([]byte(payload64))
		fmt.Println("--> Sha256 of base64 payload with no prefix:", hex.EncodeToString(hash.Sum(nil)))
		hash = sha256.New()
		hash.Write([]byte("." + payload64))
		fmt.Println("--> Sha256 of base64 payload with prefix:", hex.EncodeToString(hash.Sum(nil)))
		fmt.Println("--> base64 payload size:", len(payload64))

	case CheckCmd:
		fmt.Println("*** Check signature ***")
		status := CheckPayload(params.curve, params.signFormat, params.pemFile, params.payloadBase64, params.signatureBase64)

		if status.isOk {
			fmt.Println("--> Payload is CORRECTLY signed")
		} else {
			fmt.Println("--> Payload is NOT CORRECTLY signed")
		}

	case ReadProtobufCmd:
		fmt.Println("*** Extract protobuf file info ***")
		info, format := ReadPayload(params.payloadBase64)

		fmt.Print("--> Info (", format, "):\n{\n", info, "}\n\n")

	case SigToHex:
		fmt.Println("*** Read pubkey to hex ***")
		curve := crypto.K1Curve{}
		hexValue := curve.ConvertPrivPEMtoHexKey(params.pemFile)

		fmt.Println("--> Hex value:", hexValue)

	case SignProviderInfo:
		fmt.Println("*** Read pubkey to hex ***")
		curve := crypto.K1Curve{}
		pubKey := curve.ReadPublicKey(params.pemFile)
		signature := crypto.SignProviderInfo(params.providerName, pubKey)

		fmt.Println("--> Signature value:", signature)
	case GenerateCalInfo:
		fmt.Println("*** Generate CAL format info ***")
		curve := params.curve
		pubKey := curve.ReadPublicKey(params.pemFile)
		signature := crypto.SignProviderInfo(params.providerName, pubKey)
		calInfo := encode.CalFormatProviderInfo(params.providerName, curve.Name(), curve.ConvertPEMtoHexKey(params.pemFile), signature)

		fmt.Println("--> CAL info:", calInfo.String())
		fmt.Println("--> CAL info:", calInfo.Pretty())
	}
}

type marshalFile func() []byte

func GenerateProtoAndSign(curve crypto.Curve, signFormat crypto.SignFormat, privFilename string, fnMarshalledFile marshalFile) (payload64 string, sign64 string) {
	privateKey, _ := curve.ReadPrivateKey(privFilename)
	payload64 = fileToBase64(fnMarshalledFile)
	sign64 = crypto.SignMessageInRS(payload64, privateKey, signFormat)
	return
}

func fileToBase64(fnMarshalledFile marshalFile) string {
	payloadMarshalled := fnMarshalledFile()
	return encode.EncodeBase64(payloadMarshalled)
}

// Check provided base64 URL encoded payload (which must be a binary protobuf) that match the signature
func CheckPayload(curve crypto.Curve, signFormat crypto.SignFormat, pubFilename, payload, signature string) checkStatus {
	publicKey := curve.ReadPublicKey(pubFilename)

	signatureByte, format := encode.CascadeDecodeBase64(signature)
	status := checkStatus{
		base64Format: format,
		isOk:         true,
	}

	status.isOk = crypto.VerifySignature(publicKey, payload, signatureByte, signFormat)

	return status
}

// Base64 decode payload and extract its info from binary result
func ReadPayload(payload string) (string, encode.Base64Format) {
	payloadByte, format := encode.CascadeDecodeBase64(payload)
	payloadJson := encode.DecodeSwapProtobuf(payloadByte)

	return payloadJson.String(), format
}
