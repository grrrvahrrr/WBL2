package main

import (
	"fmt"
	"sort"
	"strings"
)

//dictionary and words
var (
	words      = []string{"меч", "пятка", "столик", "голова", "еда"}
	dictionary = []string{"голова", "еда", "листок", "меч", "пятак", "пятка", "слиток", "столик", "тяпка"}
)

func main() {
	//Creating map of anagramms
	anagramms := make(map[string][]string)

	//Ranging through the slice of words
	for _, word := range words {
		//bringing word to lower case
		word := strings.ToLower(word)
		//calculating total weight of the word
		var vTotal int
		var wordSlice []string
		for _, v := range word {
			vTotal += int(v)
		}

		//Going through dictionary to search through words with same weight
		for _, wordDict := range dictionary {
			var vTotalDict int
			for _, v := range wordDict {
				vTotalDict += int(v)
			}
			if vTotal == vTotalDict {
				wordSlice = append(wordSlice, wordDict)
			}
		}
		//If there is more that one word of this weight append it write the slice to map
		if len(wordSlice) > 1 {
			sort.Strings(wordSlice)
			anagramms[wordSlice[0]] = wordSlice
		}
	}

	fmt.Println(anagramms)
}
