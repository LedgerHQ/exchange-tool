package encode_test

import (
	"encoding/hex"
	"testing"

	"exchange.ledger.fr/encode"
)

func TestConvertToSellProtobuf(t *testing.T) {
	// Given
	payload := encode.SellDevicePayload{
		TraderEmail:         "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
		InCurrency:          "BTC",
		InAmount:            1150,
		InAddress:           "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
		OutCurrency:         "EUR",
		OutAmount:           12,
		DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
	}

	// When
	bin := encode.EncodeDevicePaylod(payload)

	// Then
	if bin == nil {
		t.Fatal("Unable to convert Protobuf payload")
	}
	if hex.EncodeToString(bin) != "0a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766356d647112034254431a02047e222a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766347465712a03455552320c0a08000000000000000c10023a20350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b" {
		t.Fatal("Incorrect encoded payload")
	}
}

func TestConvertProtobufToSell(t *testing.T) {
	t.Skip("Review UDecimal encoding/decoding first")
	// Given
	bin, _ := hex.DecodeString("0a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766356d647112034254431a02047e222a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766347465712a03455552320c0a08000000000000000c10023a20350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b")

	// When
	payload := encode.DecodeSellProtobuf(bin)

	// Then
	expectedPayload := encode.SellDevicePayload{
		TraderEmail:         "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
		InCurrency:          "BTC",
		InAmount:            1150,
		InAddress:           "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
		OutCurrency:         "EUR",
		OutAmount:           12,
		DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
	}
	if payload != expectedPayload {
		t.Fatalf("Incorrect decoded payload.\nExpected: %v\nResult: %v", expectedPayload, payload)
	}
}

func TestConvertToSwapProtobuf(t *testing.T) {
	// Given
	payload := encode.SwapDevicePayload{
		PayinAddress:        "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
		RefundAddress:       "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
		PayoutAddress:       "0xb794f5ea0ba39494ce839613fffba74279579268",
		CurrencyFrom:        "BTC",
		CurrencyTo:          "BAT",
		AmountToProvider:    1150,
		AmountToWallet:      6000000000000,
		DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
	}

	// When
	bin := encode.EncodeDevicePaylod(payload)

	// Then
	if bin == nil {
		t.Fatal("Unable to convert Protobuf payload")
	}
	if hex.EncodeToString(bin) != "0a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766356d64711a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766347465712a2a3078623739346635656130626133393439346365383339363133666666626137343237393537393236383a0342544342034241544a02047e52060574fbde60006220350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b" {
		t.Fatal("Incorrect encoded payload")
	}
}

func TestConvertProtobufToSwap(t *testing.T) {
	// Given
	bin, _ := hex.DecodeString("0a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766356d64711a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766347465712a2a3078623739346635656130626133393439346365383339363133666666626137343237393537393236383a0342544342034241544a02047e52060574fbde60006220350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b")

	// When
	payload := encode.DecodeSwapProtobuf(bin)

	// Then
	expectedPayload := encode.SwapDevicePayload{
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
		t.Fatalf("Incorrect decoded payload.\nExpected: %v\nResult: %v", expectedPayload, payload)
	}
}

func TestConvertToFundProtobuf(t *testing.T) {
	// Given
	payload := encode.FundDevicePayload{
		UserId:              "12343456",
		AccountName:         "John Doe",
		InCurrency:          "BTC",
		InAmount:            1150,
		InAddress:           "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
		DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
	}

	// When
	bin := encode.EncodeDevicePaylod(payload)

	// Then
	if bin == nil {
		t.Fatal("Unable to convert Protobuf payload")
	}
	if hex.EncodeToString(bin) != "0a08313233343334353612084a6f686e20446f651a034254432202047e2a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766347465713220350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b" {
		t.Fatal("Incorrect encoded payload")
	}
}

func TestConvertProtobufToFund(t *testing.T) {
	// Given
	bin, _ := hex.DecodeString("0a08313233343334353612084a6f686e20446f651a034254432202047e2a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766347465713220350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b")

	// When
	payload := encode.DecodeFundProtobuf(bin)

	// Then
	expectedPayload := encode.FundDevicePayload{
		UserId:              "12343456",
		AccountName:         "John Doe",
		InCurrency:          "BTC",
		InAmount:            1150,
		InAddress:           "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
		DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
	}
	if payload != expectedPayload {
		t.Fatal("Incorrect decoded payload")
	}
}
