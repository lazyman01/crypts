package CryptsRC4

import (
	"crypto/rc4"
)
//rc4加密
// strToEncode 需要加密的字符串； keyStr 密钥
func RC4Encode(bytesToEncode []byte, key []byte) []byte {
	rc4obj, _ := rc4.NewCipher(key) //返回 Cipher
	plaintext := make([]byte, len(bytesToEncode)) //
	rc4obj.XORKeyStream(plaintext, bytesToEncode)
	//XORKeyStream方法将src的数据与秘钥生成的伪随机位流取XOR并写入dst。
	//plaintext就是你加密的返回过来的结果了，注意：plaintext为base-16 编码的字符串，每个字节使用2个字符表示 必须格式化成字符串
	//stringinf := fmt.Sprintf("%x\n", plaintext) //转换字符串
	return plaintext
}

//rc4解密
func RC4Decode(strToDecode []byte, key []byte, ) []byte  {
	maxLen := 1024
	decodeBytes := make([]byte, maxLen)
	cipher, _ := rc4.NewCipher(key) // 切记：这里不能重用cipher1，必须重新生成新的
	cipher.XORKeyStream(decodeBytes, strToDecode)
	return decodeBytes
}
