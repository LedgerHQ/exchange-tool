package cmd

import (
	"crypto/ecdsa"
	"fmt"

	"exchange.ledger.fr/crypto"
	"exchange.ledger.fr/encode"
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the payload validity against the provided signature",
	Run:   check,
}

type checkStatus struct {
	isOk         bool
	base64Format encode.Base64Format
}

func init() {
	CheckCmd.Flags().StringP("curve", "c", "", "Curve: k1 or r1")
	CheckCmd.Flags().StringP("public", "p", "", "Public key file")
	CheckCmd.Flags().StringP("hex", "x", "", "Public key hex value")
	CheckCmd.Flags().StringP("format", "f", "jwt", "Sign format: raw or jwt")
	CheckCmd.MarkFlagRequired("curve")
	CheckCmd.MarkFlagsMutuallyExclusive("public", "hex")
}

func convertCheckParameter(cmd *cobra.Command, args []string) *params {
	curve := cmd.Flags().Lookup("curve").Value.String()
	pemFile := cmd.Flags().Lookup("public").Value.String()
	pemHex := cmd.Flags().Lookup("hex").Value.String()
	signFormat := cmd.Flags().Lookup("format").Value.String()

	params := &params{
		curve:           parseCurve(curve),
		pemFile:         pemFile,
		pemHex:          pemHex,
		signatureBase64: args[1], // Expected in base64
		signFormat:      parseSignFormat(signFormat),
	}
	params.fillJWTPayload(args[0])
	return params
}

func check(cmd *cobra.Command, args []string) {
	fmt.Println("*** Check signature ***")

	params := convertCheckParameter(cmd, args)

	var status checkStatus
	if params.pemHex == "" {
		status = checkPayloadWithKeyFile(params.curve, params.signFormat, params.pemFile, params.payloadBase64, params.signatureBase64)
	} else {
		status = checkPayloadWithKeyHex(params.curve, params.signFormat, params.pemHex, params.payloadBase64, params.signatureBase64)
	}

	if status.isOk {
		fmt.Println("--> Payload is CORRECTLY signed")
	} else {
		fmt.Println("--> Payload is NOT CORRECTLY signed")
	}
}

// Check provided base64 URL encoded payload (which must be a binary protobuf) that match the signature
func checkPayloadWithKeyFile(curve crypto.Curve, signFormat crypto.SignFormat, pubFilename, payload, signature string) checkStatus {
	publicKey := curve.ReadPublicKeyFile(pubFilename)
	return checkPayload(publicKey, signFormat, payload, signature)
}

func checkPayloadWithKeyHex(curve crypto.Curve, signFormat crypto.SignFormat, pubHexValue, payload, signature string) checkStatus {
	publicKey := curve.ReadHexPublicKey(pubHexValue)
	return checkPayload(publicKey, signFormat, payload, signature)
}

func checkPayload(publicKey *ecdsa.PublicKey, signFormat crypto.SignFormat, payload, signature string) checkStatus {
	signatureByte, format := encode.CascadeDecodeBase64(signature)
	status := checkStatus{
		base64Format: format,
		isOk:         true,
	}

	status.isOk = crypto.VerifyRSSignature(publicKey, payload, signatureByte, signFormat)

	return status
}
