package paillier

import "math/big"

func _equals(a *big.Int, b int64) bool {
	return a.Cmp(big.NewInt(b)) == 0
}

func _bigMul(a, b, mod *big.Int) *big.Int {
	bigAB := big.NewInt(0).Mul(a, b)
	return bigAB.Mod(bigAB, mod)
}

func _pow(a, step, mod *big.Int) *big.Int {
	a = big.NewInt(0).Mod(a, mod)

	if _equals(step, 0) {
		return big.NewInt(1)
	} else if _equals(step, 1) {
		return a
	}

	h := _pow(_bigMul(a, a, mod), big.NewInt(0).Div(step, big.NewInt(2)), mod)
	flag := big.NewInt(0).Mod(step, big.NewInt(2))

	if _equals(flag, 0) {
		return h
	} else {
		return _bigMul(h, a, mod)
	}
}

func _gcd(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	} else {
		return _gcd(b, big.NewInt(0).Mod(a, b))
	}
}

func _lcm(a, b *big.Int) *big.Int {
	g := _gcd(a, b)
	h := big.NewInt(0).Div(a, g)
	return h.Mul(h, b)
}

func _dec(a *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, big.NewInt(1))
}

func _square(a *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, a)
}

func _l(u, n *big.Int) *big.Int {
	h1 := _dec(u)
	return h1.Div(h1, n)
}

func _rev(a, p, q *big.Int) *big.Int {
	n := big.NewInt(0).Mul(p, q)
	h1 := big.NewInt(0).Sub(n, p)
	h2 := big.NewInt(0).Sub(n, q)
	h1.Mul(h1, h2)
	h1.Sub(h1, big.NewInt(1))
	return _pow(a, h1, n)
}
