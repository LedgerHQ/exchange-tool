// SPDX-FileCopyrightText: 2024 Ledger
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"exchange.ledger.fr/crypto"
)

type Params struct {
	command         Command
	curve           crypto.Curve
	pemFile         string
	payloadFilename string
	payloadBase64   string
	signatureBase64 string
	payloadHeader   string
	signFormat      crypto.SignFormat
	providerName    string
}

type Command string

const (
	GenerateCmd         Command = "generate"
	GenerateSellCmd     Command = "generate-sell"
	CheckCmd            Command = "check"
	ReadProtobufCmd     Command = "read"
	ReadSellProtobufCmd Command = "read-sell"
	SigToHex            Command = "hex"
	SignProviderInfo    Command = "sign"
	GenerateCalInfo     Command = "cal"
)

func ReadParams(args []string) *Params {
	if len(args) < 2 {
		printUsage()
		os.Exit(0)
	}

	switch args[1] {
	case string(GenerateCmd), string(GenerateSellCmd):
		signFormat := crypto.Jwt
		if len(args) > 5 {
			signFormat = parseSignFormat(args[5])
		}

		var cmd Command
		if args[1] == string(GenerateCmd) {
			cmd = GenerateCmd
		} else {
			cmd = GenerateSellCmd
		}

		return &Params{
			command:         cmd,
			curve:           parseCurve(args[2]),
			pemFile:         args[3],
			payloadFilename: args[4],
			signFormat:      signFormat,
		}
	case string(CheckCmd):
		signFormat := crypto.Jwt
		if len(args) > 6 {
			signFormat = parseSignFormat(args[6])
		}
		params := &Params{
			command:         CheckCmd,
			curve:           parseCurve(args[2]),
			pemFile:         args[3],
			signatureBase64: args[5], // Expected in base64
			signFormat:      signFormat,
		}
		params.fillJWTPayload(args[4])
		return params
	case string(ReadProtobufCmd):
		return &Params{
			command:       ReadProtobufCmd,
			payloadBase64: args[2], // Expected in base64
		}
	case string(SigToHex):
		return &Params{
			command: SigToHex,
			pemFile: args[2],
		}
	case string(SignProviderInfo):
		return &Params{
			command:      SignProviderInfo,
			providerName: args[2],
			pemFile:      args[3],
		}
	case string(GenerateCalInfo):
		return &Params{
			command:      GenerateCalInfo,
			providerName: args[2],
			curve:        parseCurve(args[3]),
			pemFile:      args[4],
		}
	default:
		fmt.Println("Unknown command!")
		printUsage()
		os.Exit(0)
	}
	return nil
}

func parseCurve(input string) crypto.Curve {
	switch input {
	case "k1":
		return crypto.K1Curve{}
	case "r1":
		return crypto.R1Curve{}
	default:
		log.Fatalln("Unknown curve!")
		return nil
	}
}

func parseSignFormat(input string) crypto.SignFormat {
	switch input {
	case "raw":
		return crypto.Raw
	case "jwt":
		return crypto.Jwt
	default:
		log.Fatalln("Unknown signFormat!")
		return ""
	}
}

// JWT format: <Base64Url(HEADER)>.<Base64Url(PAYLOAD)>.<Base64Url(SIGNATURE)>
// where: SIGNATURE is computed on <Base64Url(HEADER)>.<Base64Url(PAYLOAD)>
func (p *Params) fillJWTPayload(jwtPayload string) {
	jwtTokenized := strings.Split(jwtPayload, ".")
	switch len(jwtTokenized) {
	case 3:
		p.payloadHeader = jwtTokenized[0]
		p.payloadBase64 = jwtTokenized[1]
		p.signatureBase64 = jwtTokenized[2]
	case 2:
		p.payloadHeader = ""
		p.payloadBase64 = jwtTokenized[0]
		p.signatureBase64 = jwtTokenized[1]
	case 1:
		p.payloadHeader = ""
		p.payloadBase64 = jwtTokenized[0]
		// keep existing sign info
		// params.payloadSign
	default:
		log.Fatalln("Invalid input for checks")
	}
}

func printUsage() {
	message := `
	Usage: command options
		command:
			- generate
			- check
			- read
			- hex
			- sign
			- cal
		options:
			- for 'generate': <curve> <private key file> <payload in json format> <sign format>
			- for 'check': <curve> <public key file> <payload in base64Url> <payload signature in base64Url> <sign format>
			- for 'read': <payload in base64Url>
			- for 'hex': <public key file>
			- for 'sign': <provider name> <public key file>
			- for 'cal': <provider name> <curve> <public key file>
		curve:
			- k1
			- r1
		sign format:
			- raw: take the payload as it is to sign it
			- jwt: add a '.' (dot) in front of the base64 encoded payload to sign it
	`
	fmt.Println(message)
}
