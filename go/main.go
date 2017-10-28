package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	pubK, privK := generateKeyPair()
	color.Blue("Public Key:")
	fmt.Println(pubK)
	color.Green("Private Key:")
	fmt.Println(privK)

	m := "hola provant això"
	fmt.Println("m (original message): " + m)

	c := encryptM(m, pubK)
	color.Yellow("c (message encrypted):")
	fmt.Println(c)

	m2 := decryptC(c, privK)
	color.Green("m (message decrypted):")
	fmt.Println(m2)
}
