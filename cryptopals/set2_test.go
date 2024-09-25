package main

import (
	// "encoding/base64"
	// "fmt"
	// "io/ioutil"
	"crypto/aes"
	"log"
	"testing"
)

func TestPKCSPadding(t *testing.T) {
	unPadded20String := "YELLOW SUBMARINE"
	expected20Result := "YELLOW SUBMARINE\x04\x04\x04\x04"
	er := PKCSPadding([]byte(unPadded20String), 20)

	if string(er) != expected20Result {
		t.Errorf("%s does match %s", er, expected20Result)
	}

}

func TestCBCmode(t *testing.T) {

	msg := string(PKCSPadding([]byte("hello? how is the thing?"),aes.BlockSize))

	encryptedMsg := aesCBCEncrypt("YELLOW SUBMARINE", []byte(msg))
	decryptedMsg := aesCBCDecrypt("YELLOW SUBMARINE", encryptedMsg)

	if string(decryptedMsg) != msg {
		log.Println("hello",string(decryptedMsg),"end")
		log.Println("hello3",string(encryptedMsg),"end")
		t.Errorf("Encryption didn't work ")
	}
}
