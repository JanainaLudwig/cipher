package main

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
		converted += c.rotateChar(uint8(char), key)
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

	if mod == 0 {
		return string(char)
	}

	if key > 0 {
		rotated = ((int(char) - min + key) % mod) + min
	} else {
		rotated = ((int(char) - max + key) % mod) + max
	}

	return string(rune(rotated))
}
