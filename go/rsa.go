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

func (rsa RSA) blind(m []int, r int, pubK RSAPublicKey, privK RSAPrivateKey) []int {
	var mBlinded []int
	rBigInt := big.NewInt(int64(r))
	for i := 0; i < len(m); i++ {
		mBigInt := big.NewInt(int64(m[i]))
		rE := new(big.Int).Exp(rBigInt, pubK.E, nil)
		mrE := new(big.Int).Mul(mBigInt, rE)
		mrEmodN := new(big.Int).Mod(mrE, privK.N)
		mBlinded = append(mBlinded, int(mrEmodN.Int64()))
	}
	return mBlinded
}

func (rsa RSA) sign(m []int, pubK RSAPublicKey, privK RSAPrivateKey) []int {
	var r []int
	for i := 0; i < len(m); i++ {
		mBigInt := big.NewInt(int64(m[i]))
		sigma := new(big.Int).Exp(mBigInt, privK.D, pubK.N)
		r = append(r, int(sigma.Int64()))
	}
	return r
}
func (rsa RSA) unblind(sigma []int, r int, pubK RSAPublicKey) []int {
	var signature []int
	rBigInt := big.NewInt(int64(r))
	for i := 0; i < len(sigma); i++ {
		sigmaBigInt := big.NewInt(int64(sigma[i]))
		r1 := new(big.Int).Exp(rBigInt, big.NewInt(int64(-1)), nil)
		sigmar := new(big.Int).Mul(sigmaBigInt, r1)
		sig := new(big.Int).Mod(sigmar, pubK.N)
		signature = append(signature, int(sig.Int64()))
	}
	return signature
}

func (rsa RSA) homomorphicMultiplication(c1 int, c2 int, pubK RSAPublicKey) int {
	c1BigInt := big.NewInt(int64(c1))
	c2BigInt := big.NewInt(int64(c2))
	c1c2 := new(big.Int).Mul(c1BigInt, c2BigInt)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	d := new(big.Int).Mod(c1c2, n2)
	r := int(d.Int64())
	return r
}
