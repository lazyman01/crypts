package CryptsSha256

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
)

func TestHmacSha256(t *testing.T) {
	ToEncode := "wef34r2d23r43g5g4tuh67j78lo98;l09ytbrefvcwdcqwex"
	key := "123456"
	encodeBytes := HmacSha256([]byte(ToEncode), []byte(key))
	encodeB64 := base64.StdEncoding.EncodeToString(encodeBytes)
	resultBase64 :="qdWsgFDNZCcNz1r8zaDyvOZfqyZ07/wOyEGemCQx+rQ="
	if !reflect.DeepEqual(encodeB64, resultBase64) {
		t.Error("not equal: ")
		fmt.Println(encodeB64)
		fmt.Println(resultBase64)
	}
}
