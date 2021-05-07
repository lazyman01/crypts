package CryptRSA

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	data, _ := RsaEncrypt([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))
}

func TestRsaDecrypt(t *testing.T) {
	data, _ := RsaEncrypt([]byte("hello world"))
	origData, _ := RsaDecrypt(data)
	fmt.Println(string(origData))
}