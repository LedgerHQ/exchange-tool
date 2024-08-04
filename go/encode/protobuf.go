package encode

import (
	//"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"

	swap "exchange.ledger.fr/proto"
	"google.golang.org/protobuf/proto"
)

func EncodeDevicePaylod[T SwapDevicePayload | SellDevicePayload](payload T) []byte {
	switch m := any(payload).(type) {
	case SwapDevicePayload:
		message := convertSwapDevicePaylod(m)
		return marshalProtobuf(&message)
	case SellDevicePayload:
		message := convertSellDevicePaylod(m)
		return marshalProtobuf(&message)
	default:
		log.Fatalln("Unknown DevicePayload type")
	}
	return nil
}

func DecodeSwapProtobuf(payload []byte) SwapDevicePayload {
	message := swap.NewTransactionResponse{}
	proto.Unmarshal(payload, &message)

	providerAmount := new(big.Int)
	providerAmount.SetBytes(message.AmountToProvider)
	walletAmount := new(big.Int)
	walletAmount.SetBytes(message.AmountToWallet)

	fmt.Println("Decoded nonce: -> ", message.DeviceTransactionIdNg)

	// Print the array of bytes
	fmt.Printf("Byte array: %v\n", message.DeviceTransactionIdNg)

	// Read the nonce as Hex
	nonce := fmt.Sprintf("%x", message.DeviceTransactionIdNg)

	return SwapDevicePayload{
		PayinAddress:     message.PayinAddress,
		PayinExtraId:     message.PayinExtraId,
		RefundAddress:    message.RefundAddress,
		RefundExtraId:    message.RefundExtraId,
		PayoutAddress:    message.PayoutAddress,
		PayoutExtraId:    message.PayoutExtraId,
		CurrencyFrom:     message.CurrencyFrom,
		CurrencyTo:       message.CurrencyTo,
		AmountToProvider: providerAmount.String(),
		AmountToWallet:   walletAmount.String(),
		Nonce:            nonce,
	}
}

func convertSwapDevicePaylod(payload SwapDevicePayload) swap.NewTransactionResponse {
	nonce, _ := hex.DecodeString(payload.Nonce)

	return swap.NewTransactionResponse{
		PayinAddress:          payload.PayinAddress,
		RefundAddress:         payload.RefundAddress,
		PayoutAddress:         payload.PayoutAddress,
		CurrencyFrom:          payload.CurrencyFrom,
		CurrencyTo:            payload.CurrencyTo,
		AmountToProvider:      []byte(payload.AmountToProvider),
		AmountToWallet:        []byte(payload.AmountToWallet),
		DeviceTransactionIdNg: nonce,
	}
}

func DecodeSellProtobuf(payload []byte) SellDevicePayload {
	message := swap.NewSellResponse{}
	proto.Unmarshal(payload, &message)

	inAmount := new(big.Int)
	inAmount.SetBytes(message.InAmount)

	outAmount, _ := strconv.ParseUint(message.OutAmount.String(), 10, 64)

	return SellDevicePayload{
		TraderEmail:         message.TraderEmail,
		InCurrency:          message.InCurrency,
		InAmount:            inAmount.Uint64(),
		InAddress:           message.InAddress,
		OutCurrency:         message.OutCurrency,
		OutAmount:           outAmount,
		DeviceTransactionId: string(message.DeviceTransactionId),
	}
}

func convertSellDevicePaylod(payload SellDevicePayload) swap.NewSellResponse {
	bigNumberIntAmount := new(big.Int)
	bigNumberIntAmount.SetUint64(payload.InAmount)
	nonce, _ := hex.DecodeString(payload.DeviceTransactionId)

	outAmount := swap.UDecimal{
		Coefficient: binary.BigEndian.AppendUint64([]byte{}, payload.OutAmount),
		Exponent:    2,
	}

	return swap.NewSellResponse{
		TraderEmail:         payload.TraderEmail,
		InCurrency:          payload.InCurrency,
		InAmount:            bigNumberIntAmount.Bytes(),
		InAddress:           payload.InAddress,
		OutCurrency:         payload.OutCurrency,
		OutAmount:           &outAmount,
		DeviceTransactionId: nonce,
	}
}

// Encode Protobuf to binary
type Message[E swap.NewTransactionResponse | swap.NewSellResponse] interface {
	*E
	proto.Message
}

func marshalProtobuf[M Message[E], E swap.NewTransactionResponse | swap.NewSellResponse](message M) []byte {
	payloadMarshalled, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("Error while marshalling payload to protobuf:", err)
	}

	return payloadMarshalled
}
