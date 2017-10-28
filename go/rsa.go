package main

import (
	"math/big"
	"math/rand"
	"time"
)

type PublicKey struct {
	E *big.Int `json:"e"`
	N *big.Int `json:"n"`
}
type PrivateKey struct {
	D *big.Int `json:"d"`
	N *big.Int `json:"n"`
}

const maxPrime = 1000

func generateKeyPair() (PublicKey, PrivateKey) {

	rand.Seed(time.Now().Unix())
	p := randPrime(0, maxPrime)
	q := randPrime(0, maxPrime)

	n := p * q
	phi := (p - 1) * (q - 1)
	e := 65537
	var pubK PublicKey
	pubK.E = big.NewInt(int64(e))
	pubK.N = big.NewInt(int64(n))

	d := new(big.Int).ModInverse(big.NewInt(int64(e)), big.NewInt(int64(phi)))

	var privK PrivateKey
	privK.D = d
	privK.N = big.NewInt(int64(n))

	return pubK, privK
}
func encryptM(m string, pubK PublicKey) []int {
	var c []int
	mBytes := []byte(m)
	for _, byte := range mBytes {
		c = append(c, encrypt(int(byte), pubK))
	}
	return c
}
func decryptC(c []int, privK PrivateKey) string {
	var m string
	var mBytes []byte
	for _, indC := range c {
		mBytes = append(mBytes, byte(decrypt(indC, privK)))
	}
	m = string(mBytes)
	return m
}
func encrypt(char int, pubK PublicKey) int {
	charBig := big.NewInt(int64(char))
	Me := charBig.Exp(charBig, pubK.E, nil)
	c := Me.Mod(Me, pubK.N)
	return int(c.Int64())
}
func decrypt(val int, privK PrivateKey) int {
	valBig := big.NewInt(int64(val))
	Cd := valBig.Exp(valBig, privK.D, nil)
	m := Cd.Mod(Cd, privK.N)
	return int(m.Int64())
}
