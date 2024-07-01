package cmd

import (
	"fmt"

	"exchange.ledger.fr/crypto"
	"exchange.ledger.fr/encode"
	"github.com/spf13/cobra"
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
	CalCmd.Flags().UintP("version", "v", 2, "app-exchange version signaure (1 or 2)")
	CalCmd.Flags().StringP("app-name", "a", "swap", "application name")
	CalCmd.MarkFlagRequired("curve")
	CalCmd.MarkFlagRequired("public")
	CalCmd.MarkFlagRequired("name")
	CalCmd.MarkFlagRequired("version")
}

func convertCalParameter(cmd *cobra.Command) *params {
	curve := cmd.Flags().Lookup("curve").Value.String()
	pemFile := cmd.Flags().Lookup("public").Value.String()
	name := cmd.Flags().Lookup("name").Value.String()
	version, _ := cmd.Flags().GetUint("version")
	appName := cmd.Flags().Lookup("app-name").Value.String()

	return &params{
		providerName: name,
		curve:        parseCurve(curve),
		pemFile:      pemFile,
		version:      version,
		appName: 	appName,
	}
}

func cal(cmd *cobra.Command, args []string) {
	fmt.Println("*** Generate CAL format info ***")

	params := convertCalParameter(cmd)

	calInfo := generateCal(params.curve, params.pemFile, params.providerName, params.version, params.appName)

	fmt.Println("--> CAL info (copy/paste for Live):", calInfo.Pretty())
}

func generateCal(curve crypto.Curve, filename string, providerName string, version uint, appName string) encode.CalInfo {
	pubKey := curve.ConvertPEMtoHexKey(filename)
	apdu, signature := crypto.SignProviderInfo(providerName, pubKey, curve, version)
	fmt.Println("Signature:", signature)
	fmt.Println("APDU:", apdu)
	calInfo := encode.CalFormatProviderInfo(providerName, curve.Name(), pubKey, version, appName)
	return calInfo
}
