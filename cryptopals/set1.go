package main

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/bits"
	"strings"
)

func decodeHex(s string) ([]byte, string, error) {
	ds, err := hex.DecodeString(s)
	res := fmt.Sprintf("%s", ds)
	return ds, res, err
}

func hexToBase64(s string) (string, error) {
	hexDecode, _, err := decodeHex(s)
	if err != nil {
		return "", err
	}
	str := base64.StdEncoding.EncodeToString(hexDecode)
	return str, nil
}
func max(a ...int) int {
	b := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > b {
			b = a[i]
		}
	}
	return b
}

func min(a ...int) int {
	b := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < b {
			b = a[i]
		}
	}
	return b

}

func padByteArr(a *[]byte, lenght int) {
	toAddLen := lenght - len(*a)
	if toAddLen <= 0 {
		return
	}
	toAdd := make([]byte, toAddLen)
	*a = append(*a, toAdd...)
}

func xorHex(a, b string) ([]byte, string, error) {
	dA, _, err := decodeHex(a)
	if err != nil {
		return nil, "", err
	}
	dB, _, err := decodeHex(b)
	if err != nil {
		return nil, "", err
	}
	maxAB := max(len(dA), len(dB))

	padByteArr(&dA, maxAB)
	padByteArr(&dB, maxAB)

	res := make([]byte, maxAB)
	for i := 0; i < maxAB; i++ {
		res[i] = dA[i] ^ dB[i]
	}

	return res, fmt.Sprintf("%s", res), nil
}

func scoreEnglishText(text string) float32 {
	//taking care of other white space character for conformity
	for i := 0; i < 20; i++ {
		text = strings.ReplaceAll(text, strings.Repeat("\n", i), " ")
		text = strings.ReplaceAll(text, strings.Repeat("\t", i), " ")
	}
	characters := strings.Split("abcdefghijklmnopqstuvwxzy", "")
	cw := float32(strings.Count(text, " "))
	for _, c := range characters {
		cw += float32(strings.Count(text, c))
	}

	return cw / float32(len(text))

}

func singlebyteXORcipher(s string) (string, string, error) {
	_, hexString, err := decodeHex(s)
	if err != nil {
		return "", "", err
	}
	charArr := make([]byte, len(hexString))
	result := ""
	countOfRes := float32(0)
	key := ""
	//all english letters
	for c := 0; c <= 255; c++ {
		for i := 0; i < len(hexString); i++ {
			charArr[i] = byte(c)
		}
		_, resString, err := xorHex(s, fmt.Sprintf("%X", charArr))

		if err != nil {
			fmt.Println("something not right", err, fmt.Sprintf("%s", charArr))
		}
		if scoreEnglishText(resString) > countOfRes {
			countOfRes = scoreEnglishText(resString)
			result = resString
			key = fmt.Sprintf("%s", charArr)
		}
	}
	return key, result, nil
}

func singleByteCharXor(msgs []string) (string, error) {

	countOfRes := float32(0)
	result := ""

	for _, msg := range msgs {
		_, decryptString, err := singlebyteXORcipher(msg)

		if err == nil && scoreEnglishText(decryptString) > countOfRes {

			countOfRes = scoreEnglishText(decryptString)
			result = decryptString
		}
	}
	return result, nil
}
func detectSingleCharacterXOR() (string, error) {
	content, err := ioutil.ReadFile("testfiles/detectsinglecharxor.in")
	if err != nil {
		return "", err
	}
	msgs := strings.Split(string(content), "\n")
	return singleByteCharXor(msgs)
}

func repeatingkeyXOR(text string, key []byte) []byte {

	textByte := []byte(text)
	res := make([]byte, len(textByte))
	keyIndex := 0
	for i := 0; i < len(textByte); i++ {
		keyIndex = keyIndex % len(key)
		res[i] = key[keyIndex] ^ textByte[i]
		keyIndex++
	}
	return res
}

func editDistance(a, b []byte) int {

	maxAB := max(len(a), len(b))

	padByteArr(&a, maxAB)
	padByteArr(&b, maxAB)
	res := 0
	for i := 0; i < maxAB; i++ {
		res += bits.OnesCount(uint(a[i] ^ b[i]))
	}

	return res

}

func findKeySize(f []byte) int {
	keySize := 0
	dist := float32(2 << 31)

	for curkeysize := 2; curkeysize <= 40; curkeysize++ {
		a := f[:curkeysize*4]
		b := f[curkeysize*4 : curkeysize*2*4]
		ed := editDistance(a, b)

		if float32(ed)/float32(curkeysize) < dist {
			dist = float32(ed) / float32(curkeysize)
			keySize = curkeysize
		}
	}
	return keySize
}

func breakRepeatingkeyXOR() error {
	file, err := ioutil.ReadFile("testfiles/repeatingkeyxor.in") // For read access.
	if err != nil {
		return err
	}
	f, err := base64.StdEncoding.DecodeString(string(file))

	if err != nil {
		return err
	}

	keySize := findKeySize(f)
	blocks := len(f) / keySize
	keys := make([]string, blocks)

	for i := 0; i < keySize; i++ {
		blk := make([]byte, blocks)
		k := 0
		for j := 0; j < len(f); j += keySize {
			if j+i < len(f) && k < blocks {
				blk[k] = f[j+i]
			}
			k++
		}
		key, _, err := singlebyteXORcipher(fmt.Sprintf("%x", blk))
		if err != nil {
			fmt.Println(err)
		}
		keys[i] = strings.Split(key, "")[0]
	}

	joinKey := ""
	for _, key := range keys {
		joinKey += key
	}
	fmt.Println(joinKey)
	fmt.Println(string(repeatingkeyXOR(string(f), []byte(joinKey))))

	return nil
}

func aesECBDecrypt(key string, ciphertext []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	res := make([]byte, len(ciphertext))
	for i := 0; i < len(res); i += aes.BlockSize {
		block.Decrypt(res[i:], ciphertext[i:])
	}

	return res
}

func detectAESECB() error {
	text, err := ioutil.ReadFile("testfiles/detectaesecbmode.in")
	if err != nil {
		return err
	}

	for _, ciphertext := range strings.Split(string(text), "\n") {
		checker := make(map[string]int)

		for i := 0; i < len(ciphertext); i += aes.BlockSize {

			k := ciphertext[i : i+aes.BlockSize]
			if _, ok := checker[k]; ok {
				fmt.Println(ciphertext)
				break
			}
			checker[k]++
		}
	}
	return nil

}
