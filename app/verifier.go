package main
import (
	"crypto"
	"fmt"
	"math"
)

type Verifier struct {
	var A int
	var x, z []int
	var y, e int
}

func (verifier Verifier) firstVerification(N *big.Int) bool {
	if _, verifier.y := range createArray(0, verifier.A - 1) && N > 1 {

		for i := 0; i < m; i++ {
			var value_check int = math.Pow(verifier.z[i], verifier.y - verifier.e * N) % N

			if verifier.x[i] == value_check && N % 2 != 0 && N % 3 != 0 && N % 5 != 0 {
				var verificationResult bool = true

			} else {
				var verificationResult bool = false
			}
		}	
	} else {
		var verificationResult bool = false
	}
	return verificationResult
}