package main

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type SecretShare struct {
	Secret string `json:"secret"`
}

/*
t: polynomial will be t-1 degree
n: number of shares
k: sand alone value of the polynomial
p: polynomial will be evaluated at mod p
*/
func (secretShare SecretShare) create(t int, n int, k int, p int, s string) ([]string, error) {
	if k > p {
		return nil, errors.New("Error: need k<p")
	}

	secret := secretShare.stringArrayToBigIntArray(strings.Split(s, ""))
	fmt.Println(secret)

	//generate the basePolynomial
	var basePolynomial []*big.Int
	basePolynomial = append(basePolynomial, big.NewInt(int64(k)))
	for i := 0; i < len(secret)-1; i++ {
		randBigInt := big.NewInt(int64(randPrime(minPrime, maxPrime)))

		basePolynomial = append(basePolynomial, randBigInt)
	}
	fmt.Println(basePolynomial)

	//next 6 lines are to test
	basePolynomial = []*big.Int{
		big.NewInt(int64(k)),
		big.NewInt(int64(8)),
		big.NewInt(int64(7)),
	}
	fmt.Println(basePolynomial)

	//calculate shares, based on the basePolynomial
	var shares []*big.Int
	for i := 1; i < n+1; i++ {
		fmt.Print("i: ")
		fmt.Println(i)
		var pResultMod *big.Int
		pResult := big.NewInt(int64(0))
		for x, polElem := range basePolynomial {
			fmt.Print("i: ")
			fmt.Print(i)
			fmt.Print(", x: ")
			fmt.Print(x)
			if x == 0 {
				pResult = pResult.Add(pResult, polElem)
				fmt.Println("")
			} else {
				iBigInt := big.NewInt(int64(i))
				xBigInt := big.NewInt(int64(x))
				iPowed := iBigInt.Exp(iBigInt, xBigInt, nil)
				fmt.Print(" iPowed: ")
				fmt.Print(iPowed)
				fmt.Print(" polElem: ")
				fmt.Print(polElem)
				currElem := iPowed.Mul(iPowed, polElem)
				fmt.Print(" currElem: ")
				fmt.Println(currElem)
				pResult = pResult.Add(pResult, currElem)
				pResultMod = pResult.Mod(pResult, big.NewInt(int64(p)))
			}
		}
		fmt.Println(pResultMod)
		//pResult = big.NewInt(int64(0))
		shares = append(shares, pResultMod)
	}
	fmt.Println(shares)
	fmt.Println(shares)
	result := secretShare.bigIntArrayToStringArray(shares)
	return result, nil
}

func (secretShare SecretShare) stringArrayToBigIntArray(s []string) []*big.Int {
	var r []*big.Int
	//sBytes := []byte(s)
	for _, sElem := range s {
		b := []byte(sElem)[0]
		bBig := big.NewInt(int64(b))
		r = append(r, bBig)
	}
	return r
}
func (secretShare SecretShare) bigIntArrayToStringArray(b []*big.Int) []string {
	var r []string
	for _, bigint := range b {
		r = append(r, bigint.String())
	}
	return r
}

func (secretShare SecretShare) LagrangeInterpolation(shares []string, p int) *big.Int {

	result := big.NewInt(int64(0))
	/*
		sharesBigInt := secretShare.stringArrayToBigIntArray(shares)
		var lambda *big.Int
		for i, share := range sharesBigInt {

		}
	*/

	resultMod := result.Mod(result, big.NewInt(int64(p)))
	return resultMod
}

//l is the others values != i
//j is the share i
func (secretShare SecretShare) LagrangeBasis(l *big.Int, j *big.Int) *big.Int {
	l_j := new(big.Int).Sub(l, j)
	sublambda := new(big.Int).Div(l, l_j)
	return sublambda
}
