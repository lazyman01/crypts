package CryptsECDSA

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func NewSigningKey() (*ecdsa.PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return key, err
}

// 用私钥对明文进行签名
func Sign(data []byte, privkey *ecdsa.PrivateKey) ([]byte, error) {
	// 对明文进行sha256散列，生成一个长度为32的字节数组
	digest := sha256.Sum256(data)

	// 通过椭圆曲线方法对散列后的明文进行签名，返回两个big.int类型的大数
	r, s, err := ecdsa.Sign(rand.Reader, privkey, digest[:])
	if err != nil {
		return nil, err
	}
	//将大数转换成字节数组，并拼接起来，形成签名
	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

// 通过公钥验证签名
func Verify(data, signature []byte, pubkey *ecdsa.PublicKey) bool {
	// 将明文转换成字节数组
	digest := sha256.Sum256(data)

	//声明两个大数r，s
	r := big.Int{}
	s := big.Int{}
	//将签名平均分割成两部分切片，并将切片转换成*big.int类型
	sigLen := len(signature)
	r.SetBytes(signature[:(sigLen / 2)])
	s.SetBytes(signature[(sigLen / 2):])

	//通过公钥对得到的r，s进行验证
	return ecdsa.Verify(pubkey, digest[:], &r, &s)
}