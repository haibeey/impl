package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestHexTobase64(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	hexStringDecoded := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	ds, err := hexToBase64(hexString)
	if err != nil {
		t.Errorf("%s", err)
	}
	if ds != hexStringDecoded {
		t.Errorf("decoded value and real value doesn't match")
	}
}

func TestHexXOR(t *testing.T) {

	a := make([]byte, 10)
	b := make([]byte, 20)

	maxAB := max(len(a), len(b))

	padByteArr(&a, maxAB)
	padByteArr(&b, maxAB)

	if len(a) != len(b) {
		t.Errorf("%d %d", len(a), len(b))
		t.Errorf("len must be equal after padding")
	}

	hexStringA := "1c0111001f010100061a024b53535009181c"
	hexStringB := "686974207468652062756c6c277320657965"
	expectedRes := "746865206b696420646f6e277420706c6179"

	b, _, err := xorHex(hexStringA, hexStringB)
	if err != nil {
		t.Errorf("%s", err)
	}
	if expectedRes != hex.EncodeToString(b) {
		t.Errorf("computed results and real value doesn't match")
	}
}

func TestSinglebyteXORcipher(t *testing.T) {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	_, ds, err := singlebyteXORcipher(hexString)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Println(ds)
}

func TestFileSinglebyteXORcipher(t *testing.T) {

	res, err := detectSingleCharacterXOR()
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Println(res)
}

func TestRepeatingXor(t *testing.T) {

	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	realResult := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	res := fmt.Sprintf("%x", repeatingkeyXOR(text, []byte("ICE")))

	if res != realResult {
		fmt.Println(res)
		t.Errorf("computed result and real result does not match")
	}
}

func TestEditDistance(t *testing.T) {
	textA := []byte("this is a test")
	textB := []byte("wokka wokka!!!")
	expectedRes := 37
	if editDistance(textA, textB) != expectedRes {
		fmt.Println("computed editdistance ", editDistance(textA, textB))
		t.Errorf("wrong computation in edit distance ")
	}
}

func TestBreakingRepeatingXor(t *testing.T) {
	err := breakRepeatingkeyXOR()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestAesECBmode(t *testing.T) {
	text, err := ioutil.ReadFile("testfiles/ecbmode.in")
	if err != nil {
		t.Errorf("%s", err)
	}

	fileContent, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		t.Errorf("%s", err)
	}

	fmt.Println(fmt.Sprintf("%s", aesECBDecrypt("YELLOW SUBMARINE", fileContent))[0])

}

func TestDetectECB(t *testing.T) {
	err := detectAESECB()
	if err != nil {
		t.Errorf("%s", err)
	}

}
