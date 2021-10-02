package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/google/uuid"
)

// XOR
//if a, b := true, true; a != b {
//
//}

func main() {
	cifrar := flag.String("c", "", "Cifrar")
	decifrar := flag.String("d", "", "Decifrar")

	flag.Parse()

	var read string
	vernam := NewVernam()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		read = scanner.Text()
		if *cifrar != "" {
			vernam.Crypt(read)
			// todo
		} else if *decifrar != "" {
			//Decrypt(read)
		} else {
			log.Println("Please use -c or -d option with the .dat file.")
			return
		}
	}


}

type Vernam struct {
	Key string
	KeyLen int
}

func NewVernam() *Vernam {
	key := uuid.New().String()

	return &Vernam{
		Key: key,
		KeyLen: len(key),
	}
}

func (v *Vernam) Crypt(text string) {
	for i, char := range text {
		fmt.Printf("%b", uint8(char) ^ v.Key[i % v.KeyLen])
	}
}

func (v *Vernam) Decrypt(text string) {

}