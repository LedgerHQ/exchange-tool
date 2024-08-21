package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
)

func VerifySignature(publicKey *ecdsa.PublicKey, payload string, signature []byte, format SignFormat) bool {
	r, s := extractRS(signature)

	payload = formatMessageToSign(payload, format)
	fmt.Print("Payload to verify: ", payload)

	hashes := crypto.SHA256.New()
	hashes.Write([]byte(payload))
	return ecdsa.Verify(publicKey, hashes.Sum(nil), r, s)
}

func extractRS(signature []byte) (r *big.Int, s *big.Int) {
	if len(signature) != 64 {
		log.Fatalln("Signature is not 64 long", len(signature))
	}
	r = &big.Int{}
	r.SetBytes(signature[:len(signature)/2])
	s = &big.Int{}
	s.SetBytes(signature[len(signature)/2:])
	return
}
