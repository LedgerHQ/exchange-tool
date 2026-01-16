// SPDX-FileCopyrightText: Ledger SAS 2024
//
// SPDX-License-Identifier: Apache-2.0

package encode

import (
	"encoding/binary"
	"encoding/hex"
	"log"
	"math/big"

	swap "exchange.ledger.fr/proto"
	"google.golang.org/protobuf/proto"
)

func EncodeDevicePaylod[T SwapDevicePayload | SellDevicePayload | FundDevicePayload](payload T) []byte {
	switch m := any(payload).(type) {
	case SwapDevicePayload:
		message := convertSwapDevicePaylod(m)
		return marshalProtobuf(&message)
	case SellDevicePayload:
		message := convertSellDevicePaylod(m)
		return marshalProtobuf(&message)
	case FundDevicePayload:
		message := convertFundDevicePaylod(m)
		return marshalProtobuf(&message)
	default:
		log.Fatalln("Unknown DevicePayload type")
	}
	return nil
}

func DecodeSwapProtobuf(payload []byte) SwapDevicePayload {
	message := swap.NewTransactionResponse{}
	proto.Unmarshal(payload, &message)

	nonce := hex.EncodeToString(message.DeviceTransactionIdNg)
	if len(nonce) != 0 && len(nonce) != 64 && len(nonce) != 10 {
		log.Printf("Incorrect nonce size. Nonce value received: %s (length: %d)\n", nonce, len(nonce))
	}

	providerAmount := new(big.Int)
	providerAmount.SetBytes(message.AmountToProvider)
	walletAmount := new(big.Int)
	walletAmount.SetBytes(message.AmountToWallet)

	return SwapDevicePayload{
		PayinAddress:        message.PayinAddress,
		PayinExtraId:        message.PayinExtraId,
		RefundAddress:       message.RefundAddress,
		RefundExtraId:       message.RefundExtraId,
		PayoutAddress:       message.PayoutAddress,
		PayoutExtraId:       message.PayoutExtraId,
		CurrencyFrom:        message.CurrencyFrom,
		CurrencyTo:          message.CurrencyTo,
		AmountToProvider:    providerAmount.Uint64(),
		AmountToWallet:      walletAmount.Uint64(),
		DeviceTransactionId: nonce,
	}
}

func convertSwapDevicePaylod(payload SwapDevicePayload) swap.NewTransactionResponse {
	bigNumberToProvider := new(big.Int)
	bigNumberToProvider.SetUint64(payload.AmountToProvider)
	bigNumberToWallet := new(big.Int)
	bigNumberToWallet.SetUint64(payload.AmountToWallet)
	nonce, _ := hex.DecodeString(payload.DeviceTransactionId)

	return swap.NewTransactionResponse{
		PayinAddress:          payload.PayinAddress,
		PayinExtraId:          payload.PayinExtraId,
		RefundAddress:         payload.RefundAddress,
		RefundExtraId:         payload.RefundExtraId,
		PayoutAddress:         payload.PayoutAddress,
		PayoutExtraId:         payload.PayoutExtraId,
		CurrencyFrom:          payload.CurrencyFrom,
		CurrencyTo:            payload.CurrencyTo,
		AmountToProvider:      bigNumberToProvider.Bytes(),
		AmountToWallet:        bigNumberToWallet.Bytes(),
		DeviceTransactionIdNg: nonce,
	}
}

func DecodeSellProtobuf(payload []byte) SellDevicePayload {
	message := swap.NewSellResponse{}
	proto.Unmarshal(payload, &message)

	nonce := hex.EncodeToString(message.DeviceTransactionId)
	if len(nonce) != 0 && len(nonce) != 64 {
		log.Printf("Incorrect nonce size. Nonce value received: %s (length: %d)\n", nonce, len(nonce))
	}

	inAmount := new(big.Int)
	inAmount.SetBytes(message.InAmount)

	var coefficient uint64
	switch len(message.OutAmount.Coefficient) {
	case 2:
		coefficient = uint64(binary.BigEndian.Uint16(message.OutAmount.Coefficient))
	case 4:
		coefficient = uint64(binary.BigEndian.Uint32(message.OutAmount.Coefficient))
	case 8:
		coefficient = binary.BigEndian.Uint64(message.OutAmount.Coefficient)
	default:
		log.Fatalln("Incorrect Coefficient size:", len(message.OutAmount.Coefficient))
	}

	return SellDevicePayload{
		TraderEmail: message.TraderEmail,
		InCurrency:  message.InCurrency,
		InAmount:    inAmount.Uint64(),
		InAddress:   message.InAddress,
		OutCurrency: message.OutCurrency,
		OutAmount: Decimal{
			Coefficient: coefficient,
			Exponent:    message.OutAmount.GetExponent(),
		},
		DeviceTransactionId: nonce,
		InExtraId:           message.InExtraId,
	}
}

func convertSellDevicePaylod(payload SellDevicePayload) swap.NewSellResponse {
	bigNumberIntAmount := new(big.Int)
	bigNumberIntAmount.SetUint64(payload.InAmount)
	nonce, _ := hex.DecodeString(payload.DeviceTransactionId)

	outAmount := swap.UDecimal{
		Coefficient: binary.BigEndian.AppendUint64([]byte{}, payload.OutAmount.Coefficient),
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
		InExtraId:           payload.InExtraId,
	}
}

func DecodeFundProtobuf(payload []byte) FundDevicePayload {
	message := swap.NewFundResponse{}
	proto.Unmarshal(payload, &message)

	nonce := hex.EncodeToString(message.DeviceTransactionId)
	if len(nonce) != 0 && len(nonce) != 64 {
		log.Printf("Incorrect nonce size. Nonce value received: %s (length: %d)\n", nonce, len(nonce))
	}

	inAmount := new(big.Int)
	inAmount.SetBytes(message.InAmount)

	return FundDevicePayload{
		UserId:              message.UserId,
		AccountName:         message.AccountName,
		InCurrency:          message.InCurrency,
		InAmount:            inAmount.Uint64(),
		InAddress:           message.InAddress,
		DeviceTransactionId: nonce,
	}
}

func convertFundDevicePaylod(payload FundDevicePayload) swap.NewFundResponse {
	bigNumberIntAmount := new(big.Int)
	bigNumberIntAmount.SetUint64(payload.InAmount)
	nonce, _ := hex.DecodeString(payload.DeviceTransactionId)

	return swap.NewFundResponse{
		UserId:              payload.UserId,
		AccountName:         payload.AccountName,
		InCurrency:          payload.InCurrency,
		InAmount:            bigNumberIntAmount.Bytes(),
		InAddress:           payload.InAddress,
		DeviceTransactionId: nonce,
	}
}

// Encode Protobuf to binary
type Message[E swap.NewTransactionResponse | swap.NewSellResponse | swap.NewFundResponse] interface {
	*E
	proto.Message
}

func marshalProtobuf[M Message[E], E swap.NewTransactionResponse | swap.NewSellResponse | swap.NewFundResponse](message M) []byte {
	payloadMarshalled, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("Error while marshalling payload to protobuf:", err)
	}

	return payloadMarshalled
}
