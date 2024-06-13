package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"exchange.ledger.fr/crypto"
	"exchange.ledger.fr/encode"
)

var CalCmd = &cobra.Command{
	Use:   "cal",
	Short: "Generate CAL info for given provider",
	Run:   cal,
}

func init() {
	CalCmd.Flags().StringP("curve", "c", "", "Curve of provder's pubkey: k1 or r1")
	CalCmd.Flags().StringP("public", "p", "", "Public key file")
	CalCmd.Flags().StringP("name", "n", "", "Provider's name")
	CalCmd.MarkFlagRequired("curve")
	CalCmd.MarkFlagRequired("public")
	CalCmd.MarkFlagRequired("name")
}

func convertCalParameter(cmd *cobra.Command) *params {
	curve := cmd.Flags().Lookup("curve").Value.String()
	pemFile := cmd.Flags().Lookup("public").Value.String()
	name := cmd.Flags().Lookup("name").Value.String()

	return &params{
		providerName: name,
		curve:        parseCurve(curve),
		pemFile:      pemFile,
	}
}

func cal(cmd *cobra.Command, args []string) {
	fmt.Println("*** Generate CAL format info ***")

	params := convertCalParameter(cmd)

	curve := params.curve
	pubKey := curve.ReadPublicKey(params.pemFile)
	signature := crypto.SignProviderInfo(params.providerName, pubKey)
	calInfo := encode.CalFormatProviderInfo(params.providerName, curve.Name(), curve.ConvertPEMtoHexKey(params.pemFile), signature)

	fmt.Println("--> CAL info (copy/paste for Live):", calInfo.String())
	fmt.Println("--> CAL info:", calInfo.Pretty())
}
