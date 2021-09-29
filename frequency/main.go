package main

import (
	"fmt"
	"io"
	"log"
	"sort"
	"unicode"
)

func UpperMostFrequent() []uint8 {
	return []uint8{
		'A',
		'E',
		'O',
		'S',
		'R',
	}
}

func main() {

	charsMap := readAllChars()
	mostFrequentRunes := sortChars(charsMap)

	for _, char := range mostFrequentRunes {
		log.Println(string(char))
	}
}

func readAllChars() map[rune]int {
	var read string

	charsMap := make(map[rune]int)
	for {
		_, err := fmt.Scanf("%s", &read)
		if err == io.EOF {
			break
		}

		// Busca as letras mais comuns
		for _, char := range read {
			if !unicode.IsLetter(char) {
				continue
			}

			upperChar := unicode.ToUpper(char)

			_, ok := charsMap[upperChar]
			if !ok {
				charsMap[upperChar] = 1
				continue
			}

			charsMap[upperChar]++
		}
	}

	return charsMap
}

func sortChars(charsMap map[rune]int) []rune {
	chars := make([]rune, 0, len(charsMap))
	for key := range charsMap {
		chars = append(chars, key)
	}

	sort.Slice(chars, func(i, j int) bool {
		return charsMap[chars[i]] > charsMap[chars[j]]
	})

	return chars
}
