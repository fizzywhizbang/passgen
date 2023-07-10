package main

import (
	"crypto/rand"
	"flag"
	"math/big"
)

// characters for password
const Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"

func main() {
	//create flag for length
	length := flag.Int("l", 20, "password length")
	flag.Parse()

	println(genPass(*length))
}

// generate password
func genPass(length int) string {
	max := len(Chars)
	pass := ""
	//loop the length/number of characters required for password
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
		pass = pass + string(Chars[n.Int64()])
	}

	return pass
}
