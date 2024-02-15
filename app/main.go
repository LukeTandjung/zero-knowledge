package main
import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	person1 := Prover{
		l: 1,
		k: 256, 
		A: int(math.Pow(2, 2048)),
		B: int(math.Pow(2, 256)), 
		m: 80
	}
	x_claim, y_claim, z_claim, e_claim := person1.firstClaim(87654345678)
	person2 := Verifier{
		A: math.Pow(2, 2048), 
		x: x_claim,
		y: y_claim, 
		z: z_claim, 
		e: e_claim
	}
	person1.firstVerification(87654345678)
}