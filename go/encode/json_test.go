// SPDX-FileCopyrightText: 2024 Ledger
//
// SPDX-License-Identifier: Apache-2.0

package encode

import (
	"testing"
)

func TestSwapDecode(t *testing.T) {
	// Given
	filename := "../../samples/swap-payload-example.json"

	// When
	payload := ConvertFileToDevicePayload[SwapDevicePayload](filename)

	// Then
	expectedPayload := SwapDevicePayload{
		PayinAddress:        "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
		RefundAddress:       "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
		PayoutAddress:       "0xb794f5ea0ba39494ce839613fffba74279579268",
		CurrencyFrom:        "BTC",
		CurrencyTo:          "BAT",
		AmountToProvider:    1150,
		AmountToWallet:      6000000000000,
		DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
	}
	if payload != expectedPayload {
		t.Fatalf("Unable to correctly read Swap file.\nExpected: %v\nReceived: %v", expectedPayload, payload)
	}
}

func TestFundDecode(t *testing.T) {
	// Given
	filename := "../../samples/fund-payload-example.json"

	// When
	payload := ConvertFileToDevicePayload[FundDevicePayload](filename)

	// Then
	expectedPayload := FundDevicePayload{
		UserId:              "12343456",
		AccountName:         "John Doe",
		InCurrency:          "BTC",
		InAmount:            1150,
		InAddress:           "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
		DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
	}
	if payload != expectedPayload {
		t.Fatalf("Unable to correctly read Fund file.\nExpected: %v\nReceived: %v", expectedPayload, payload)
	}
}
