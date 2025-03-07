// SPDX-FileCopyrightText: 2024 Ledger
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"

	"exchange.ledger.fr/crypto"
	"github.com/spf13/cobra"
)

var HexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Convert asn1 private key to hex value",
	Run:   hexConvert,
}

func init() {
	HexCmd.Flags().StringP("curve", "c", "", "Curve: k1 or r1")
	HexCmd.Flags().StringP("type", "t", "", "Type: private or public")
	HexCmd.MarkFlagRequired("curve")
	HexCmd.MarkFlagRequired("type")
}

func convertHexParameter(cmd *cobra.Command) *params {
	curve := cmd.Flags().Lookup("curve").Value.String()
	keyType := cmd.Flags().Lookup("type").Value.String()

	params := &params{
		curve:   parseCurve(curve),
		keyType: keyType,
	}
	return params
}

func hexConvert(cmd *cobra.Command, args []string) {
	fmt.Println("*** Convert private key to hex ***")

	params := convertHexParameter(cmd)

	if params.keyType == "private" {
		curve := crypto.K1Curve{}
		hexValue := curve.ConvertPrivPEMtoHexKey(args[0])
		fmt.Println("--> Hex value:", hexValue)
	} else {
		hexValue := params.curve.ConvertPEMtoHexKey(args[0])
		fmt.Println("--> Hex value:", hexValue)
	}
}
