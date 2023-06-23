package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var regSplit = regexp.MustCompile(`[\s.,!?"':;]+`)

func Top10(text string) []string {
	wordsCount := make(map[string]int)
	words := regSplit.Split(text, -1)
	if len(words) == 0 {
		return nil
	}
	for _, word := range words {
		word = strings.ReplaceAll(strings.ToLower(word), "-", "")
		if len(word) > 0 {
			wordsCount[word] += 1
		}
	}
	keys := make([]string, 0, len(wordsCount))
	for key := range wordsCount {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	sort.SliceStable(keys, func(i, j int) bool {
		return wordsCount[keys[i]] > wordsCount[keys[j]]
	})
	if len(keys) > 9 {
		return keys[:10]
	} else {
		return keys
	}

}
