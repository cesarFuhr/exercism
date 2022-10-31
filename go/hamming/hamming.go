package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences don't have the same length")
	}

	misses := 0
	for i := range a {
		if a[i] != b[i] {
			misses++
		}
	}

	return misses, nil
}
