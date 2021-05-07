package CryptsBase64

import (
	"reflect"
	"testing"
)

func TestBase64(t *testing.T)  {
	var (
		b64RawStr = "123456789abc"
		b64Str = "MTIzNDU2Nzg5YWJj"
	)

	if Base64Encode([]byte(b64RawStr)) != b64Str {
		t.Errorf("Base64Encode(%s) = %s , expecte: %s", b64RawStr, Base64Encode([]byte(b64RawStr)), b64Str)
	}
	decodeStr, err := Base64Decode(b64Str)
	if err != nil {
		t.Error("Base64Decode: ", err)
	}
	if !reflect.DeepEqual(decodeStr, []byte(b64RawStr))  {
		t.Errorf("Base64Decode(%s) = %s , expecte: %s", b64RawStr, Base64Encode([]byte(b64RawStr)), b64Str)
	}
}
