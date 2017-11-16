package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

type RSAPublicKey struct {
	E *big.Int `json:"e"`
	N *big.Int `json:"n"`
}
type RSAPrivateKey struct {
	D *big.Int `json:"d"`
	N *big.Int `json:"n"`
}

type RSA struct {
	PubK  RSAPublicKey
	PrivK RSAPrivateKey
}

func (rsa RSA) generateKeyPair() RSA {

	rand.Seed(time.Now().Unix())
	p := randPrime(minPrime, maxPrime)
	q := randPrime(minPrime, maxPrime)
	fmt.Print("p:")
	fmt.Println(p)
	fmt.Print("q:")
	fmt.Println(q)

	n := p * q
	phi := (p - 1) * (q - 1)
	e := 65537
	var pubK RSAPublicKey
	pubK.E = big.NewInt(int64(e))
	pubK.N = big.NewInt(int64(n))

	d := new(big.Int).ModInverse(big.NewInt(int64(e)), big.NewInt(int64(phi)))

	var privK RSAPrivateKey
	privK.D = d
	privK.N = big.NewInt(int64(n))

	rsa.PubK = pubK
	rsa.PrivK = privK
	return rsa
}
func (rsa RSA) encrypt(m string, pubK RSAPublicKey) []int {
	var c []int
	mBytes := []byte(m)
	for _, byte := range mBytes {
		c = append(c, rsa.encryptInt(int(byte), pubK))
	}
	return c
}
func (rsa RSA) decrypt(c []int, privK RSAPrivateKey) string {
	var m string
	var mBytes []byte
	for _, indC := range c {
		mBytes = append(mBytes, byte(rsa.decryptInt(indC, privK)))
	}
	m = string(mBytes)
	return m
}
func (rsa RSA) encryptInt(char int, pubK RSAPublicKey) int {
	charBig := big.NewInt(int64(char))
	Me := charBig.Exp(charBig, pubK.E, nil)
	c := Me.Mod(Me, pubK.N)
	return int(c.Int64())
}
func (rsa RSA) decryptInt(val int, privK RSAPrivateKey) int {
	valBig := big.NewInt(int64(val))
	Cd := valBig.Exp(valBig, privK.D, nil)
	m := Cd.Mod(Cd, privK.N)
	return int(m.Int64())
}
