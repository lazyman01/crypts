package CryptsRand

import (
	"crypto/rand"
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
)

//伪随机数
func PseudoRandom()  {
	fmt.Println("伪随机数: ",mrand.Int())
}

//真随机数
func UniformRandom()  {
	result,_ := crand.Int(rand.Reader, big.NewInt(100000000000000000))
	fmt.Println("真随机数: ",result)
}
