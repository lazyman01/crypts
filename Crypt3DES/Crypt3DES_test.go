package Crypt3DES

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func main()  {
	fmt.Println("hello world")
	var key=[]byte("123456789012345678901239")
	var encirtcode =TripleEncrypt([]byte("hello world"),key)
	var decryptcode=TrileDesDecrypt(encirtcode,key)
	fmt.Printf("%x\n",encirtcode)
	fmt.Println(string(decryptcode))

}
func TestTrileDesDecrypt(t *testing.T) {
	var key=[]byte("123456789012345678901239")
	var encirtcode =TripleEncrypt([]byte("hello world"),key)
	fmt.Printf("%x\n",encirtcode)
}

func TestTrileDesDecrypt2(t *testing.T) {
	var key=[]byte("123456789012345678901239")
	encodeStr := Hex2Bytes("d2cd8901a667746605f7ce753e2480eb")
	var decryptcode=TrileDesDecrypt(encodeStr, key)
	fmt.Println(string(decryptcode))

}

func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}