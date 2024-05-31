package encode

import (
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
