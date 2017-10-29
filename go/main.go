package main

import (
	"fmt"

	"github.com/fatih/color"
)

const maxPrime = 500
const minPrime = 70

func main() {
	//RSA
	color.Yellow("RSA test")
	var rsa RSA
	rsa = rsa.generateKeyPair()
	color.Blue("Public Key:")
	fmt.Println(rsa.PubK)
	color.Green("Private Key:")
	fmt.Println(rsa.PrivK)

	m := "Hi, trying RSA encryption"
	fmt.Println("m (original message): " + m)

	c := rsa.encryptM(m, rsa.PubK)
	color.Yellow("c (message encrypted):")
	fmt.Println(c)

	m2 := rsa.decryptC(c, rsa.PrivK)
	color.Green("m (message decrypted):")
	fmt.Println(m2)

	fmt.Println("-----")
	//Paillier
	color.Yellow("Paillier test")

	var paillier Paillier
	paillier = paillier.generateKeyPair()
	fmt.Println(paillier)

	m = "Hi, here trying Paillier encryption"
	fmt.Println("m (original message): " + m)

	c = paillier.encryptM(m, paillier.PubK)
	color.Yellow("c (message encrypted):")
	fmt.Println(c)

	m2 = paillier.decryptC(c, paillier.PubK, paillier.PrivK)
	color.Green("m (message decrypted):")
	fmt.Println(m2)
}
