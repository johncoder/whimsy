package whimsy

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewPhraseForZeroWordCount(t *testing.T) {
	phrase := NewPhrase(0)
	if phrase != "" {
		fmt.Println("New phrase with zero words was not empty")
		t.Fail()
	}
}

func TestNewPhraseForNonZeroWordCount(t *testing.T) {
	for i := 1; i < 10; i++ {
		genPhrase(i, t)
	}
}

func genPhrase(wordCount int, t *testing.T) {
	phrase := NewPhrase(wordCount)

	if phrase == "" {
		fmt.Printf("New phrase was empty, but expected %v words", wordCount)
		t.Fail()
	}

	words := strings.Split(phrase, " ")

	if len(words) != wordCount {
		fmt.Printf("Expected %v words, but got %v", wordCount, len(words))
		t.Fail()
	}
}
