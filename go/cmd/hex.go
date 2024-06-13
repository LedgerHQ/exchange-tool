package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"exchange.ledger.fr/crypto"
)

var HexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Convert asn1 private key to hex value",
	Run:   hexConvert,
}

func init() {
	HexCmd.Flags().StringP("curve", "c", "", "Curve: k1 or r1")
	HexCmd.Flags().StringP("type", "t", "", "Type: private or public")
	// HexCmd.MarkFlagRequired("curve")
	HexCmd.MarkFlagRequired("type")
}

func convertHexParameter(cmd *cobra.Command) *params {
	keyType := cmd.Flags().Lookup("type").Value.String()

	params := &params{
		keyType: keyType,
	}
	return params
}

func hexConvert(cmd *cobra.Command, args []string) {
	fmt.Println("*** Convert private key to hex ***")

	params := convertHexParameter(cmd)

	curve := crypto.K1Curve{}
	if (params.keyType == "private") {
		hexValue := curve.ConvertPrivPEMtoHexKey(args[0])
		fmt.Println("--> Hex value:", hexValue)
		} else {
		hexValue := curve.ConvertPEMtoHexKey(args[0])
		fmt.Println("--> Hex value:", hexValue)
	}
}
