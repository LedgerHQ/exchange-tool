// SPDX-FileCopyrightText: 2024 Ledger
//
// SPDX-License-Identifier: Apache-2.0

package encode

import (
	"encoding/base64"
	"log"
)

type Base64Format string

const (
	rawUrl Base64Format = "RawUrl"
	url    Base64Format = "Url"
	rawStd Base64Format = "RawStd"
	std    Base64Format = "Std"
)

func EncodeBase64(payload []byte) string {
	return base64.URLEncoding.EncodeToString(payload)
}

func CascadeDecodeBase64(payload string) ([]byte, Base64Format) {
	format := rawUrl
	payloadByte, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		format = url
		payloadByte, err = base64.URLEncoding.DecodeString(payload)
		if err != nil {
			format = rawStd
			payloadByte, err = base64.RawStdEncoding.DecodeString(payload)
			if err != nil {
				format = std
				payloadByte, err = base64.StdEncoding.DecodeString(payload)
				if err != nil {
					format = ""
					log.Fatal("Unable to base64 decode payload", err)
				}
			}
			log.Println("WARNING: this base64 string is encoded in Base64Std, not in required Base64URL")
		}
	}

	return payloadByte, format
}
