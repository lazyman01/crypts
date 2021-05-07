package CryptsMd5

import "crypto/md5"

func Md532(data string) []byte  {
	h := md5.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}
