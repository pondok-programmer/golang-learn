package sidik

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
		result := make(map[string]int)

		for _, word := range words{
			result[word] += 1
		}

		return result
}

func type23() {
	wc.Test(WordCount)
}
