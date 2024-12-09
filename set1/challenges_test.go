package set1

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestHexToBase64(T *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result := HexToBase64([]byte(input))

	if string(result) != expectedOutput {
		T.FailNow()
	}
}

func TestFixedXOR(T *testing.T) {
	left, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	right, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	expectedOutput := "746865206b696420646f6e277420706c6179"
	result := FixedXor(left, right)

	if hex.EncodeToString(result) != expectedOutput {
		log.Println(string(result))
		T.FailNow()
	}
}

func TestSingeByteXOR(T *testing.T) {
	data, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	expectedOutput := "Cooking MC's like a pound of bacon"
	var results []int
	for i := 0; i < 127; i++ {
		entry := SingleByteXor(data, byte(i))
		num := numberOfValidChar(entry)
		results = append(results, num)
	}
	best, index := 0, 0
	for i := 0; i < len(results); i++ {
		if best < results[i] {
			best = results[i]
			index = i
		}
	}
	gamble := SingleByteXor(data, byte(index))
	if string(gamble) != expectedOutput {
		T.FailNow()
	}
}
