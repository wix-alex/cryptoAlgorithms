package main

import (
	"fmt"

	"github.com/fatih/color"

	ownPaillier "./ownPaillier"
	ownRSA "./ownRSA"
	ownSecretShare "./ownSecretShare"
)

func main() {
	//RSA

	color.Magenta("RSA test")
	var rsa ownRSA.RSA
	rsa = ownRSA.GenerateKeyPair()
	color.Blue("Public Key:")
	fmt.Println(rsa.PubK)
	color.Green("Private Key:")
	fmt.Println(rsa.PrivK)

	/*
		//using the whole message to bigint
		m1 := "Hi"
		fmt.Println("m (original message): " + m1)
		m1Bytes := []byte(m1)
		fmt.Println(m1Bytes)
		m1BigInt := new(big.Int).SetBytes(m1Bytes)
		fmt.Println(m1BigInt)
		c := rsa.encryptBigInt(m1BigInt, rsa.PubK)
		fmt.Print("c: ")
		fmt.Println(c)

		m1decryptedBigInt := rsa.decryptBigInt(c, rsa.PrivK)
		m1decryptedBytes := m1decryptedBigInt.Bytes()
		fmt.Println(m1decryptedBytes)
		m1decrypted := string(m1decryptedBytes)
	*/

	m1 := "Hi, trying RSA encryption"
	c := ownRSA.Encrypt(m1, rsa.PubK)
	color.Yellow("c (message encrypted):")
	fmt.Println(c)

	m1decrypted := ownRSA.Decrypt(c, rsa.PrivK)
	color.Green("m (message decrypted):")
	fmt.Println(m1decrypted)

	color.Blue("\n-----\n\n\n")

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
	mBlinded := ownRSA.BlindSign(m, rsa.PubK, rsa.PrivK) //here the pubK and privK is the user's one
	fmt.Print("Message blinded': ")
	fmt.Println(mBlinded)
	//sign
	sigma := ownRSA.BlindSign(mBlinded, rsa.PubK, rsa.PrivK) //here the privK will be the CA privK, not the m emmiter's one. The pubK is the user's one
	fmt.Print("Sigma': ")
	fmt.Println(sigma)
	//unblind
	signature := ownRSA.Unblind(sigma, r, rsa.PubK)
	fmt.Print("Sigma (signature): ")
	fmt.Println(signature)

	color.Blue("\n-----\n\n\n")

	//RSA Homomorphic addition
	color.Magenta("RSA Homomorphic multiplication")
	m3 := 11
	m4 := 15
	fmt.Print("Message A: ")
	fmt.Print(m3)
	fmt.Print(", Message B: ")
	fmt.Println(m4)

	c3 := ownRSA.EncryptInt(m3, rsa.PubK)
	c4 := ownRSA.EncryptInt(m4, rsa.PubK)
	fmt.Print("message A encrypted: ")
	fmt.Print(c3)
	fmt.Print(", message B encrypted: ")
	fmt.Println(c4)

	c3c4 := ownRSA.HomomorphicMultiplication(c3, c4, rsa.PubK)
	fmt.Print("Homomorphic multiplication A * B: ")
	fmt.Println(c3c4)
	m5decrypted := ownRSA.DecryptInt(c3c4, rsa.PrivK)
	color.Green("Homomorphic multiplication result decrypted):")
	fmt.Println(m5decrypted)

	color.Blue("\n-----\n\n\n")

	//Paillier
	color.Magenta("Paillier test")

	var paillier ownPaillier.Paillier
	paillier = ownPaillier.GenerateKeyPair()
	fmt.Println(paillier)

	m2 := "Hi, here trying Paillier encryption"
	fmt.Println("m (original message): " + m2)

	c2 := ownPaillier.Encrypt(m2, paillier.PubK)
	color.Yellow("c (message encrypted):")
	fmt.Println(c)

	m2decrypted := ownPaillier.Decrypt(c2, paillier.PubK, paillier.PrivK)
	color.Green("m (message decrypted):")
	fmt.Println(m2decrypted)

	color.Blue("\n-----\n\n\n")

	//Paillier Homomorphic addition
	color.Magenta("Paillier Homomorphic addition")
	m6 := 2
	m7 := 10
	fmt.Print("Message A: ")
	fmt.Print(m6)
	fmt.Print(", Message B: ")
	fmt.Println(m7)

	c6 := ownPaillier.EncryptInt(m6, paillier.PubK)
	c7 := ownPaillier.EncryptInt(m7, paillier.PubK)
	fmt.Print("message A encrypted: ")
	fmt.Print(c6)
	fmt.Print(", message B encrypted: ")
	fmt.Println(c7)

	c6c7 := ownPaillier.HomomorphicAddition(c6, c7, paillier.PubK)
	fmt.Print("Homomorphic addition A + B: ")
	fmt.Println(c6c7)
	m8decrypted := ownPaillier.DecryptInt(c6c7, paillier.PubK, paillier.PrivK)
	color.Green("Homomorphic addition result decrypted):")
	fmt.Println(m8decrypted)

	color.Blue("\n-----\n\n\n")

	color.Magenta("Secret Share test")
	//var secretShare ownSecretShare.SecretShare

	//ownSecretShare.Create(needed shares, num of shares, p_mod, text)
	k := 1234
	fmt.Print("original secret: ")
	fmt.Println(k)
	shares, err := ownSecretShare.Create(3, 6, 1613, k)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(shares)

	//generate sharesToUse
	var sharesToUse [][]string
	sharesToUse = append(sharesToUse, shares[2])
	sharesToUse = append(sharesToUse, shares[1])
	sharesToUse = append(sharesToUse, shares[0])
	secr := ownSecretShare.LagrangeInterpolation(sharesToUse, 1613)
	fmt.Print("secret result: ")
	fmt.Println(secr)

}
