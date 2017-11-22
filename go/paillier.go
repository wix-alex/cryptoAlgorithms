package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

type PaillierPublicKey struct {
	N *big.Int `json:"n"`
	G *big.Int `json:"g"`
}
type PaillierPrivateKey struct {
	Lambda *big.Int `json:"lambda"`
	Mu     *big.Int `json:"mu"`
}

type Paillier struct {
	PubK  PaillierPublicKey
	PrivK PaillierPrivateKey
}

func (paillier Paillier) generateKeyPair() Paillier {

	rand.Seed(time.Now().Unix())
	p := big.NewInt(int64(randPrime(minPrime, maxPrime)))
	q := big.NewInt(int64(randPrime(minPrime, maxPrime)))
	fmt.Print("p: ")
	fmt.Println(p)
	fmt.Print("q: ")
	fmt.Println(q)
	pq := new(big.Int).Mul(p, q)
	p1q1 := big.NewInt((p.Int64() - 1) * (q.Int64() - 1))
	gcd := new(big.Int).GCD(nil, nil, pq, p1q1)
	fmt.Print("gcd comprovation: ")
	fmt.Println(gcd)

	n := new(big.Int).Mul(p, q)
	lambda := big.NewInt(int64(paillier.lcm(float64(p.Int64())-1, float64(q.Int64())-1)))
	fmt.Print("lambda: ")
	fmt.Println(lambda)

	//g generation
	alpha := big.NewInt(int64(randInt(0, int(n.Int64()))))
	beta := big.NewInt(int64(randInt(0, int(n.Int64()))))
	alphan := new(big.Int).Mul(alpha, n)
	alphan1 := new(big.Int).Add(alphan, big.NewInt(1))
	betaN := new(big.Int).Exp(beta, n, nil)
	ab := new(big.Int).Mul(alphan1, betaN)
	n2 := new(big.Int).Mul(n, n)
	g := new(big.Int).Mod(ab, n2)
	//in some Paillier implementations use this:
	//g = new(big.Int).Add(n, big.NewInt(1))
	fmt.Print("g: ")
	fmt.Println(g)

	paillier.PubK.N = n
	paillier.PubK.G = g

	//mu generation
	Glambda := new(big.Int).Exp(g, lambda, nil)
	u := new(big.Int).Mod(Glambda, n2)
	L := paillier.L(u, n)
	mu := new(big.Int).ModInverse(L, n)

	paillier.PrivK.Lambda = lambda
	paillier.PrivK.Mu = mu

	return paillier
}
func (paillier Paillier) lcm(a, b float64) float64 {
	r := (a * b) / float64(gcd(int(a), int(b)))
	return r

}
func (paillier Paillier) L(u *big.Int, n *big.Int) *big.Int {
	u1 := new(big.Int).Sub(u, big.NewInt(1))
	L := new(big.Int).Div(u1, n)
	return L
}
func (paillier Paillier) encrypt(m string, pubK PaillierPublicKey) []int {
	var c []int
	mBytes := []byte(m)
	fmt.Print("mBytes: ")
	fmt.Println(mBytes)
	for _, byte := range mBytes {
		c = append(c, paillier.encryptInt(int(byte), pubK))
	}
	return c
}
func (paillier Paillier) decrypt(c []int, pubK PaillierPublicKey, privK PaillierPrivateKey) string {
	var m string
	var mBytes []byte
	for _, indC := range c {
		mBytes = append(mBytes, byte(paillier.decryptInt(indC, pubK, privK)))
	}
	m = string(mBytes)
	return m
}
func (paillier Paillier) encryptInt(char int, pubK PaillierPublicKey) int {
	m := big.NewInt(int64(char))
	gM := new(big.Int).Exp(pubK.G, m, nil)
	r := big.NewInt(int64(randInt(0, int(pubK.N.Int64()))))
	rN := new(big.Int).Exp(r, pubK.N, nil)
	gMrN := new(big.Int).Mul(gM, rN)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	c := new(big.Int).Mod(gMrN, n2)
	return int(c.Int64())
}
func (paillier Paillier) decryptInt(val int, pubK PaillierPublicKey, privK PaillierPrivateKey) int {
	c := big.NewInt(int64(val))
	cLambda := new(big.Int).Exp(c, privK.Lambda, nil)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	u := new(big.Int).Mod(cLambda, n2)
	L := paillier.L(u, pubK.N)
	LMu := new(big.Int).Mul(L, privK.Mu)
	m := new(big.Int).Mod(LMu, pubK.N)
	return int(m.Int64())
}

/*func (paillier Paillier) homomorphicAddition(c1 []int, c2 []int, pubK PaillierPublicKey) []int {
	var r []int
	for i := 0; i < len(c1); i++ {
		c1BigInt := big.NewInt(int64(c1[i]))
		c2BigInt := big.NewInt(int64(c2[i]))
		c1c2 := new(big.Int).Mul(c1BigInt, c2BigInt)
		n2 := new(big.Int).Mul(pubK.N, pubK.N)
		d := new(big.Int).Mod(c1c2, n2)
		r = append(r, int(d.Int64()))
	}
	return r
}*/

func (paillier Paillier) homomorphicAddition(c1 int, c2 int, pubK PaillierPublicKey) int {
	c1BigInt := big.NewInt(int64(c1))
	c2BigInt := big.NewInt(int64(c2))
	c1c2 := new(big.Int).Mul(c1BigInt, c2BigInt)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	d := new(big.Int).Mod(c1c2, n2)
	r := int(d.Int64())
	return r
}
