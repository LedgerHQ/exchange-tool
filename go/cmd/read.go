package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"exchange.ledger.fr/encode"
)

var ReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read payload base64Url encoded",
	Args:  cobra.MinimumNArgs(1),
	Run:   read,
}

func read(cmd *cobra.Command, args []string) {
	fmt.Println("*** Extract protobuf file info ***")
	info, format := readPayload(args[0])

	fmt.Println("--> Info (encoded format:", format, ")\n", info)
}

// Base64 decode payload and extract its info from binary result
func readPayload(payload string) (string, encode.Base64Format) {
	payloadByte, format := encode.CascadeDecodeBase64(payload)
	payloadJson := encode.DecodeSwapProtobuf(payloadByte)

	return payloadJson.String(), format
}
