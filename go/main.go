package main

import (
	"fmt"

	"github.com/fatih/color"
)

const maxPrime = 500
const minPrime = 70

func main() {
	//RSA

	color.Magenta("RSA test")
	var rsa RSA
	rsa = rsa.generateKeyPair()
	color.Blue("Public Key:")
	fmt.Println(rsa.PubK)
	color.Green("Private Key:")
	fmt.Println(rsa.PrivK)

	m1 := "Hi, trying RSA encryption"
	fmt.Println("m (original message): " + m1)

	c := rsa.encrypt(m1, rsa.PubK)
	color.Yellow("c (message encrypted):")
	fmt.Println(c)

	m1decrypted := rsa.decrypt(c, rsa.PrivK)
	color.Green("m (message decrypted):")
	fmt.Println(m1decrypted)

	color.Blue("-----\n\n")

	//RSA blind signature
	color.Magenta("RSA blind signature")
	r := 101
	msg := "hola"
	//convert msg to []int
	var m []int
	mBytes := []byte(msg)
	for _, byte := range mBytes {
		m = append(m, int(byte))
	}
	//blind
	mBlinded := rsa.sign(m, rsa.PubK, rsa.PrivK) //here the pubK and privK is the user's one
	fmt.Print("Message blinded': ")
	fmt.Println(mBlinded)
	//sign
	sigma := rsa.sign(mBlinded, rsa.PubK, rsa.PrivK) //here the privK will be the CA privK, not the m emmiter's one. The pubK is the user's one
	fmt.Print("Sigma': ")
	fmt.Println(sigma)
	//unblind
	signature := rsa.unblind(sigma, r, rsa.PubK)
	fmt.Print("Sigma (signature): ")
	fmt.Println(signature)

	color.Blue("-----\n\n")

	//RSA Homomorphic addition
	color.Magenta("RSA Homomorphic multiplication")
	m3 := 11
	m4 := 15
	fmt.Print("Message A: ")
	fmt.Print(m3)
	fmt.Print(", Message B: ")
	fmt.Println(m4)

	c3 := rsa.encryptInt(m3, rsa.PubK)
	c4 := rsa.encryptInt(m4, rsa.PubK)
	fmt.Print("message A encrypted: ")
	fmt.Print(c3)
	fmt.Print(", message B encrypted: ")
	fmt.Println(c4)

	c3c4 := rsa.homomorphicMultiplication(c3, c4, rsa.PubK)
	fmt.Print("Homomorphic multiplication A * B: ")
	fmt.Println(c3c4)
	m5decrypted := rsa.decryptInt(c3c4, rsa.PrivK)
	color.Green("Homomorphic multiplication result decrypted):")
	fmt.Println(m5decrypted)

	color.Blue("-----\n\n")

	//Paillier
	color.Magenta("Paillier test")

	var paillier Paillier
	paillier = paillier.generateKeyPair()
	fmt.Println(paillier)

	m2 := "Hi, here trying Paillier encryption"
	fmt.Println("m (original message): " + m2)

	c = paillier.encrypt(m2, paillier.PubK)
	color.Yellow("c (message encrypted):")
	fmt.Println(c)

	m2decrypted := paillier.decrypt(c, paillier.PubK, paillier.PrivK)
	color.Green("m (message decrypted):")
	fmt.Println(m2decrypted)

	color.Blue("-----\n\n")

	//Paillier Homomorphic addition
	color.Magenta("Paillier Homomorphic addition")
	m6 := 2
	m7 := 10
	fmt.Print("Message A: ")
	fmt.Print(m6)
	fmt.Print(", Message B: ")
	fmt.Println(m7)

	c6 := paillier.encryptInt(m6, paillier.PubK)
	c7 := paillier.encryptInt(m7, paillier.PubK)
	fmt.Print("message A encrypted: ")
	fmt.Print(c6)
	fmt.Print(", message B encrypted: ")
	fmt.Println(c7)

	c6c7 := paillier.homomorphicAddition(c6, c7, paillier.PubK)
	fmt.Print("Homomorphic addition A + B: ")
	fmt.Println(c6c7)
	m8decrypted := paillier.decryptInt(c6c7, paillier.PubK, paillier.PrivK)
	color.Green("Homomorphic addition result decrypted):")
	fmt.Println(m8decrypted)

	color.Blue("-----\n\n")

	color.Magenta("Secret Share test")
	var secretShare SecretShare
	shares, err := secretShare.create(2, 5, 17, "hola")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(shares)

	//generate sharesToUse
	var sharesToUse [][]string
	sharesToUse = append(sharesToUse, shares[0])
	sharesToUse = append(sharesToUse, shares[1])
	sharesToUse = append(sharesToUse, shares[3])
	secr := secretShare.LagrangeInterpolation(sharesToUse, 17)
	fmt.Println(secr)

}
