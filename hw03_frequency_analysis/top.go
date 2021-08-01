package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const (
	topResultLen = 10
	topLen       = 30
)

func Top10(text string) []string {
	var result = make([]string, 0, topResultLen)
	if len(text) == 0 {
		return result
	}

	wordsMap := splitWordByFrequency(text)
	keys := getSortedKeys(wordsMap)

	for _, key := range keys {
		words := wordsMap[key]
		sort.Strings(words)
		result = append(result, words...)
	}

	if len(result) > topResultLen {
		return result[:topResultLen]
	}

	return result
}

func getSortedKeys(wordsMap map[int][]string) []int {
	var keys = make([]int, 0, topLen)
	for k, _ := range wordsMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	return keys
}

func splitWordByFrequency(text string) map[int][]string {
	var dict = make(map[string]int)
	for _, s := range strings.Fields(text) {
		dict[s]++
	}

	var result = make(map[int][]string, topLen)
	for word, cnt := range dict {
		result[cnt] = append(result[cnt], word)
	}

	return result
}
