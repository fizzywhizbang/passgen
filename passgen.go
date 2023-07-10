package main

import (
	"crypto/rand"
	"flag"
	"math/big"
)

const Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"

func main() {
	length := flag.Int("l", 20, "password length")
	flag.Parse()

	println(genPass(*length))
}

func genPass(length int) string {
	max := len(Chars)
	pass := ""

	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
		pass = pass + string(Chars[n.Int64()])
	}

	return pass
}
