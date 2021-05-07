package CryptsDes

import (
	"fmt"
	"testing"
)

func TestMyDesEncrypt(t *testing.T) {
	key :=[]byte("12345698")
	origData := "hello world!"
	iv := []byte("43218765")
	encodeBytes, err := DesEncrypt([]byte(origData), key, iv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Bytes2Hex(encodeBytes))
}

func TestDecrptogDES(t *testing.T) {
	key :=[]byte("aFefbFca")
	origData := "Hello World!"
	iv := []byte("Passw0rd")
	encodestr := "e529b5bbd2e0e3fb62eeb88ba6a22367"

	//encodeBytes, _ := base64.StdEncoding.DecodeString(encodeB64)
	data,err := DesDecrypt([]byte(encodestr), key, iv)
	fmt.Println(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(data) != origData {
		t.Error("not equal")
	}
}

