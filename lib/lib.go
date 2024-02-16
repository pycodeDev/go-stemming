package lib

import "strings"

func StemIndonesianWord(word string) string {
	// Daftar awalan dan akhiran yang mungkin dalam bahasa Indonesia
	prefixes := []string{"peng", "pen", "ber", "per", "pe", "di", "ke", "se", "te", "me"}
	suffixes := []string{"lah", "kah", "tah", "pun", "ku", "mu", "nya", "kan", "an"}

	// Menghapus awalan yang sesuai dari kata
	for _, prefix := range prefixes {
		if strings.HasPrefix(word, prefix) {
			word = word[len(prefix):]
			break
		}
	}

	// Menghapus akhiran yang sesuai dari kata
	for _, suffix := range suffixes {
		if strings.HasSuffix(word, suffix) {
			word = word[:len(word)-len(suffix)]
			break
		}
	}

	return word
}
