package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"exchange.ledger.fr/crypto"
)

var SignCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign public key with Ledger's debug key",
	Run:   sign,
}

func init() {
	SignCmd.Flags().StringP("curve", "c", "", "Curve of provder's pubkey: k1 or r1")
	SignCmd.Flags().StringP("name", "n", "", "Provider's name")
	SignCmd.MarkFlagRequired("curve")
	SignCmd.MarkFlagRequired("name")
}

func convertSignParameter(cmd *cobra.Command, args []string) *params {
	curve := cmd.Flags().Lookup("curve").Value.String()
	name := cmd.Flags().Lookup("name").Value.String()

	return &params{
		curve:        parseCurve(curve),
		pemFile:      args[0],
		providerName: name,
	}
}

func sign(cmd *cobra.Command, args []string) {
	fmt.Println("*** Sign public key ***")

	params := convertSignParameter(cmd, args)

	pubKey := params.curve.ReadPublicKey(params.pemFile)
	signature := crypto.SignProviderInfo(params.providerName, pubKey)

	fmt.Println("--> Signature value:\n", signature)
}
