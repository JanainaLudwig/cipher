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

	text := "ABCxyz"

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


func (c *Cesar) getMinMaxMod(char uint8) (min int, max int, mod int) {
	if char >= 'A' && char <= 'Z' { // A - Z
		min, max, mod = 'A', 'Z', 'Z' - 'A' + 1
	}

	if char >= 'a' && char <= 'z' { // a - z
		min, max, mod = 'a', 'z', 'z' - 'a' + 1
	}

	if char >= '0' && char <= '9' { // a - z
		min, max, mod = '0', '9', '9' - '0' + 1
	}
	
	return min, max, mod
}

func (c *Cesar) rotateChar(char uint8, key int) string {
	var rotated int

	min, max, mod := c.getMinMaxMod(char)

	if key > 0 {
		rotated = ((int(char) - min + key) % mod) + min
	} else {
		rotated = ((int(char) - max + key) % mod) + max
	}

	return string(rune(rotated))
}
