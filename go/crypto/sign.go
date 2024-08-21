package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"log"
	"math/big"

	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/crypto/cryptobyte/asn1"
)

type SignFormat string

const (
	Raw SignFormat = "raw"
	Jwt SignFormat = "jwt"
	Jws SignFormat = "jws"
)

// Generate the payload from the provided JSON file
//
// Returns the base64URL encoded payload and its base64UL encoded signature
func SignMessageInRS(message string, privKey *ecdsa.PrivateKey, format SignFormat) []byte {
	//-- Sign base64URL payload (don't forget to add a '.' before as we don't require the alg info)
	//-- [JWS RFC](https://www.rfc-editor.org/rfc/rfc7515#section-5.1)
	//payloadSignature := signPayload([]byte("."+payload64), &privKey)
	signature := SignMessageInDER(message, privKey, format)
	return convertToRS(signature)
}

func SignMessageInDER(message string, privKey *ecdsa.PrivateKey, format SignFormat) []byte {
	message = formatMessageToSign(message, format)

	return signPayloadNew([]byte(message), privKey)
}

func signPayloadNew(data []byte, privateKey *ecdsa.PrivateKey) []byte {
	hashes := crypto.SHA256.New()
	hashes.Write(data)
	sign, err := ecdsa.SignASN1(rand.Reader, privateKey, hashes.Sum(nil))
	if err != nil {
		log.Fatal("Unable to sign payload:", err)
	}

	return sign
}

func convertToRS(signature []byte) []byte {
	r := new(big.Int)
	s := new(big.Int)
	value := new(cryptobyte.String)
	sign := cryptobyte.String(signature)
	if !sign.ReadASN1(value, asn1.SEQUENCE) ||
		!value.ReadASN1Integer(r) ||
		!value.ReadASN1Integer(s) {
		log.Fatal("Unable to convert DER to R|S value")
	}

	return concatRS(r, s)
}

func concatRS(r, s *big.Int) []byte {
	sign := make([]byte, 0)
	sign = append(sign, r.Bytes()...)
	sign = append(sign, s.Bytes()...)

	return sign
}

func formatMessageToSign(message string, format SignFormat) string {
	if format == Jwt {
		return "." + message
	}
	return message
}
