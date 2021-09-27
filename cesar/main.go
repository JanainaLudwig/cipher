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

	log.Println("cifrar:", *cifrar)
	log.Println("decifrar:", *decifrar)
	log.Println("chave:", *chave)
	log.Println("input:", *input)
	log.Println("output:", *output)

	c := Cesar{
		Key: *chave,
	}

	read, err := Read(*input)

	err = Write(*output, c.Crypt(read))
	if err != nil {
		log.Fatal(err)
	}
}
