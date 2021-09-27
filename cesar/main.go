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

	c.rotateChar('a')
	c.rotateChar('A')
	c.rotateChar('0')
	c.rotateChar('z')
	c.rotateChar('Z')
	c.rotateChar('9')
}


type Cesar struct {
	Key int
}

func (c *Cesar) rotateChar(char uint8) string {
	var rotated int

	if c.Key > 0 {
		switch char {
		case 122: // z
			rotated = 96 + c.Key // a - 1
		case 90: // Z
			rotated = 64 + c.Key // A - 1
		case 57: // 9
			rotated = 47 + c.Key // 0 - 1
		default:
			rotated = int(char) + c.Key
		}
	} else {
		switch char {
		case 97: // - a
			rotated = 123 + c.Key // a -
		case 65: // - A-
			rotated = 91  + c.Key // A -
		case 48: // - 0-
			rotated = 58  + c.Key // 0 -
		default:
			rotated = int(char) + c.Key
		}
	}

	//log.Println(string(char), char, "=>", rotated, rune(rotated), string(rune(rotated)))
	return string(rune(rotated))
}
