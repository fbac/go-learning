package expersonal

import (
	"fmt"
	"strings"
)

// First part of the exercise:
// 	Return the word/s with more characters
// Second part of the exercise:
// 	Return the word/s with more unique characters

func main() {

	words := []string{"a", "ab", "abc", "abcd", "aaaa"}
	res := ""

	// loop over all the words
	for i := 0; i < len(words); i++ {
		// if the current word lenght is equal
		// than the word at result
		// the current word should be added as well
		// to the result
		if len(words[i]) == len(res) {
			res = fmt.Sprintf("%s %s", res, words[i])
		}

		// if current word is bigger
		// it means that it will be the new result
		if len(words[i]) > len(res) {
			res = words[i]
		}
	}
	fmt.Println("[1] Bigger word/s are: ", res)

	// unique will hold words and the number of unique chars
	// map[a:1 aaaa:1 ab:2 abc:3 abcd:4]
	unique := make(map[string]int)

	// resString will hold the unique characters for each word
	resString := ""

	// loop over all the words
	for i := 0; i < len(words); i++ {
		// loop over all the chars in every word
		for j := 0; j < len(words[i]); j++ {
			currentChar := string(words[i][j])
			if !strings.Contains(resString, currentChar) {
				// if resString doesn't contain the currentChar, add it
				resString = fmt.Sprintf("%s%s", resString, currentChar)
			}
		}
		// initialize the map for each ["word" = int]
		// where int is the len of resString (unique chars)
		// and the index is the word itself
		unique[words[i]] = len(resString)
		resString = ""
	}

	// initialize counter to count unique chars
	uniqueCounter := 0

	// initialize empty string to show the final result
	uniqueFinal := ""

	// loop over the map unique
	for i, j := range unique {
		// if the number of unique chars (j)
		// is bigger than the counter (initially 0)
		// then assign the final word to uniqueFinal and update the counter
		// when the program step into a bigger j (more unique chars)
		// uniqueFinal will be updated with the new bigger word
		if j > uniqueCounter {
			uniqueFinal = i
			uniqueCounter = j
		}
	}
	fmt.Println("[2] Most unique word: ", uniqueFinal)
}
