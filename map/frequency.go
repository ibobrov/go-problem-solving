package _map

import "sort"

func Frequencies(words []string) map[string]int {
	result := make(map[string]int)

	for _, word := range words {
		result[word]++
	}

	return result
}

func MostFrequent(freq map[string]int) (word string, count int, ok bool) {
	if len(freq) == 0 {
		return "", 0, false
	}

	type WordCount struct {
		Word  string
		Count int
	}

	words := make([]WordCount, 0, len(freq))
	for k, v := range freq {
		words = append(words, WordCount{k, v})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].Count != words[j].Count {
			return words[i].Count > words[j].Count
		}
		return words[i].Word < words[j].Word
	})

	return words[0].Word, words[0].Count, true
}
