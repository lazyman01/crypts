package CryptsSha512

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
)

func TestHmacSha512(t *testing.T) {
	ToEncode := "wef34r2d23r43g5g4tuh67j78lo98;l09ytbrefvcwdcqwex"
	key := "123456"
	encodeBytes := HmacSha512([]byte(ToEncode), []byte(key))
	encodeB64 := base64.StdEncoding.EncodeToString(encodeBytes)
	resultBase64 :="S/zLfY2WQLiCvQ0TUNNoPKn6aYgiUksAY3w7yiIz4cz4b91bpETAp9g3mx+fgdQVa8W2M8zXs9wkdJoB5DlM3w=="
	if !reflect.DeepEqual(encodeB64, resultBase64) {
		t.Error("not equal: ")
		fmt.Println(encodeB64)
		fmt.Println(resultBase64)
	}
}

