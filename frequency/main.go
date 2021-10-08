package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
	"unicode"
)

func UpperMostFrequent() []rune {
	return []rune{
		'A',
		'E',
		'O',
		'S',
		//'R',
		//'I',
		//'N',
		//'D',
		//'M',
		//'U',
	}
}

type Result struct {
	Key int
	Percentage int
}

var minPercentage = 100

func main() {
	text := readText()
	charsMap := readAllCharsNormalized(text)
	mostFrequentRunes := sortChars(charsMap)

	result := make(chan Result)


	ticker := time.NewTicker(20 * time.Millisecond)

	for  {
		select {
		case <- ticker.C:
			minPercentage--
		case r := <- result:
			log.Printf("Chave encontrada: %v", r.Key)

			c := Cesar{Key: r.Key}

			fmt.Println(c.Decrypt(text))
			fmt.Printf("\nChave utilizada: %v", r.Key)
			return
		default:
			go findKeyPercentage(result, UpperMostFrequent(), mostFrequentRunes) // Probably this is correct one
			for i := 0; i < 10; i++ {	// But also try other ones
				go findKeyPercentage(result, shuffle(UpperMostFrequent()), mostFrequentRunes)
			}
		}
	}

}

func findKeyPercentage(result chan<-Result, commomWordsInPt, sortedChars []rune) {
	key, percentage := findKeyAndPercentage(commomWordsInPt, sortedChars)

	if percentage >= minPercentage {
		result <- Result{
			Key: key,
			Percentage: percentage,
		}
	}
}

func shuffle(shuffleFrequent []rune) []rune {
	for i := range shuffleFrequent {
		j := rand.Intn(i + 1)
		shuffleFrequent[i], shuffleFrequent[j] = shuffleFrequent[j], shuffleFrequent[i]
	}

	return shuffleFrequent
}

func findKeyAndPercentage(commomWordsInPt, sortedChars []rune) (key int, percentage int) {
	qtdFrequentWords := len(commomWordsInPt)

	commonKeys := make(map[int]int)

	for i, frequentWord := range commomWordsInPt { // Passa pelas letras mais comuns
		key := getKey(sortedChars[i], frequentWord) // Compara cada letra
		//log.Println("comparing ", string(sortedChars[i]), string(frequentWord), "key: ", key)
		commonKeys[key]++ // soma quantidade de vezes que a chave apareceu
	}

	mostFrequentKey := commonKeys[0]
	for key, count := range commonKeys {
		if count > mostFrequentKey {
			mostFrequentKey = key
		}
	}

	return mostFrequentKey, commonKeys[mostFrequentKey] * 100 / qtdFrequentWords
}

func readText() string {
	var read string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		read += scanner.Text()
	}

	return read
}

func readAllCharsNormalized(read string) map[rune]int {
	charsMap := make(map[rune]int)

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