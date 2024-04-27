package patternsearch

import "math"

const RABIN_KARP_BASE = 256
const RABIN_KARP_MODULUS = math.MaxInt

// Rabin-Karp Algorithm uses a rolling hash to quickly filter out positions of the text that cannot match the pattern,
// and then checks for a match at the remaining positions.
// https://www.geeksforgeeks.org/rabin-karp-algorithm-for-pattern-searching/
func RabinKarpSearch(txt string, pat string) []int {
	var p int = 0 // hash of the pattern
	var t int = 0 // hash of the window of txt
	var h int = 1
	// The value of h would be "pow(d, M-1)%q"
	for i := 0; i < len(pat)-1; i++ {
		h = (RABIN_KARP_BASE * h) % RABIN_KARP_MODULUS
	}

	// Calculate the hash value of pattern and first window of text
	for i := 0; i < len(pat); i++ {
		p = (RABIN_KARP_BASE*p + int(pat[i])) % RABIN_KARP_MODULUS
		t = (RABIN_KARP_BASE*t + int(txt[i])) % RABIN_KARP_MODULUS
	}

	var ret []int

	for i := 0; i <= len(txt)-len(pat); i++ {
		// Check the hash values of current window of text and pattern.
		// If the hash values match then only check for characters one by one
		if p == t {
			var isFound bool = true
			for j := 0; j < len(pat); j++ {
				if pat[j] != txt[i+j] {
					isFound = false
					break
				}
			}
			if isFound {
				ret = append(ret, i)
			}
		}

		// Calculate hash value for next window of text:
		// Remove leading digit, add trailing digit
		if i < len(txt)-len(pat) {
			t = (RABIN_KARP_BASE*(t-int(txt[i])*h) + int(txt[i+len(pat)])) % RABIN_KARP_MODULUS
			// We might get negative value of t, converting it to positive
			if t < 0 {
				t = t + RABIN_KARP_MODULUS
			}
		}
	}

	return ret
}
