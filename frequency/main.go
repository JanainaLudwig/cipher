package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"unicode"
)

func UpperMostFrequent() []rune {
	return []rune{
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

	key0 := getKey(mostFrequentRunes[0], UpperMostFrequent()[0])
	if key0 == getKey(mostFrequentRunes[1], UpperMostFrequent()[1]) {
		log.Println("\nKey:", key0)
	}

	correctPercentage(mostFrequentRunes)
}

func correctPercentage(sortedChars []rune) {
	//shuffleFrequent := UpperMostFrequent()
	//for i := range shuffleFrequent {
	//	j := rand.Intn(i + 1)
	//	shuffleFrequent[i], shuffleFrequent[j] = shuffleFrequent[j], shuffleFrequent[i]
	//}

	//qtdFrequentWords := len(shuffleFrequent)


	commonKeys := make(map[int]int)

	for i, frequentWord := range UpperMostFrequent() { // Passa pelas letras mais comuns
		//for _, char := range sortedChars[0:10] {

		key := getKey(sortedChars[i], frequentWord) // Compara cada letra
		log.Println("comparing ", string(sortedChars[i]), string(frequentWord), "key: ", key)
		commonKeys[key]++ // soma quantidade de vezes que a chave apareceu
		//}

	}

	log.Println(commonKeys)


	//for j := 1; j < total; j++ {
	//	if n[0] < n[j] {
	//		n[0] = n[j]
	//	}
	//
	//}
}

func readAllChars() map[rune]int {
	var read string

	charsMap := make(map[rune]int)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		read = scanner.Text()

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

func getKey(from, to rune) int {
	fromNormalized := from - 'A'
	toNormalized := to - 'A'

	return int(fromNormalized - toNormalized)
}

func getReversedKey(from, to rune) int {
	fromNormalized := from - 'A'
	toNormalized := to - 'A'

	return int(toNormalized - fromNormalized)
}