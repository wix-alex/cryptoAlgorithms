package ownSecretShare

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	ownPrime "../ownPrime"
)

type SecretShare struct {
	Secret string `json:"secret"`
}

/*
t: min needed shares, polynomial will be t-1 degree
n: number of shares
k: sand alone value of the polynomial
p: polynomial will be evaluated at mod p
*/
func Create(t int, n int, p int, k int) ([][]string, error) {
	//next line is for developing test
	//k = 1234

	if k > p {
		return nil, errors.New("Error: need k<p")
	}

	//secret := CharactersArrayToBigIntArray(strings.Split(s, ""))

	//generate the basePolynomial
	var basePolynomial []*big.Int
	basePolynomial = append(basePolynomial, big.NewInt(int64(k)))
	//for i := 0; i < len(secret)-1; i++ {
	for i := 0; i < t-1; i++ {
		randBigInt := big.NewInt(int64(ownPrime.RandPrime(ownPrime.MinPrime, ownPrime.MaxPrime)))

		basePolynomial = append(basePolynomial, randBigInt)
	}

	//next 6 lines are to test
	/*basePolynomial = []*big.Int{
		big.NewInt(int64(k)),
		big.NewInt(int64(166)),
		big.NewInt(int64(94)),
	}*/

	//calculate shares, based on the basePolynomial
	var shares []*big.Int
	for i := 1; i < n+1; i++ {
		var pResultMod *big.Int
		pResult := big.NewInt(int64(0))
		for x, polElem := range basePolynomial {
			if x == 0 {
				pResult = pResult.Add(pResult, polElem)
			} else {
				iBigInt := big.NewInt(int64(i))
				xBigInt := big.NewInt(int64(x))
				iPowed := iBigInt.Exp(iBigInt, xBigInt, nil)
				currElem := iPowed.Mul(iPowed, polElem)
				pResult = pResult.Add(pResult, currElem)
				pResultMod = pResult.Mod(pResult, big.NewInt(int64(p)))
			}
		}
		//pResult = big.NewInt(int64(0))
		shares = append(shares, pResultMod)
	}
	sharesString := BigIntArrayToCharactersArray(shares)
	//put the share together with his p value
	result := PackSharesAndI(sharesString)
	return result, nil
}
func StringArrayToBigIntArray(s []string) []*big.Int {
	var r []*big.Int
	//sBytes := []byte(s)
	for _, sElem := range s {
		b, err := strconv.Atoi(sElem)
		if err != nil {
			fmt.Println(err)
		}
		bBig := big.NewInt(int64(b))
		r = append(r, bBig)
	}
	return r
}
func CharactersArrayToBigIntArray(s []string) []*big.Int {
	var r []*big.Int
	//sBytes := []byte(s)
	for _, sElem := range s {
		b := []byte(sElem)[0]
		bBig := big.NewInt(int64(b))
		r = append(r, bBig)
	}
	return r
}
func BigIntArrayToCharactersArray(b []*big.Int) []string {
	var r []string
	for _, bigint := range b {
		r = append(r, bigint.String())
	}
	return r
}
func PackSharesAndI(sharesString []string) [][]string {
	var r [][]string
	for i, share := range sharesString {
		curr := []string{share, strconv.Itoa(i + 1)}
		r = append(r, curr)
	}
	return r
}
func UnpackSharesAndI(sharesPacked [][]string) ([]string, []string) {
	var shares []string
	var i []string
	for _, share := range sharesPacked {
		shares = append(shares, share[0])
		i = append(i, share[1])
	}
	return shares, i
}

func LagrangeInterpolation(sharesGiven [][]string, p int) *big.Int {

	//result := big.NewInt(int64(0))
	resultN := big.NewInt(int64(0))
	resultD := big.NewInt(int64(0))

	//unpack shares
	shares, sharesI := UnpackSharesAndI(sharesGiven)

	sharesBigInt := StringArrayToBigIntArray(shares)
	sharesIBigInt := StringArrayToBigIntArray(sharesI)
	for i := 0; i < len(sharesBigInt); i++ {
		lagrangeNumerator := big.NewInt(int64(1))
		lagrangeDenominator := big.NewInt(int64(1))
		for j := 0; j < len(sharesBigInt); j++ {
			if sharesI[i] != sharesI[j] {
				//lagrBasis := secretShare.LagrangeBasis(new(big.Float).SetInt(sharesIBigInt[j]), new(big.Float).SetInt(sharesIBigInt[i]))
				currLagrangeNumerator := sharesIBigInt[j]
				currLagrangeDenominator := new(big.Int).Sub(sharesIBigInt[j], sharesIBigInt[i])

				lagrangeNumerator = new(big.Int).Mul(lagrangeNumerator, currLagrangeNumerator)
				lagrangeDenominator = new(big.Int).Mul(lagrangeDenominator, currLagrangeDenominator)

			}
		}
		numerator := new(big.Int).Mul(sharesBigInt[i], lagrangeNumerator)
		quo := new(big.Int).Quo(numerator, lagrangeDenominator)
		//fmt.Println(quo)
		if quo.Int64() != 0 {
			resultN = resultN.Add(resultN, quo)
		} else {
			resultNMULlagrangeDenominator := new(big.Int).Mul(resultN, lagrangeDenominator)
			resultN = new(big.Int).Add(resultNMULlagrangeDenominator, numerator)

			resultD = resultD.Add(resultD, lagrangeDenominator)
		}
	}

	fmt.Print("nominator: ")
	fmt.Print(resultN)
	fmt.Print(", denominator: ")
	fmt.Println(resultD)
	fmt.Print("result: ")
	fmt.Print(resultN)
	fmt.Print("/")
	fmt.Print(resultD)
	fmt.Print(" mod ")
	fmt.Println(p)
	var modinvMul *big.Int
	if resultD.Int64() != 0 {
		modinv := new(big.Int).ModInverse(resultD, big.NewInt(int64(p)))
		fmt.Println("modinv: ")
		fmt.Println(modinv)
		modinvMul = new(big.Int).Mul(resultN, modinv)
	} else {
		modinvMul = resultN
	}
	r := new(big.Int).Mod(modinvMul, big.NewInt(int64(p)))
	//resultMod, asdf := new(big.Int).DivMod(resultN, resultD, big.NewInt(int64(p)))
	fmt.Println(r)
	return r
}

//l is the others values != i
//j is the share i
func LagrangeBasis(l *big.Float, j *big.Float) *big.Float {
	l_j := new(big.Float).Sub(l, j)
	sublambda := new(big.Float).Quo(l, l_j)
	return sublambda
}
