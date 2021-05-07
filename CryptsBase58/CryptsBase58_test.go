package CryptsBase58

import (
	"fmt"
	"testing"
)

func TestBase58Encode(t *testing.T) {
	org := []byte("qwerty")
	fmt.Println(string(org))

	ReverseBytes(org)

	fmt.Println(string(org))
	enBytes := Base58Encode([]byte("hello jonson"))
	//测试编码
	fmt.Printf("%s\n",string(enBytes))


	deBytes := Base58Decode(enBytes)
	fmt.Printf("%s\n", deBytes)
}