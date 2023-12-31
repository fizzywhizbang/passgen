package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

// allowed characters for password
const Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+,.?/:;{}[]`~"

func main() {
	//create flag for length
	length := flag.Int("l", 20, "password length")
	writeToFile := flag.Bool("f", false, "write to file")
	flag.Parse()
	pass := genPass(*length)
	if *writeToFile {
		fileName := "password.txt"
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		fmt.Println("your password is", pass, "and has been written to", fileName)
		f.WriteString(pass)

	} else {
		fmt.Println(pass)
	}
}

// generate password
func genPass(length int) string {
	max := len(Chars)
	pass := ""
	//iterate over the length/number of characters required for password
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
		pass = pass + string(Chars[n.Int64()])
	}

	return pass
}
