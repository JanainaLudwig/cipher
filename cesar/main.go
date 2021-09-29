package main

import (
	"flag"
	"fmt"
	"io"
	"log"
)

func main() {
	cifrar := flag.Bool("c", false, "Cifrar")
	decifrar := flag.Bool("d", false, "Decifrar")
	chave := flag.Int("k", 1, "valor da chave a ser usada")

	flag.Parse()

	if *chave == 0 {
		log.Println("Please use a key different than zero.")
	}

	c := Cesar{
		Key: *chave,
	}

	var read string

	for {
		_, err := fmt.Scanf("%s", &read)
		if err == io.EOF {
			break
		}

		if *cifrar {
			fmt.Println(c.Crypt(read))
		} else if *decifrar {
			fmt.Println(c.Decrypt(read))
		} else {
			log.Println("Please use -c or -d option.")
		}
	}
}
