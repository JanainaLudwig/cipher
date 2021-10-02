package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	//scanner := bufio.NewScanner(os.Stdin)

	var t string
	for {
		_, err := fmt.Scan(&t)
		if err == io.EOF {
			break
		}

		read += t + " "
	}

	log.Println(read)

	if *cifrar != "" {
		file = *cifrar
		vernam = NewVernam(len(read))
		vernam.Crypt(read)
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

func NewVernam(textLen int) *Vernam {
	key := uuid.New().String()

	for len(key) < textLen {
		key += uuid.New().String()
	}
	//
	//// 80c82918-c66c-44c2-8eb3-cfb09d1216c156e2071e-bb93-4b37-bee6-92e871a5da4b
	key = strings.ReplaceAll(key, "-", "")
	return &Vernam{
		Key: key,
		KeyLen: len(key),
	}
}

func (v *Vernam) Crypt(text string) {
	textBinary := text
	keyBinary := v.Key

	// 80c82918-c66c-44c2-8eb3-cfb09d1216c156e2071e-bb93-4b37-bee6-92e871a5da4b

	for i, char := range textBinary {
		//log.Println("dec: ", i, v.KeyLen, i % v.KeyLen)
		fmt.Printf("%s", string(char ^ rune(keyBinary[i % v.KeyLen])))
	}
}

func (v *Vernam) Decrypt(text string) {
	keyBinary := v.Key


	for i, char := range text {
		//log.Println("dec: ", i, v.KeyLen, i % v.KeyLen)
		fmt.Printf("%s", string(char ^ rune(keyBinary[i % v.KeyLen])))
	}
}
