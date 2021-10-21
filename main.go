package main

import (
	"fmt"

	_rc4 "github.com/BEN00262/rc4-spritz/RC4"
)

func main() {
	enc := _rc4.NewRC4("john kerama")
	data := enc.Encrypt([]byte("juma is the best in the game"))

	fmt.Println("After encryption")
	fmt.Println("-----------------")
	fmt.Println(data)

	fmt.Println("After decryption")
	fmt.Println("-----------------")
	fmt.Println(
		string(enc.Encrypt(data)),
	)
}
