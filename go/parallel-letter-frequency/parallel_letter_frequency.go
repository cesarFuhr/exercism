package letter

import (
	"runtime"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	var wg sync.WaitGroup
	freqMaps := make([]FreqMap, len(l))

	wg.Add(len(l))
	for i := range l {
		go func(i int) {
			defer wg.Done()

			freqMaps[i] = Frequency(l[i])
		}(i)
	}
	wg.Wait()

	final := make(FreqMap)
	for _, fm := range freqMaps {
		for key, freq := range fm {
			final[key] += freq
		}
	}

	return final
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequencyChannel(l []string) FreqMap {
	ch := make(chan FreqMap, runtime.NumCPU())

	for i := range l {
		go func(i int) {
			ch <- Frequency(l[i])
		}(i)
	}

	final := make(FreqMap, 50)
	for range l {
		for key, freq := range <-ch {
			final[key] += freq
		}
	}

	return final
}
