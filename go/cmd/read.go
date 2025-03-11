// SPDX-FileCopyrightText: Ledger SAS 2024
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"log"

	"exchange.ledger.fr/encode"
	"github.com/spf13/cobra"
)

var ReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read payload base64Url encoded",
	Args:  cobra.MinimumNArgs(1),
	Run:   read,
}

func init() {
	ReadCmd.Flags().StringP("exchangeType", "e", "", "Exchange type (i.e. swap, sell)")
	ReadCmd.MarkFlagRequired("exchangeType")
}

func read(cmd *cobra.Command, args []string) {
	fmt.Println("*** Extract protobuf file info ***")

	exchangeType := cmd.Flags().Lookup("exchangeType").Value.String()
	info, format := readPayload(args[0], exchangeType)

	fmt.Println("--> Info (encoded format:", format, ")\n", info)
}

// Base64 decode payload and extract its info from binary result
func readPayload(payload string, exchangeType string) (string, encode.Base64Format) {
	payloadByte, format := encode.CascadeDecodeBase64(payload)
	log.Println("Read results:\n", payloadByte, "\n", format)
	var payloadFormatted string
	switch exchangeType {
	case "sell":
		payloadFormatted = encode.DecodeSellProtobuf(payloadByte).String()
	// case "fund":
	// 	payloadFormatted = encode.DecodeFundProtobuf(payloadByte).String()
	default:
		payloadFormatted = encode.DecodeSwapProtobuf(payloadByte).String()
	}

	return payloadFormatted, format
}
