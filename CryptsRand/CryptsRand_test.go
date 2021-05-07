package CryptsRand

import "testing"

func TestPseudoRandom(t *testing.T) {
	for i:=0; i<100;i+=1 {
		PseudoRandom()
	}
}

func TestUniformRandom(t *testing.T) {
	for i:=0; i<100; i+=1 {
		UniformRandom()
	}
}