package set1

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

/*
	Challenege 1: Convert hex to base64

Pre: parameter must be a valid hex encoded byte array
Post: New Byte Array encoded into base64
*/
func HexToBase64(payload []byte) []byte {
	plainText, err := hex.DecodeString(string(payload))
	if err != nil {
		log.Panic("Invalid Input; Byte data is not encoded into hex")
	}
	return []byte(base64.RawURLEncoding.EncodeToString(plainText))
}
