package encode

import (
	"testing"
)

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
		t.Fatal("Unable to corretly read Fund file")
	}
}
