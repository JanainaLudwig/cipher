package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		read += scanner.Text()  + "\n"
	}

	read = read[0: len(read)-1]

	if *cifrar {
		fmt.Printf("%v", c.Crypt(read))
	} else if *decifrar {
		fmt.Printf("%v", c.Decrypt(read))
	} else {
		log.Println("Please use -c or -d option.")
	}
}
