package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
)

func VerifyRSSignature(publicKey *ecdsa.PublicKey, payload string, signature []byte, format SignFormat) bool {
	r, s := extractRS(signature)

	payload = formatMessageToSign(payload, format)
	log.Println("Payload to verify:", payload)

	hashes := crypto.SHA256.New()
	hashes.Write([]byte(payload))
	return ecdsa.Verify(publicKey, hashes.Sum(nil), r, s)
}

func VerifyRSSignature2(publicKey *ecdsa.PublicKey, payload []byte, signature []byte) bool {
	r, s := extractRS(signature)

	log.Println("Payload to verify:", payload)

	hashes := crypto.SHA256.New()
	hashes.Write([]byte(payload))
	return ecdsa.Verify(publicKey, hashes.Sum(nil), r, s)
}

func VerifyDERSignature(publicKey *ecdsa.PublicKey, payload string, signature []byte) bool {
	hashes := crypto.SHA256.New()
	// hashes.Write([]byte(payload))
	payloadBuffer, _ := hex.DecodeString(payload)
	hashes.Write(payloadBuffer)
	return ecdsa.VerifyASN1(publicKey, hashes.Sum(nil), signature)
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
