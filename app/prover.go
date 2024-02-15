package main
import (
	"crypto"
	"fmt"
	"math"
)

type Prover struct {
	var l, k, A, B, m int
}

func (prover Prover) firstClaim(N *big.Int) ([]int, int []int, int) {

	var r int = randomFromArraySecure(createArray(0, prover.A-1))
	var z [prover.m]int
	var x [prover.m]int

	for i := 0; i < prover.m; i++ {
		z[i] := randomFromArraySecure(createArray(1, N-1))
		x[i] := math.Pow(z[i], prover.r) % N
	}
	
	var e int = randomFromArraySecure(createArray(0, prover.B-1))
	var y int = r + e * (N - eulerTotientBig(N))

	return x, y, z, e
}
