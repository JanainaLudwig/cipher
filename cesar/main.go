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


func (c *Cesar) getMinMod(char uint8) (min int, mod int) {
	if char >= 'A' && char <= 'Z' { // A - Z
		min, mod = 65, 26
	}

	if char >= 'a' && char <= 'z' { // a - z
		min, mod = 97, 26
	}

	if char >= '0' && char <= '9' { // a - z
		min, mod = 48, 10
	}
	
	return min, mod
}

func (c *Cesar) rotateChar(char uint8, key int) string {
	var rotated int

	min, mod := c.getMinMod(char)

	rotated = ((int(char) - min + key) % mod) + min

	return string(rune(rotated))
}
