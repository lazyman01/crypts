package CryptsECDSA

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	//明文
	message := []byte("Hello world")

	//获取私钥
	key, err := NewSigningKey()
	if err != nil {
		return
	}

	//用私钥对明文进行签名
	signature, err := Sign(message, key)

	fmt.Printf("签名后：%x\n", signature)
	if err != nil {
		return
	}

	//用公钥对签名进行验证，确认签名是否是对当前明文的有效
	if !Verify(message, signature, &key.PublicKey) {
		fmt.Println("验证失败！")
		return
	}else{
		fmt.Println("验证成功！")
	}
}
