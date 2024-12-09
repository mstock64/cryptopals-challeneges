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

/*
Challenge 2: Fixed XOR
Pre: Byte Arrays are the same length
Post
*/
func FixedXor(left, right []byte) []byte {
	leng := len(left)
	result := make([]byte, leng)

	for i := 0; i < leng; i++ {
		result[i] = left[i] ^ right[i]
	}

	return result
}

/*
Challenge 2: Single Byte Xor
Pre:
Post
*/
func SingleByteXor(data []byte, single byte) []byte {
	leng := len(data)
	result := make([]byte, leng)

	for i := 0; i < leng; i++ {
		result[i] = data[i] ^ single
	}

	return result
}

func numberOfValidChar(data []byte) int {
	count := 0
	for _, entry := range data {
		if entry > 63 && entry < 91 {
			count++
		} else if entry > 96 && entry < 123 {
			count++
		}
	}
	return count
}
