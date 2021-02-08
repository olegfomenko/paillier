---
    title: Paillier cryptosystem implementation
    author: olegfomenko 
    date: 08 Fabruary 2021 
---

# Paillier cryptosystem implementation lib
This is an implementation of additive homomorphic Paillier cryptosystem this means that, given only the public key and the encryption of m1 and m2, one can compute the encryption of m1 + m2.
Also, we can get multiplication of encrypted value m1 on unencrypted m2. More information: https://en.wikipedia.org/wiki/Paillier_cryptosystem 

# Example of usage

###Import:
```go
import "github.com/olegfomenko/paillier"
```

###Crating keys:
```go
scheme := paillier.GetInstance(rand.Reader, 128)
privateKey := scheme.GenKeypair()
publicKey := privateKey.PublicKey
```

###Encryption/Decryption:
```go
encVal := scheme.Encrypt(publicKey, &PrivateValue{Val: big.NewInt(10000)})
decVal := scheme.Decrypt(privateKey, encVal).Val // equals to big int 10000
```

###Operations:
```go
var s1, s1 *paillier.PublicValue // Encrypted values

// Addition of two encrypted values
s := scheme.Add(s1, s2, publicKey)

// Subtraction of two encrypted values
s := scheme.Sub(s1, s2, publicKey)

// Multiplication on unencrypted value
ss := scheme.Mul(s1, big.NewInt(100), publicKey)
```

