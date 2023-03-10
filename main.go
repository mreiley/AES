package main_test

import (
	"crypto/aes"
	"fmt"
	"reflect"
)

/*
Implement AES, the Advanced Encryption Standard. AES is one of the most popular and
secure encryption standards, and was adopted by the US government in 2002.

Test vectors are sets of known inputs and keys that generate known encrypted results (ciphers).
Inputs, keys and outputs are not necessarily text characters. NIST distributes the reference of
AES test vectors as AES Known Answer Test (KAT) Vectors.

cripto.block.Encrypt() function use parameters
src: data to cypher
dts: store data to cypher
*/
func main() {

}

/*
Inputs, keys and outputs are not necessarily text characters. Accordingly,
the test vectors are all presented,here as double-byte characters,
in other words as 16 bits per character.
*/

func VectorLinear(src []byte, expected []byte) bool {
	key := []byte{0x0, 0x1, 0x2, 0x3, 0x5, 0x6, 0x7, 0x8, 0xA, 0xB, 0xC, 0xD, 0xF, 0x10, 0x11, 0x12}
	dst := make([]byte, len(src))

	// To get a test vector to pass you also need to the algorithm to use 128 bit encryption = blocksize of 16
	blockCypher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Erro:" + err.Error())
	}

	blockCypher.Encrypt(dst, src)
	if reflect.DeepEqual(dst, expected) {
		return true
	} else {
		return false
	}

}
