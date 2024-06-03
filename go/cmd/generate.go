package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"swap.ledger.fr/crypto"
	"swap.ledger.fr/encode"
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
	fmt.Println("*** Generate Swap Proto ***")

	params := convertGeneateParameter(args)

	marshalledPayload := func() []byte {
		payloadJson := encode.ConvertFileToDevicePayload[T](params.payloadFilename)
		return encode.EncodeDevicePaylod(payloadJson)
	}
	payload64, sign64 := generateProtoAndSign(params.curve, params.signFormat, params.pemFile, marshalledPayload)

	fmt.Println("--> Result payload base64:\n", payload64)
	fmt.Println("--> Result signature base64:\n", sign64)

	// hash := sha256.New()
	// hash.Write([]byte(payload64))
	// fmt.Println("--> Sha256 of base64 payload with no prefix:\n", hex.EncodeToString(hash.Sum(nil)))
	// hash = sha256.New()
	// hash.Write([]byte("." + payload64))
	// fmt.Println("--> Sha256 of base64 payload with prefix:\n", hex.EncodeToString(hash.Sum(nil)))
	// fmt.Println("--> base64 payload size:", len(payload64))
}

func generateProtoAndSign(curve crypto.Curve, signFormat crypto.SignFormat, privFilename string, fnMarshalledFile marshalFile) (payload64 string, sign64 string) {
	privateKey, _ := curve.ReadPrivateKey(privFilename)
	payload64 = fileToBase64(fnMarshalledFile)
	sign64 = crypto.SignMessageInRS(payload64, privateKey, signFormat)
	return
}

func fileToBase64(fnMarshalledFile marshalFile) string {
	payloadMarshalled := fnMarshalledFile()
	return encode.EncodeBase64(payloadMarshalled)
}
