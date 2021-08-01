package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const (
	topLen = 10
)

func Top10(text string) []string {
	var result = make([]string, 0, topLen)
	if len(text) == 0 {
		return result
	}

	dict := splitWordByFrequency(text)

	var result_words = make(map[int][]string, topLen)
	for word, cnt := range dict {
		result_words[cnt] = append(result_words[cnt], word)
	}
	var keys = make([]int, 0, topLen)
	for k, _ := range result_words {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {

		val := result_words[k]
		sort.Strings(val)
		result = append(result, val...)
	}

	if len(result) > topLen {
		return result[:topLen]
	}

	return result
}

func splitWordByFrequency(text string) map[string]int {
	var dict = make(map[string]int)
	for _, s := range strings.Fields(text) {
		dict[s]++
	}

	return dict
}
