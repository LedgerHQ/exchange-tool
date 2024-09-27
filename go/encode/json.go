package encode

import (
	"encoding/json"
	"log"
	"os"
)

type SwapDevicePayload struct {
	PayinAddress          string `json:"payinAddress"`
	RefundAddress         string `json:"refundAddress"`
	PayoutAddress         string `json:"payoutAddress"`
	CurrencyFrom          string `json:"currencyFrom"`
	CurrencyTo            string `json:"currencyTo"`
	AmountToProvider      uint64 `json:"amountToProvider"`
	AmountToWallet        uint64 `json:"amountToWallet"`
	DeviceTransactionId   string `json:"nonce"`
	DeviceTransactionIdNg string `json:"deviceTransactionIdNg"`
	PayinExtraData        string `json:"payinExtraData"`
}

type Decimal struct {
	Coefficient uint64 `json:"coefficient"`
	Exponent    uint32 `json:"exponent"`
}
type SellDevicePayload struct {
	TraderEmail         string  `json:"traderEmail"`
	InCurrency          string  `json:"inCurrency"`
	InAmount            uint64  `json:"inAmount"`
	InAddress           string  `json:"inAddress"`
	OutCurrency         string  `json:"outCurrency"`
	OutAmount           Decimal `json:"outAmount"`
	DeviceTransactionId string  `json:"nonce"`
}

type FundDevicePayload struct {
	UserId              string `json:"userId"`
	AccountName         string `json:"accountName"`
	InCurrency          string `json:"inCurrency"`
	InAmount            uint64 `json:"inAmount"`
	InAddress           string `json:"inAddress"`
	DeviceTransactionId string `json:"nonce"`
}

func ConvertFileToDevicePayload[T SwapDevicePayload | SellDevicePayload | FundDevicePayload](filename string) T {
	fileByte, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error while reading payload file:", filename)
	}

	payload := new(T)
	err = json.Unmarshal(fileByte, payload)
	if err != nil {
		log.Fatal("Invalid json payload file:", err)
	}

	return *payload
}

func (p SwapDevicePayload) String() string {
	jsonValue, _ := json.Marshal(p)
	return string(jsonValue)
}

func (p SellDevicePayload) String() string {
	jsonValue, _ := json.Marshal(p)
	return string(jsonValue)
}
