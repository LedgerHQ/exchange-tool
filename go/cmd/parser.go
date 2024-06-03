package cmd

import (
	"log"
	"strings"

	"swap.ledger.fr/crypto"
)

type params struct {
	curve           crypto.Curve
	pemFile         string
	payloadFilename string
	payloadBase64   string
	signatureBase64 string
	payloadHeader   string
	signFormat      crypto.SignFormat
	providerName    string
}

// JWT format: <Base64Url(HEADER)>.<Base64Url(PAYLOAD)>.<Base64Url(SIGNATURE)>
// where: SIGNATURE is computed on <Base64Url(HEADER)>.<Base64Url(PAYLOAD)>
func (p *params) fillJWTPayload(jwtPayload string) {
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

type marshalFile func() []byte

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
