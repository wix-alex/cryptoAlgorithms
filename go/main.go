package main

import (
	"fmt"

	"github.com/fatih/color"
)

const maxPrime = 500
const minPrime = 70

func main() {
	//RSA
	/*
		color.Yellow("RSA test")
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

		fmt.Println("-----")
		//Paillier
		color.Yellow("Paillier test")

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
	*/
	color.Yellow("Secret Share test")
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
	r := secretShare.LagrangeInterpolation(sharesToUse, 17)
	fmt.Println(r)
}
