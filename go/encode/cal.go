package encode

import (
	"encoding/json"

	"exchange.ledger.fr/crypto"
)

type CalInfo struct {
	Name                           string            `json:"name"`
	PayloadSignatureComputedFormat crypto.SignFormat `json:"payloadSignatureComputedFormat"`
	PublicKey                      CalPubKey         `json:"publicKey"`
	//Signature                      string            `json:"signature"`
	//Version                        uint              `json:"version"`
	//Apdu                           string            `json:"-"`
	Service                        Service           `json:"service"`
}

type CalPubKey struct {
	Curve string `json:"curve"`
	Data  string `json:"data"`
}

type Service struct {
	AppVersion uint   `json:"appVersion"`
	Name       string `json:"name"`
}

func CalFormatProviderInfo(name string, curve string, pubKey string, version uint, serviceName string) CalInfo {
	return CalInfo{
		Name:                           name,
		PayloadSignatureComputedFormat: crypto.Jws,
		PublicKey: CalPubKey{
			Curve: curve,
			Data:  pubKey,
		},
		//Signature: signApdu,
		//Version:   version,
		//Apdu:      apdu,
		Service: Service{
			AppVersion: version,
			Name:       serviceName,
		},
	}
}

func (cal CalInfo) String() string {
	jsonValue, _ := json.Marshal(cal)
	return string(jsonValue)
}

func (cal CalInfo) Pretty() string {
	jsonValue, _ := json.MarshalIndent(cal, "", "\t")
	return string(jsonValue)
}
