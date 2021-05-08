package CryptsPBKDF2

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"os"
)

/*
官方提供 “golang.org/x/crypto/pbkdf2″ 包封装此加密算法
func Key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte
	password：用户密码，加密的原材料，以 []byte 传入；
	salt：随机盐，以 []byte 传入；
	iter：迭代次数，次数越多加密与破解所需时间越长;
	keylen：期望得到的密文长度;
	Hash：加密所用的 hash 函数，默认使用为 sha1，这里我们使用sha256
*/

func CryptsPBKDF2()  {
	//生成随机盐
	salt := make([]byte, 32)
	fmt.Println(salt)
	_, err := rand.Read(salt)
	checkErr(err)
	fmt.Println(salt)
	//生成密文
	dk := pbkdf2.Key([]byte("mimashi1323"), salt, 1, 32, sha256.New)
	fmt.Println(dk)
	fmt.Println(hex.EncodeToString(dk))
}

func checkErr(err error){
	if err != nil {
		fmt.Fprintln(os.Stderr,"Fatal error:",err.Error())
		os.Exit(1)
	}
}