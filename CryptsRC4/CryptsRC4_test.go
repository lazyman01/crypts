package CryptsRC4

import (
	"fmt"
	"testing"
)

func TestRC4Encode(t *testing.T) {
	//加密
	rawStr := "hellowordwefwef232r43f"
	keyStr := "123456"
	encodeBytes := RC4Encode([]byte(rawStr), []byte(keyStr))
	fmt.Println("ecodestr = ", fmt.Sprintf("%x\n", encodeBytes))

	//解密
	decodeBytes := RC4Decode(encodeBytes, []byte(keyStr))
	fmt.Println("decodeBytes = ", string(decodeBytes))
}
