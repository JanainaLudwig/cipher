package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
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
	var file string
	var vernam *Vernam

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		read = scanner.Text()
		if *cifrar != "" {
			file = *cifrar
			vernam = NewVernam()
			vernam.Crypt(read)
			// todo
		} else if *decifrar != "" {
			file = *decifrar

			file, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}

			keyBytes, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}

			vernam = &Vernam{
				Key:    string(keyBytes),
				KeyLen: len(string(keyBytes)),
			}

			vernam.Decrypt(read)
		} else {
			log.Println("Please use -c or -d option with the .dat file.")
			return
		}
	}

	err := ioutil.WriteFile(file, []byte(vernam.Key), 0644)
	if err != nil {
		log.Fatal(err)
		return
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
	textBinary := text
	keyBinary := v.Key


	for i, char := range textBinary {
		fmt.Printf("%s", string(uint8(char) ^ keyBinary[i % v.KeyLen]))
	}
}

func (v *Vernam) Decrypt(text string) {
	keyBinary := v.Key

	for i, char := range text {
		fmt.Printf("%s", string(uint8(char) ^ keyBinary[i % v.KeyLen]))
	}
}
