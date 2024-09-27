package cmd

import (
	"fmt"

	"exchange.ledger.fr/crypto"
	"exchange.ledger.fr/encode"
	"github.com/spf13/cobra"
)

var curve string
var pemFile string
var signFormat string

var (
	GenerateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate a Base64Url encoded payload with its signature",
		Long: `Generate a Base64Url encoded payload with its signature.
	The payload is based on the json format inputs.`,
	}

	generateSwapCmd = &cobra.Command{
		Use:   "swap",
		Short: "Generate a Swap Base64Url encoded payload with its signature",
		Long: `Generate a Swap Base64Url encoded payload with its signature.
	The payload is based on the json format inputs.`,
		Run: generate[encode.SwapDevicePayload],
	}

	generateSellCmd = &cobra.Command{
		Use:   "sell",
		Short: "Generate a Sell Base64Url encoded payload with its signature",
		Long: `Generate a Sell Base64Url encoded payload with its signature.
	The payload is based on the json format inputs.`,
		Run: generate[encode.SellDevicePayload],
	}
)

func init() {
	GenerateCmd.AddCommand(generateSwapCmd, generateSellCmd)

	GenerateCmd.PersistentFlags().StringVarP(&curve, "curve", "c", "", "Curve: k1 or r1")
	GenerateCmd.PersistentFlags().StringVarP(&pemFile, "private", "p", "", "Private key file")
	GenerateCmd.PersistentFlags().StringVarP(&signFormat, "format", "f", "jwt", "Sign format: raw or jwt")
	GenerateCmd.MarkPersistentFlagRequired("curve")
	GenerateCmd.MarkPersistentFlagRequired("private")
}

func convertGeneateParameter(args []string) *params {
	return &params{
		curve:           parseCurve(curve),
		pemFile:         pemFile,
		payloadFilename: args[0],
		signFormat:      parseSignFormat(signFormat),
	}
}

func generate[T encode.SwapDevicePayload | encode.SellDevicePayload](cmd *cobra.Command, args []string) {
	switch any(new(T)).(type) {
	case *encode.SwapDevicePayload:
		fmt.Println("*** Generate Swap Proto ***")
	case *encode.SellDevicePayload:
		fmt.Println("*** Generate Sell Proto ***")
	default:
		fmt.Println("*** Unknown Payload Type ***")
	}

	params := convertGeneateParameter(args)

	payload64, sign64 := generateProtoAndSign[T](params.curve, params.signFormat, params.pemFile, params.payloadFilename)

	fmt.Println("--> Result payload base64:\n", payload64)
	fmt.Println("--> Result signature base64:\n", sign64)
}

func generateProtoAndSign[T encode.SwapDevicePayload | encode.SellDevicePayload](curve crypto.Curve, signFormat crypto.SignFormat, privFilename, payloadFilename string) (payload64 string, sign64 string) {
	fnMarshalledFile := func() []byte {
		payloadJson := encode.ConvertFileToDevicePayload[T](payloadFilename)
		return encode.EncodeDevicePaylod(payloadJson)
	}

	privateKey, _ := curve.ReadPrivateKey(privFilename)
	payload64 = fileToBase64(fnMarshalledFile)
	signature := crypto.SignMessageInRS(payload64, privateKey, signFormat)
	sign64 = encode.EncodeBase64(signature)
	return
}

func fileToBase64(fnMarshalledFile marshalFile) string {
	payloadMarshalled := fnMarshalledFile()
	return encode.EncodeBase64(payloadMarshalled)
}
