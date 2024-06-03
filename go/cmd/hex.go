package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"swap.ledger.fr/crypto"
)

var HexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Convert asn1 private key to hex value",
	Run:   HexConvert,
}

func HexConvert(cmd *cobra.Command, args []string) {
	fmt.Println("*** Convert private key to hex ***")

	curve := crypto.K1Curve{}
	hexValue := curve.ConvertPrivPEMtoHexKey(args[0])

	fmt.Println("--> Hex value:", hexValue)
}
