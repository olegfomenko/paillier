package paillier

import (
	"fmt"
	"math/big"
)

func _equals(a *big.Int, b int64) bool {
	return a.Cmp(big.NewInt(b)) == 0
}

func _bigMul(a, b, mod *big.Int) *big.Int {
	bigAB := big.NewInt(0).Mul(a, b)
	return bigAB.Mod(bigAB, mod)
}

func _pow(a, step, mod *big.Int) *big.Int {
	return big.NewInt(0).Exp(a, step, mod)
}

func _gcd(a, b *big.Int) *big.Int {
	return big.NewInt(0).GCD(nil, nil, a, b)
}

func _lcm(a, b *big.Int) *big.Int {
	g := _gcd(a, b)
	h := big.NewInt(0).Div(a, g)
	return h.Mul(h, b)
}

func _dec(a *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, big.NewInt(1))
}

func _inc(a *big.Int) *big.Int {
	return big.NewInt(0).Add(a, big.NewInt(1))
}

func _square(a *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, a)
}

func _l(u, n *big.Int) *big.Int {
	/*if !_equals(big.NewInt(0).Mod(u, n), 1) {
		panic("division error")
	}*/
	h1 := _dec(u)
	return h1.Div(h1, n)
}

func _rev(a, n *big.Int) (*big.Int, error) {
	result := big.NewInt(0).ModInverse(a, n)

	if result == nil {
		fmt.Println("Mod Inverse operation error!")
		return nil, InverseError{}
	} else {
		return result, nil
	}
}
