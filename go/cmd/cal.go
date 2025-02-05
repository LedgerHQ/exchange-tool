package cmd

import (
	"fmt"
	"log"

	"exchange.ledger.fr/crypto"
	"exchange.ledger.fr/encode"
	"github.com/spf13/cobra"
)

var (
	CalCmd = &cobra.Command{
		Use:   "cal",
		Short: "Generate CAL info like",
	}

	mockCalCmd = &cobra.Command{
		Use:   "cal",
		Short: "Generate CAL info for given provider in strict CAL format",
		Run:   cal,
	}

	liveCalCmd = &cobra.Command{
		Use:   "live",
		Short: "Generate CAL info for given provider in LLD compact format",
		Long:  "Generate CAL info for given provider in LLD compact format. This format is used for swap test to set a mock provider",
		Run:   live,
	}

	coinCalCmd = &cobra.Command{
		Use:   "coin",
		Short: "Generate CAL info for a coin",
		Run:   coin,
	}
)

func init() {
	CalCmd.AddCommand(mockCalCmd, liveCalCmd, coinCalCmd)

	mockCalCmd.Flags().StringP("curve", "c", "", "Curve of provder's pubkey: k1 or r1")
	mockCalCmd.Flags().StringP("public", "p", "", "Public key file")
	mockCalCmd.Flags().StringP("name", "n", "", "Provider's name")
	mockCalCmd.Flags().UintP("version", "v", 2, "app-exchange version signaure (1 or 2)")
	mockCalCmd.Flags().StringP("app-name", "a", "swap", "application name")
	mockCalCmd.MarkFlagRequired("curve")
	mockCalCmd.MarkFlagRequired("public")
	mockCalCmd.MarkFlagRequired("name")
	mockCalCmd.MarkFlagRequired("version")

	liveCalCmd.Flags().StringP("curve", "c", "", "Curve of provder's pubkey: k1 or r1")
	liveCalCmd.Flags().StringP("public", "p", "", "Public key file")
	liveCalCmd.Flags().StringP("name", "n", "", "Provider's name")
	liveCalCmd.Flags().UintP("version", "v", 2, "app-exchange version signaure (1 or 2)")
	liveCalCmd.Flags().StringP("app-name", "a", "swap", "application name")
	liveCalCmd.MarkFlagRequired("curve")
	liveCalCmd.MarkFlagRequired("public")
	liveCalCmd.MarkFlagRequired("name")
	liveCalCmd.MarkFlagRequired("version")

	coinCalCmd.Flags().StringP("ticker", "t", "", "Coin's ticker")
	coinCalCmd.Flags().StringP("application", "a", "", "AppCoin name (beware of the case)")
	coinCalCmd.Flags().UintP("magnitude", "m", 0, "Coin's magnitude")
	coinCalCmd.Flags().UintP("chain Id", "c", 0, "In EVM cases only")
	coinCalCmd.MarkFlagRequired("ticker")
	coinCalCmd.MarkFlagRequired("application")
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
		appName:      appName,
	}
}

func common(cmd *cobra.Command) encode.CalInfo {
	params := convertCalParameter(cmd)

	return generateCal(params.curve, params.pemFile, params.providerName, params.version, params.appName)
}

func cal(cmd *cobra.Command, args []string) {
	fmt.Println("*** Generate CAL format info ***")

	calInfo := common(cmd)

	fmt.Println("--> CAL format:\n", calInfo.CalFormat())
}

func live(cmd *cobra.Command, args []string) {
	fmt.Println("*** Generate CAL info in Live info ***")

	calInfo := common(cmd)

	fmt.Println("--> Ledger Live format:\n", calInfo.String())
}

func coin(cmd *cobra.Command, args []string) {
	fmt.Println("*** Generate CAL format info ***")

	app := cmd.Flags().Lookup("application").Value.String()
	ticker := cmd.Flags().Lookup("ticker").Value.String()
	magnitude, _ := cmd.Flags().GetUint("magnitude")
	chainId, _ := cmd.Flags().GetUint("chain Id")

	var subConfig *crypto.SubConfig = nil
	if magnitude != 0 {
		subConfig = &crypto.SubConfig{
			Ticker:    ticker,
			Magnitude: uint8(magnitude),
			ChainId:   uint16(chainId),
		}
	}

	serialized, _ := crypto.GenerateCoinConfig(crypto.CoinConfig{
		Ticker:    ticker,
		AppName:   app,
		SubConfig: subConfig,
	})

	fmt.Println("--> CAL Coin format:\n", serialized)
}

func generateCal(curve crypto.Curve, filename string, providerName string, version uint, appName string) encode.CalInfo {
	pubKey := curve.ConvertPEMtoHexKey(filename)
	apdu, signature := crypto.SignProviderInfo(providerName, pubKey, curve, version)
	log.Println("Signature:", signature)
	log.Println("APDU:", apdu)
	calInfo := encode.CalFormatProviderInfo(providerName, curve.Name(), pubKey, version, appName, signature, apdu)
	return calInfo
}
