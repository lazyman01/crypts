package CryptsMd5

import (
	"reflect"
	"testing"
)

func TestMd532(t *testing.T) {
	var (
		md5RawStr = "123456789abc"
		md5Hash = "d266f2f31cf903c870027659030e967e"
	)

	if reflect.DeepEqual(Md532(md5RawStr), []byte(md5Hash)) {
		t.Errorf("Md532(%s) = %s , expecte: %s", md5RawStr, Md532(md5RawStr), md5Hash)
	}
}