package parallelletterfrequency

import (
	"strings"
	"unicode"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMapChan map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func FrequencyChan(text string) FreqMapChan {
	freqMap := make(FreqMapChan)
	text = strings.ToLower(text)
	for _, ch := range text {
		if unicode.IsLetter(ch) {
			freqMap[ch]++
		}
	}
	return freqMap

}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequencyChan(texts []string) FreqMapChan {
	total := make(FreqMapChan)
	results := make(chan FreqMapChan)
	var wg sync.WaitGroup

	for _, text := range texts {
		wg.Add(1)

		go func(){
			defer wg.Done()
			local := FrequencyChan(text)
			results<- local

		}()
	}

	go func(){
		wg.Wait()
		close(results)
	}()

	for result := range results {
		for key, value := range result {
			total[key] += value
		}
	}
	return total

}
