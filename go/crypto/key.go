// SPDX-FileCopyrightText: 2024 Ledger
//
// SPDX-License-Identifier: Apache-2.0

package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"strings"

	ethereum "github.com/ethereum/go-ethereum/crypto"
)

type Curve interface {
	ReadPublicKeyFile(filename string) *ecdsa.PublicKey
	ReadPublicKey(contentByte []byte) *ecdsa.PublicKey
	ReadPrivateKey(filename string) (*ecdsa.PrivateKey, error)
	ConvertPEMtoHexKey(filename string) string
	ReadHexPublicKey(hexValue string) *ecdsa.PublicKey
	Name() string
	Flag() string
	Code() []byte
}

type K1Curve struct {
	Curve
}

func (c K1Curve) ConvertPEMtoHexKey(filename string) string {
	contentByte, err := readPemFile(filename)
	if err != nil {
		log.Fatal("Unable to open pub key file:", filename)
	}

	// This part is a substitute to `x509.ParsePKIXPublicKey` which throw an error "unsupported elliptic curve"
	// This is due to the fact that secp256k1 is not supported by the Golang crypto package.
	var pki publicKeyInfo
	if _, err := asn1.Unmarshal(contentByte, &pki); err != nil {
		log.Fatal("Unable to read public key file")
	}

	return hex.EncodeToString(pki.PublicKey.RightAlign())
}

func (c K1Curve) ConvertPrivPEMtoHexKey(filename string) string {
	contentByte, err := readPemFile(filename)
	if err != nil {
		log.Fatal("Unable to open private key file:", filename)
	}

	var privKey ecPrivateKey
	if _, err := asn1.Unmarshal(contentByte, &privKey); err != nil {
		log.Fatal("Unable to read private key file")
	}

	return hex.EncodeToString(privKey.PrivateKey)
}

// --

func (c K1Curve) ReadPublicKeyFile(filename string) *ecdsa.PublicKey {
	contentByte, err := readPemFile(filename)
	if err != nil {
		log.Fatal("Unable to open pub key file:", filename)
	}

	return c.ReadPublicKey(contentByte)
}

func (c K1Curve) ReadPublicKey(contentByte []byte) *ecdsa.PublicKey {
	// This part is a substitute to `x509.ParsePKIXPublicKey` which throw an error "unsupported elliptic curve"
	// This is due to the fact that secp256k1 is not supported by the Golang crypto package.
	var pki publicKeyInfo
	if _, err := asn1.Unmarshal(contentByte, &pki); err != nil {
		log.Fatal("Unable to read public key. Reason:", err.Error())
	}

	asn1Data := pki.PublicKey.RightAlign()

	pubKey, err := ethereum.UnmarshalPubkey(asn1Data)
	if err != nil {
		log.Fatal("Error while reading public key:", err)
	}

	return pubKey
}

func (c K1Curve) ReadHexPublicKey(hexValue string) *ecdsa.PublicKey {
	asn1Data, err := hex.DecodeString(hexValue)
	if err != nil {
		log.Fatalln("Unable to decode pubkey value")
	}

	pubKey, err := ethereum.UnmarshalPubkey(asn1Data)
	if err != nil {
		log.Fatal("Error while reading public key:", err)
	}

	return pubKey
}

// --

func (c K1Curve) ReadPrivateKey(filename string) (*ecdsa.PrivateKey, error) {
	contentByte, err := readPemFile(filename)
	if err != nil {
		log.Fatal("Unable to open private key file:", filename)
	}

	var privKey ecPrivateKey
	if _, err := asn1.Unmarshal(contentByte, &privKey); err != nil {
		log.Fatal("Unable to read private key file")
	}

	return ethereum.ToECDSA(privKey.PrivateKey)
}

func (c K1Curve) ReadPrivateKeyFromHex(hexValue string) (*ecdsa.PrivateKey, error) {
	return ethereum.HexToECDSA(hexValue)
}

func (c K1Curve) GeneratePrivateKey() *ecdsa.PrivateKey {
	init := strings.NewReader("Random truc Random truc Random truc Random truc Random truc")

	privateKey, err := ecdsa.GenerateKey(ethereum.S256(), init)
	if err != nil {
		log.Fatal("Error while generating private key:", err)
	}

	return privateKey
}

func (c K1Curve) Name() string {
	return "secp256k1"
}

func (c K1Curve) Flag() string {
	return "k1"
}

func (c K1Curve) Code() []byte {
	return []byte{0x00}
}

// -- R1
type R1Curve struct {
	Curve
}

func (c R1Curve) ConvertPEMtoHexKey(filename string) string {
	contentByte, err := readPemFile(filename)
	if err != nil {
		log.Fatal("Unable to open pub key file:", filename)
	}

	key, err := x509.ParsePKIXPublicKey(contentByte)
	if err != nil {
		log.Fatal("Error while reading public key:", err)
	}
	pubKey, isOk := key.(*ecdsa.PublicKey)
	if !isOk {
		log.Fatal("Error while type casting key")
	}

	// return hex.EncodeToString(append(pubKey.X.Bytes(), pubKey.Y.Bytes()...))
	return hex.EncodeToString(elliptic.Marshal(elliptic.P256(), pubKey.X, pubKey.Y))
}

func (c R1Curve) ReadPrivateKey(filename string) (privKey *ecdsa.PrivateKey, err error) {
	contentByte, err := readPemFile(filename)
	if err != nil {
		log.Fatal("Unable to open private key file:", filename)
	}

	privKey, err = x509.ParseECPrivateKey(contentByte)
	if err != nil {
		log.Fatal("Error while reading private key:", err)
	}

	return
}

func (c R1Curve) ReadPublicKeyFile(filename string) (pubKey *ecdsa.PublicKey) {
	contentByte, err := readPemFile(filename)
	if err != nil {
		log.Fatal("Unable to open public key file:", filename)
	}

	return c.ReadPublicKey(contentByte)
}

func (c R1Curve) ReadPublicKey(contentByte []byte) (pubKey *ecdsa.PublicKey) {
	key, err := x509.ParsePKIXPublicKey(contentByte)
	if err != nil {
		log.Fatal("Error while reading public key:", err)
	}
	pubKey, isOk := key.(*ecdsa.PublicKey)
	if !isOk {
		log.Fatal("Error while type casting key")
	}

	return pubKey
}

func (c R1Curve) ReadHexPublicKey(hexValue string) *ecdsa.PublicKey {
	contentByte, err := hex.DecodeString(hexValue)
	if err != nil {
		log.Fatalln("Unable to decode pubkey value")
	}

	x, y := elliptic.Unmarshal(elliptic.P256(), contentByte)
	if x == nil && y == nil {
		log.Fatalln("Error while reading public key content")
	}

	return &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
}

func (c R1Curve) GeneratePrivateKey() *ecdsa.PrivateKey {
	init := strings.NewReader("Random truc Random truc Random truc Random truc Random truc")

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), init)
	if err != nil {
		log.Fatal("Error while generating private key:", err)
	}

	return privateKey
}

func (c R1Curve) Name() string {
	return "secp256r1"
}

func (c R1Curve) Flag() string {
	return "r1"
}

func (c R1Curve) Code() []byte {
	return []byte{0x01}
}

// -- Utils
func readPemFile(filename string) ([]byte, error) {
	contentByte, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(contentByte)
	if block == nil {
		return nil, errors.New("Unable to decode PEM file")
	}

	return block.Bytes, nil
}

// Coming from x509 offical package
type ecPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

// Coming from x509 offical package
type publicKeyInfo struct {
	Raw       asn1.RawContent
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}
