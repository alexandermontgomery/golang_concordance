package main

import (
	"testing"
)

func TestGetWords(t *testing.T) {
	sentence := &Sentence{0, "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown..."}
	words := GetSentenceWords(sentence)
	actual := 16
	if len(words) != actual {
		t.Errorf("Failed to parse %v words from sentence. Result: %v", actual, words)
	}

	sentence.Value = "The famous Merriam-Webster dictionary was written by two individuals."
	words = GetSentenceWords(sentence)
	actual = 9
	if len(words) != actual {
		t.Errorf("Failed to parse %v words from sentence. Result: %v", actual, words)
	}

}

func TestGetSentences(t *testing.T) {

	// Test a rather normal path
	str := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.

	It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).`

	sentences := GetSentences(str)

	actual := 8
	if len(sentences) != actual {
		t.Errorf("Expected to find %v sentences but actually found %v. Result: %v", actual, len(sentences), sentences)
	}

	// Test a more exceptional example - note here that a human sees 4 sentences, but the sentence regex sees 5
	str = `Does the largest city in a state always have the highest population? In some states, the city with the highest population is the capital city. In other states, large cities are not the capital city, e.g. New York City.  This makes guessing capital cities challenging.`
	sentences = GetSentences(str)
	actual = 5
	if len(sentences) != actual {
		t.Errorf("Expected to find %v sentences but actually found %v. Result: %v", actual, len(sentences), sentences)
	}

	// Run the same test except this time lower case new york city so the sentence regex sees 4
	str = `Does the largest city in a state always have the highest population? In some states, the city with the highest population is the capital city. In other states, large cities are not the capital city, e.g. new York City.  This makes guessing capital cities challenging.`
	sentences = GetSentences(str)
	actual = 4
	if len(sentences) != actual {
		t.Errorf("Expected to find %v sentences but actually found %v. Result: %v", actual, len(sentences), sentences)
	}
}

func TestBuildConcordance(t *testing.T) {
	sentences := []*Sentence{
		&Sentence{0, "There are many programming languages in the world today"},
		&Sentence{1, "Nobody can really say that one is better than the others"},
		&Sentence{2, "Golang is certainly one of the languages I enjoy - but there are many others"},
		&Sentence{3, "Others include PHP, Java, C++, and Python"},
		&Sentence{4, "Golang's concurrency framework is pretty powerful"},
	}
	concordance := BuildConcordance(sentences)

	// Test the counts add up
	if concordance.Words["languages"].Count != 2 {
		t.Fail()
	}
	// It is a known exception that a 'strange' word like C++ will not work due to the regex
	if _, exists := concordance.Words["C++"]; exists {
		t.Fail()
	}

	if concordance.Words["Golang's"].Count != 1 {
		t.Fail()
	}

	// Test that other special char
	if concordance.Words["enjoy"].Count != 1 {
		t.Fail()
	}

}
