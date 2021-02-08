package paillier

import (
	"fmt"
	"math/big"
	"math/rand"
)

func Test() {
	fmt.Println("\n--------Paillier test--------")

	scheme := GetNewInstance()
	privateKey := scheme.GenKeypair()

	s1 := scheme.Encrypt(privateKey.PublicKey, &PrivateValue{Val: big.NewInt(102)})
	fmt.Println("Encrypted data 102:", s1)
	fmt.Println("Decrypted data 102:", scheme.Decrypt(privateKey, s1).Val.Int64())

	s2 := scheme.Encrypt(privateKey.PublicKey, &PrivateValue{Val: big.NewInt(200)})
	fmt.Println("Encrypted data 102:", s2)
	fmt.Println("Decrypted data 102:", scheme.Decrypt(privateKey, s2).Val.Int64())

	s := scheme.Add(s1, s2, privateKey.PublicKey)
	fmt.Println("Decrypted sum:", scheme.Decrypt(privateKey, s).Val.Int64())

	ss := scheme.Mul(s1, big.NewInt(100), privateKey.PublicKey)
	fmt.Println("Mul on unencrypted val", scheme.Decrypt(privateKey, ss).Val.Int64())

	sss := scheme.Sub(s2, s1, privateKey.PublicKey)
	fmt.Println("Decrypted sub:", scheme.Decrypt(privateKey, sss).Val.Int64())

	var cnt = 0

	for i := 0; i < 1000; i++ {
		val := rand.Int63()
		enc := scheme.Encrypt(privateKey.PublicKey, &PrivateValue{Val: big.NewInt(val)})
		dec := scheme.Decrypt(privateKey, enc).Val.Int64()

		if dec != val {
			cnt++
			fmt.Println(val, dec, enc)
		}
	}

	fmt.Println(cnt)

	fmt.Println("--------Paillier test end--------\n\n\n")
}
