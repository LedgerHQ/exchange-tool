// SPDX-FileCopyrightText: Ledger SAS 2024
//
// SPDX-License-Identifier: Apache-2.0

package encode

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestDecode(t *testing.T) {
	encodedSign := "aIyogZIUJFfpYsIlzVycjdP4ls9yzZqXXBSb-meyrFfO5YQxi1v2dkurEZ3v77zsaKUIi1_tOu5HYt3m20332Q"
	decoded, format := CascadeDecodeBase64(encodedSign)
	log.Println("Result:", string(decoded))
	log.Println("Length result:", len(decoded))
	log.Println("Format:", format)
}

func TestDecodeSellSignature(t *testing.T) {
	signedBuffer, _ := CascadeDecodeBase64("rMvQbQ6pZM4QTFkhqBv9FipJHrCfbLRzDGsMwM9CBtsu12uEtv3vrpo4GLObu7KuRdYJ_7RxpYtti7tyehKiFg==")
	signature := hex.EncodeToString(signedBuffer)

	if signature != "accbd06d0ea964ce104c5921a81bfd162a491eb09f6cb4730c6b0cc0cf4206db2ed76b84b6fdefae9a3818b39bbbb2ae45d609ffb471a58b6d8bbb727a12a216" {
		t.Fatalf("Signature mismatch, got %s", signature)
	}
}
