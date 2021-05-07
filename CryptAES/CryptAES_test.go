package CryptAES

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAes128Encrypt(t *testing.T) {
	SKey        := "2e32datg5ytzbefr"
	IvParameter := "0srjopxmwe3mfhwo"
	origData 	:= "Hello World！"
	enBytes, err := Aes128Encrypt([]byte(origData), []byte(SKey), []byte(IvParameter))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("enBytes=", Bytes2Hex(enBytes))
}

func TestAes128Decrypt(t *testing.T) {
	SKey        := "2e32datg5ytzbefr"
	IvParameter := "0srjopxmwe3mfhwo"
	enBytes := Hex2Bytes("e65cd4e5851c52a433f8b27193862e41")
	origData, err := Aes128Decrypt(enBytes, []byte(SKey), []byte(IvParameter))
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("origData=", string(origData))
}

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}
func Bytes2Hex(d []byte) string {
	return hex.EncodeToString(d)
}