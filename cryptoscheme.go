package paillier

import (
	"math/big"
	"math/rand"
)

const (
	defaultP int64 = 1_000_000_007
	defaultQ int64 = 41
)

type PublicKey struct {
	n *big.Int
	g *big.Int
}

type PrivateKey struct {
	n *big.Int
	l *big.Int
	u *big.Int
}

type PublicValue struct {
	Val *big.Int
}

type PrivateValue struct {
	Val *big.Int
}

type PaillierScheme interface {
	Encode(key *PublicKey, m *PrivateValue) *PublicValue

	Decode(key *PrivateKey, c *PublicValue) *PrivateValue

	GenKeypair() (*PublicKey, *PrivateKey)

	SetInitialPrimes(p int64, q int64)

	Add(a *PublicValue, b *PublicValue, key *PublicKey) *PublicValue

	Mul(a *PublicValue, b *big.Int, key *PublicKey) *PublicValue

	Sub(a *PublicValue, b *PublicValue, key *PublicKey) *PublicValue
}

type paillier struct {
	P *big.Int
	Q *big.Int
}

func (p *paillier) GenKeypair() (*PublicKey, *PrivateKey) {
	n := big.NewInt(0).Mul(p.P, p.Q)
	g := big.NewInt(rand.Int63n(p.P.Int64()))
	l := _lcm(_dec(p.P), _dec(p.Q))
	u := _rev(_l(_pow(g, l, _square(n)), n), p.P, p.Q)
	return &PublicKey{n, g}, &PrivateKey{n, l, u}
}

func (p *paillier) Encode(key *PublicKey, m *PrivateValue) *PublicValue {
	nn := _square(key.n)
	r := big.NewInt(rand.Int63n(p.P.Int64()))
	s1 := _pow(key.g, m.Val, nn)
	s2 := _pow(r, key.n, nn)
	return &PublicValue{_bigMul(s1, s2, nn)}
}

func (p *paillier) Decode(key *PrivateKey, c *PublicValue) *PrivateValue {
	nn := _square(key.n)
	h := _pow(c.Val, key.l, nn)
	hL := _l(h, key.n)
	return &PrivateValue{_bigMul(hL, key.u, key.n)}
}

func (p *paillier) SetInitialPrimes(P int64, Q int64) {
	p.P = big.NewInt(P)
	p.Q = big.NewInt(Q)
}

func GetNewInstance() PaillierScheme {
	return &paillier{big.NewInt(defaultP), big.NewInt(defaultQ)}
}

func (p *paillier) Add(a *PublicValue, b *PublicValue, key *PublicKey) *PublicValue {
	nn := _square(key.n)
	return &PublicValue{Val: _bigMul(a.Val, b.Val, nn)}
}

func (p *paillier) Mul(a *PublicValue, b *big.Int, key *PublicKey) *PublicValue {
	nn := _square(key.n)
	return &PublicValue{Val: _pow(a.Val, b, nn)}
}

func (p *paillier) Sub(a *PublicValue, b *PublicValue, key *PublicKey) *PublicValue {
	nn := _square(key.n)
	revB := _rev2(b.Val, p.P, p.Q)
	return &PublicValue{Val: _bigMul(a.Val, revB, nn)}
}
