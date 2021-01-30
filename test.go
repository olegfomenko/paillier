package paillier

import (
	"fmt"
	"math/big"
)

func Test() {
	fmt.Println("\n--------Paillier test--------")

	scheme := GetNewInstance()
	public, private := scheme.GenKeypair()
	fmt.Println("Private / Public keys:", public, private)

	s1 := scheme.Encode(public, &PrivateValue{Val: big.NewInt(102)})
	fmt.Println("Encrypted data 102:", s1)
	fmt.Println("Decrypted data 102:", scheme.Decode(private, s1).Val.Int64())

	s2 := scheme.Encode(public, &PrivateValue{Val: big.NewInt(200)})
	fmt.Println("Encrypted data 102:", s2)
	fmt.Println("Decrypted data 102:", scheme.Decode(private, s2).Val.Int64())

	s := scheme.Add(s1, s2, public)
	fmt.Println("Decrypted sum:", scheme.Decode(private, s).Val.Int64())

	ss := scheme.Mul(s1, big.NewInt(100), public)
	fmt.Println("Mul on unencrypted val", scheme.Decode(private, ss).Val.Int64())

	sss := scheme.Sub(s2, s1, public)
	fmt.Println("Decrypted sub:", scheme.Decode(private, sss).Val.Int64())

	fmt.Println("--------Paillier test end--------\n\n\n")
}
