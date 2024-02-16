package gostemming

import (
	"errors"
	"io/ioutil"

	"github.com/pycodeDev/go-stemming/lib"
)

func OneWordStem(param string) string {
	// Kata-kata yang akan distem
	stemWord := lib.StemIndonesianWord(param)

	return stemWord
}

func MultipleWordStem(param []string) []string {
	words := []string{}

	// Melakukan stemming pada setiap kata
	for _, word := range param {
		stemmedWord := lib.StemIndonesianWord(word)
		words = append(words, stemmedWord)
	}

	return words
}

func FileWordStem(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("Error reading file: " + err.Error())
	}

	text := string(content)

	stemmedText := lib.StemIndonesianWord(text)

	outputFilePath := "stem-output.txt"
	err = ioutil.WriteFile(outputFilePath, []byte(stemmedText), 0644)
	if err != nil {
		return errors.New("Error writing to file: " + err.Error())
	}

	return nil
}
