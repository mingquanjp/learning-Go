package parallelletterfrequency

import (
	"strings"
	"unicode"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

type Fmap struct {
	fmap FreqMap
	mu sync.Mutex	
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	freqMap := make(FreqMap)
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
func ConcurrentFrequency(texts []string) FreqMap {
	shareMap := Fmap{
		fmap : make(FreqMap),
	}
	var wg sync.WaitGroup
	for _,text := range texts {
		wg.Add(1)
		go func(){
			defer wg.Done()
			local := Frequency(text)
			
			shareMap.mu.Lock()

			for ch, count := range local {
				shareMap.fmap[ch] += count
			}
			shareMap.mu.Unlock()
		}()
	}
	wg.Wait()
	return shareMap.fmap
	
}
