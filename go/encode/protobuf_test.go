package encode_test

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"exchange.ledger.fr/crypto"
	"exchange.ledger.fr/encode"
	swap "exchange.ledger.fr/proto"
	"google.golang.org/protobuf/proto"
)

func TestConvertToSellProtobuf(t *testing.T) {
	// Given
	payload := encode.SellDevicePayload{
		TraderEmail: "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
		InCurrency:  "BTC",
		InAmount:    1150,
		InAddress:   "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
		OutCurrency: "EUR",
		OutAmount: encode.Decimal{
			Coefficient: 1234,
			Exponent:    2,
		},
		DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
	}

	// When
	bin := encode.EncodeDevicePaylod(payload)

	// Then
	if bin == nil {
		t.Fatal("Unable to convert Protobuf payload")
	}
	const expectedResult = "0a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766356d647112034254431a02047e222a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766347465712a03455552320c0a0800000000000004d210023a20350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b"
	if hex.EncodeToString(bin) != expectedResult {
		t.Fatal("Incorrect encoded payload. Got:", hex.EncodeToString(bin))
	}
}

func TestConvertProtobufToSell(t *testing.T) {
	firstPayload, _ := hex.DecodeString("0a2a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766356d647112034254431a02047e222a62633171617230737272723778666b7679356c3634336c79646e77397265353967747a7a7766347465712a03455552320c0a08000000000000000c10023a20350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b")
	secondPayload, _ := encode.CascadeDecodeBase64("Cg50ZXN0QGxlZGdlci5mchIDRVRIGgg5LS4r2pwAACIqMHhkNjkyQ2IxMzQ2MjYyRjU4NEQxN0I0QjQ3MDk1NDUwMWY2NzE1YTgyKgNFVVIyDAoIAAAAAAAABNIQAjogNQrqDJf3R_HQ92CBRhSkdSOAGxrrfQvLuqKk9Gv4GEs")
	thirdPayload, _ := encode.CascadeDecodeBase64("ChticmViYW4uc2VyZ2l1LWV4dEBsZWRnZXIuZnISA0JUQxoDBwTgIiMyTXhEeGh6OXJmZ3JSRjFreHgzdE1vNFlYaENzYTkxdXFmWioDRVVSMgYKAlz5EAI6IDUK6gyX90fx0PdggUYUpHUjgBsa630Ly7qipPRr-BhL")
	bin, _ := hex.DecodeString("4368526e64573174655335776432357551486c68614739764c6d4e7662524944516c524447674d474f535569497a4a4f4e564e71574464344e30703063324979517a4e715a467044516d6f79526c6c53526a6c3061556877536e6c724b674e465656497942676f4357386f51416a6f675634484655476845376549336b4b745546335f5872477436574456454d386b30654b747441317164337763")
	fourthPayload := make([]byte, base64.URLEncoding.DecodedLen(len(bin)))
	base64.URLEncoding.Decode(fourthPayload, bin)

	testCases := []struct {
		payload         []byte
		expectedPayload encode.SellDevicePayload
	}{
		{
			firstPayload,
			encode.SellDevicePayload{
				TraderEmail: "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq",
				InCurrency:  "BTC",
				InAmount:    1150,
				InAddress:   "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf4teq",
				OutCurrency: "EUR",
				OutAmount: encode.Decimal{
					Coefficient: 12,
					Exponent:    2,
				},
				DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
			},
		},
		{
			secondPayload,
			encode.SellDevicePayload{
				TraderEmail: "test@ledger.fr",
				InCurrency:  "ETH",
				InAmount:    4_120_000_000_000_000_000,
				InAddress:   "0xd692Cb1346262F584D17B4B470954501f6715a82",
				OutCurrency: "EUR",
				OutAmount: encode.Decimal{
					Coefficient: 1234,
					Exponent:    2,
				},
				DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
			},
		},
		{
			thirdPayload,
			encode.SellDevicePayload{
				TraderEmail: "breban.sergiu-ext@ledger.fr",
				InCurrency:  "BTC",
				InAmount:    460000,
				InAddress:   "2MxDxhz9rfgrRF1kxx3tMo4YXhCsa91uqfZ",
				OutCurrency: "EUR",
				OutAmount: encode.Decimal{
					Coefficient: 23801,
					Exponent:    2,
				},
				DeviceTransactionId: "350aea0c97f747f1d0f760814614a47523801b1aeb7d0bcbbaa2a4f46bf8184b",
			},
		},
		{
			fourthPayload,
			encode.SellDevicePayload{
				TraderEmail: "gummy.pwnn@yahoo.com",
				InCurrency:  "BTC",
				InAmount:    407845,
				InAddress:   "2N5SjX7x7Jtsb2C3jdZCBj2FYRF9tiHpJyk",
				OutCurrency: "EUR",
				OutAmount: encode.Decimal{
					Coefficient: 23498,
					Exponent:    2,
				},
				DeviceTransactionId: "",
			},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			// When
			payload := encode.DecodeSellProtobuf(tc.payload)

			// Then
			if payload != tc.expectedPayload {
				t.Fatalf("Incorrect decoded payload.\nExpected: %v\nResult: %v", tc.expectedPayload, payload)
			}
		})
	}
}

func TestPartnerInput(t *testing.T) {
	t.SkipNow()
	// Given
	nonce, _ := base64.StdEncoding.DecodeString("NQrqDJf3R/HQ92CBRhSkdSOAGxrrfQvLuqKk9Gv4GEs=")

	// Mercuryo should code:
	outAmount := &swap.UDecimal{ // Will correspond to 300.00€
		Coefficient: binary.BigEndian.AppendUint64(make([]byte, 0), 30000),
		Exponent:    2,
	}
	message := swap.NewSellResponse{
		TraderEmail:         "sara.neila-jimenez@ledger.fr",
		InCurrency:          "BTC",
		InAmount:            binary.BigEndian.AppendUint32(make([]byte, 0), 535382), // 0.00535382 BTC
		InAddress:           "2NA5Rgu8zXTBvUNFpnmjsDCi7vxUmvdP88P",
		OutCurrency:         "EUR",
		OutAmount:           outAmount,
		DeviceTransactionId: nonce,
	}
	payloadMarshalled, _ := proto.Marshal(&message)

	base64 := encode.EncodeBase64(payloadMarshalled)
	log.Println("Base64 payload:", base64)

	curve := crypto.K1Curve{}
	privateKey, _ := curve.ReadPrivateKeyFromHex("308f6a5369aea611d89abf937d0ffaf0b43b457d42cbf0cf754786b3088f17ae")
	signature := crypto.SignMessageInRS(base64, privateKey, crypto.Jwt)
	sign64 := encode.EncodeBase64(signature)
	log.Println("Signaure:", sign64)

	// When
	payload := encode.DecodeSellProtobuf(payloadMarshalled)

	// Then
	expectedPayload := encode.SellDevicePayload{
		TraderEmail: "sara.neila-jimenez@ledger.fr",
		InCurrency:  "BTC",
		InAmount:    535382,
		InAddress:   "2NA5Rgu8zXTBvUNFpnmjsDCi7vxUmvdP88P",
		OutCurrency: "EUR",
		OutAmount: encode.Decimal{
			Coefficient: 30000,
			Exponent:    2,
		},
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
