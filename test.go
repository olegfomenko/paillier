package paillier

import (
	"fmt"
	"math/big"
	"math/rand"
)

func Test() {
	fmt.Println("\n--------Paillier test--------")

	scheme := GetNewInstance()
	public, private := scheme.GenKeypair()
	fmt.Println("Private / Public keys:", public, private)

	s1 := scheme.Encrypt(public, &PrivateValue{Val: big.NewInt(102)})
	fmt.Println("Encrypted data 102:", s1)
	fmt.Println("Decrypted data 102:", scheme.Decrypt(private, s1).Val.Int64())

	s2 := scheme.Encrypt(public, &PrivateValue{Val: big.NewInt(200)})
	fmt.Println("Encrypted data 102:", s2)
	fmt.Println("Decrypted data 102:", scheme.Decrypt(private, s2).Val.Int64())

	s := scheme.Add(s1, s2, public)
	fmt.Println("Decrypted sum:", scheme.Decrypt(private, s).Val.Int64())

	ss := scheme.Mul(s1, big.NewInt(100), public)
	fmt.Println("Mul on unencrypted val", scheme.Decrypt(private, ss).Val.Int64())

	sss := scheme.Sub(s2, s1, public)
	fmt.Println("Decrypted sub:", scheme.Decrypt(private, sss).Val.Int64())

	var cnt = 0

	for i := 0; i < 1000; i++ {
		val := rand.Int63n(scheme.GetP())
		enc := scheme.SafeEncrypt(public, private, &PrivateValue{Val: big.NewInt(val)})
		dec := scheme.Decrypt(private, enc).Val.Int64()

		if dec != val {
			cnt++
			fmt.Println(val, dec, enc)
		}
	}

	fmt.Println(cnt)

	fmt.Println("--------Paillier test end--------\n\n\n")
}
