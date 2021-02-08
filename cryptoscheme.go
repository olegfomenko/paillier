package paillier

import (
	"crypto/rand"
	"math/big"
)

const defaultKeySize int = 128

type PublicKey struct {
	n  *big.Int
	nn *big.Int
	g  *big.Int
}

type PrivateKey struct {
	PublicKey *PublicKey
	h         *big.Int
	u         *big.Int
}

type PublicValue struct {
	Val *big.Int
}

type PrivateValue struct {
	Val *big.Int
}

type InverseError struct {
}

func (e InverseError) Error() string {
	return "Mod inverse error!"
}

type PaillierScheme interface {
	Encrypt(key *PublicKey, m *PrivateValue) *PublicValue

	Decrypt(key *PrivateKey, c *PublicValue) *PrivateValue

	GenKeypair() *PrivateKey

	Add(a *PublicValue, b *PublicValue, key *PublicKey) *PublicValue

	Mul(a *PublicValue, b *big.Int, key *PublicKey) *PublicValue

	Sub(a *PublicValue, b *PublicValue, key *PublicKey) *PublicValue
}

type paillier struct {
	P *big.Int
	Q *big.Int
}

func (p *paillier) GenKeypair() *PrivateKey {
	n := new(big.Int).Mul(p.P, p.Q)
	nn := square(n)
	g := inc(n)

	h := lcm(dec(p.P), dec(p.Q))
	u, err := rev(l(pow(g, h, nn), n), n)

	if err != nil {
		panic(err)
	}

	return &PrivateKey{
		PublicKey: &PublicKey{
			n:  n,
			nn: nn,
			g:  g,
		},
		h: h,
		u: u,
	}
}

func (p *paillier) Encrypt(key *PublicKey, m *PrivateValue) *PublicValue {
	r, err := rand.Int(rand.Reader, key.n)
	if err != nil {
		panic(err)
	}

	s1 := pow(key.g, m.Val, key.nn)
	s2 := pow(r, key.n, key.nn)
	return &PublicValue{bigMul(s1, s2, key.nn)}
}

func (p *paillier) Decrypt(key *PrivateKey, c *PublicValue) *PrivateValue {
	ch := pow(c.Val, key.h, key.PublicKey.nn)
	lVal := l(ch, key.PublicKey.n)
	return &PrivateValue{bigMul(lVal, key.u, key.PublicKey.n)}
}

func (p *paillier) Add(a *PublicValue, b *PublicValue, key *PublicKey) *PublicValue {
	return &PublicValue{Val: bigMul(a.Val, b.Val, key.nn)}
}

func (p *paillier) Mul(a *PublicValue, b *big.Int, key *PublicKey) *PublicValue {
	return &PublicValue{Val: pow(a.Val, b, key.nn)}
}

func (p *paillier) Sub(a *PublicValue, b *PublicValue, key *PublicKey) *PublicValue {
	revB, err := rev(b.Val, key.nn)
	if err != nil {
		panic(err)
	}

	return &PublicValue{Val: bigMul(a.Val, revB, key.nn)}
}

func GetNewInstance() PaillierScheme {
	var instance = paillier{}

	p, err := rand.Prime(rand.Reader, defaultKeySize)
	if err != nil {
		panic(err)
	}

	q, err := rand.Prime(rand.Reader, defaultKeySize)
	if err != nil {
		panic(err)
	}

	instance.P, instance.Q = p, q
	return &instance
}
