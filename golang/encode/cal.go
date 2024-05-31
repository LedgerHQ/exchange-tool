package encode

import "encoding/json"

type CalInfo struct {
	Name      string    `json:"name"`
	PublicKey CalPubKey `json:"publicKey"`
	Signature string    `json:"signature"`
	Version   uint      `json:"version"`
}

type CalPubKey struct {
	Curve string `json:"curve"`
	Data  string `json:"data"`
}

func CalFormatProviderInfo(name string, curve string, pubKey string, signApdu string) CalInfo {
	return CalInfo{
		Name: name,
		PublicKey: CalPubKey{
			Curve: curve,
			Data:  pubKey,
		},
		Signature: signApdu,
		Version:   2,
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
