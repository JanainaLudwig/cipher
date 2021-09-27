package main

import (
	"flag"
	"log"
)

func main() {
	cifrar := flag.Bool("c", false, "Cifrar")
	decifrar := flag.Bool("d", false, "Decifrar")
	chave := flag.Int("k", 1, "valor da chave a ser usada")
	input := flag.String("i", "texto-aberto.txt", "arquivo de input")
	output := flag.String("o", "texto-cifrado.txt", "arquivo de output")

	flag.Parse()

	if *chave == 0 {
		log.Println("Please use a key different than zero.")
	}

	c := Cesar{
		Key: *chave,
	}

	read, err := Read(*input)

	if *cifrar {
		err = Write(*output, c.Crypt(read))
		if err != nil {
			log.Fatal(err)
		}
	} else if *decifrar {
		err = Write(*output, c.Decrypt(read))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Please use -c or -d option.")
	}
}
