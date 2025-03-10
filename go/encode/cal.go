// SPDX-FileCopyrightText: Ledger SAS 2024
//
// SPDX-License-Identifier: Apache-2.0

package encode

import (
	"encoding/json"

	"exchange.ledger.fr/crypto"
)

type CalInfo struct {
	Name                           string            `json:"name"`
	PayloadSignatureComputedFormat crypto.SignFormat `json:"payloadSignatureComputedFormat"`
	PublicKey                      CalPubKey         `json:"publicKey"`
	Service                        Service           `json:"service"`
	Signature                      string            `json:"signature"`
	Apdu                           string            `json:"-"`
}

type CalPubKey struct {
	Curve string `json:"curve"`
	Data  string `json:"data"`
}

type Service struct {
	AppVersion uint   `json:"appVersion"`
	Name       string `json:"name"`
}

func CalFormatProviderInfo(name string, curve string, pubKey string, version uint, serviceName string, signApdu string, apdu string) CalInfo {
	return CalInfo{
		Name:                           name,
		PayloadSignatureComputedFormat: crypto.Jws,
		PublicKey: CalPubKey{
			Curve: curve,
			Data:  pubKey,
		},
		Service: Service{
			AppVersion: version,
			Name:       serviceName,
		},
		Signature: signApdu,
		Apdu:      apdu,
	}
}

func (cal CalInfo) String() string {
	jsonValue, _ := json.Marshal(cal)
	return string(jsonValue)
}

func (cal CalInfo) CalFormat() string {
	type calFormat struct {
		Name                           string            `json:"name"`
		PayloadSignatureComputedFormat crypto.SignFormat `json:"payloadSignatureComputedFormat"`
		PublicKey                      CalPubKey         `json:"publicKey"`
		Service                        Service           `json:"service"`
		Signature                      string            `json:"-"`
		Apdu                           string            `json:"-"`
	}
	jsonValue, _ := json.MarshalIndent(calFormat(cal), "", "\t")

	return string(jsonValue)
}
