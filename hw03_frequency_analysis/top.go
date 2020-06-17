package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
	"unicode"
)

// Top10 returns top 10 words in the text
//
func Top10(text string) []string {
	// Using map [string]int to count lowercased words in the text
	wordStats := map[string]int{}
	words := strings.FieldsFunc(text, func(c rune) bool {
		return unicode.IsSpace(c) || c == '–' || c == '—'
	})
	for _, word := range words {
		// Trimming
		_word := strings.Trim(word, "?.,!-")
		wordStats[strings.ToLower(_word)]++
	}
	delete(wordStats, "")

	// Using slice of string and sort by map count value
	wordStatsSlice := make([]string, 0, len(wordStats))
	for word := range wordStats {
		wordStatsSlice = append(wordStatsSlice, word)
	}
	sort.Slice(wordStatsSlice, func(i, j int) bool {
		return wordStats[wordStatsSlice[i]] > wordStats[wordStatsSlice[j]]
	})

	// return slice as is if length less or equal 10
	if len(wordStatsSlice) <= 10 {
		return wordStatsSlice
	}
	return wordStatsSlice[:10]
}
