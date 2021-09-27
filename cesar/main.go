package main

import (
	"flag"
	"log"
	"unicode"
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

	text := "texto de teste 1"

	crypt := c.Crypt(text)
	log.Println(crypt)
	log.Println(c.Decrypt(crypt))

}


type Cesar struct {
	Key int
}

func (c *Cesar) Crypt(text string) string {
	return c.convert(text, c.Key)
}

func (c *Cesar) Decrypt(text string) string {
	return c.convert(text, c.Key * -1)
}

func (c *Cesar) convert(text string, key int) string {
	var converted string
	for _, char := range text {
		if unicode.IsNumber(char) || unicode.IsLetter(char) {
			converted += c.rotateChar(uint8(char), key)
		} else {
			converted += string(char)
		}
	}

	return converted
}


func (c *Cesar) rotateChar(char uint8, key int) string {
	var rotated int

	if key > 0 {
		switch char {
		case 122: // z


			rotated = 96 + key // a - 1
		case 90: // Z
			rotated = 64 + key // A - 1
		case 57: // 9
			rotated = 47 + key // 0 - 1
		default:
			rotated = int(char) + key
		}
	} else {
		switch char {
		case 97: // - a
			rotated = 123 + key // a -
		case 65: // - A-
			rotated = 91  + key // A -
		case 48: // - 0-
			rotated = 58  + key // 0 -
		default:
			rotated = int(char) + key
		}
	}

	log.Println(string(char), char, "=>", rotated, rune(rotated), string(rune(rotated)))
	return string(rune(rotated))
}
