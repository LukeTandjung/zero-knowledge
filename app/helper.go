package main
import (
	"fmt"
	"crypto"
	"math"
)

func createArray(start int, end int) []int {
    arr := make([]int, end-start+1)

    for i := 0; i < len(arr); i++ { 
        arr[i] = start + i 
    }
    return arr
}

func randomFromArraySecure(arr []int) int {
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(arr)))))
	
	if err != nil {
		panic("Error generating random index: " + err.Error()) // Handle error appropriately
	}
	return arr[randomIndex.Int64()]
}

func eulerTotientBig(n *big.Int) *big.Int {
    one := big.NewInt(1)
    // Directly use crypto/rsa's optimized calculation
    totient := new(big.Int).Sub(n, one) 
    totient.GCD(nil, nil, totient, n)
    return totient
}
