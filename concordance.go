package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

const SENTENCE_SPLIT_REGEX string = "([.?!;]\\s+)(?:\\p{Lu})"
const WORD_CAPTURE_REGEX string = "(\\b[^\\s]+\\b)"
const WHITESPACE_REGEX string = "\\s+"

// ===== DOMAIN OBJECTS

/**
 * A Word in the built Concordance
 */
type Word struct {
	Count      uint32 // It is possible that a word occurs multiple times in the same sentences
	Occurences map[uint32]*Sentence
	Value      string
}

/**
 * A sentence in the inputted text
 */
type Sentence struct {
	Position uint32 // The position of the sentence in relation to the rest of the text - based on a 0 based index
	Value    string
}

/**
 * The built concordance
 */
type Concordance struct {
	Words map[string]*Word
}

// ===== LOGIC

func BuildConcordance(sentence []*Sentence) *Concordance {

	concordance := new(Concordance)
	concordance.Words = make(map[string]*Word)

	for _, sentence := range sentence {
		ProcessSentence(concordance, sentence)
	}

	return concordance
}

func ProcessSentence(concordance *Concordance, sentence *Sentence) {
	words := GetSentenceWords(sentence)
	for _, word := range words {
		if _, exists := concordance.Words[word]; !exists {
			concordance.Words[word] = &Word{0, make(map[uint32]*Sentence), word}
		}
		concordance.Words[word].Count += 1
		concordance.Words[word].Occurences[sentence.Position] = sentence
	}
}

func GetSentences(str string) []*Sentence {
	var sentences []*Sentence

	regex := regexp.MustCompile(SENTENCE_SPLIT_REGEX)
	// Unfortunately GO Regex does not support look ahead / look behind so we will have to do some extra magic to make it work.
	// Identify the string splitting match, find the whitespace within that match, then grab the string from the last whitespace we
	// split on to the new whitespace we found
	whitespaceRe := regexp.MustCompile(WHITESPACE_REGEX)

	matches := regex.FindAllStringIndex(str, -1)

	sentenceStart := 0
	sentenceStop := 0
	matchesLength := len(matches)
	for i := 0; i <= matchesLength; i++ {

		if i < matchesLength {
			match := matches[i]
			whitespaceStart := match[0]
			whitespaceStop := match[1]
			// Find the position of first whitespace within the matched substring. The matched substring will look something like '. W'
			whitespaceMatchCoord := whitespaceRe.FindStringIndex(str[whitespaceStart:whitespaceStop])
			// Add the first whitespace to the index relative to the full paragraph within the
			sentenceStop = whitespaceStart + whitespaceMatchCoord[0]
		} else { // This case means we are extracting the last sentence
			sentenceStop = len(str) - 1
		}

		// Finally, cleanup leading and trailing whitespace
		cleanSentence := strings.TrimSpace(str[sentenceStart:sentenceStop])
		sentences = append(sentences, &Sentence{uint32(i), cleanSentence})

		// Finally, move the pointer within the full string to the stop postion of the sentence we just added. Also increment the sentence index
		sentenceStart = sentenceStop
	}
	return sentences
}

// Given a sentence, return an array of String words
func GetSentenceWords(sentence *Sentence) []string {
	regex := regexp.MustCompile(WORD_CAPTURE_REGEX)
	words := regex.FindAllString(sentence.Value, -1)
	return words
}

// ===== MAIN

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintf(os.Stderr, "To use this tool data must be piped to the program\n")
		// Exit with error status
		os.Exit(1)
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read piped input. Receieved this error %v\n", err)
		// Exit with error status
		os.Exit(1)
	}

	text := string(bytes)

	// Build concordance - all of this I would consider "display logic" that would belong in a lightweight controller and View
	// For simplicity I am going to do it inline here.
	sentences := GetSentences(text)
	concordance := BuildConcordance(sentences)

	if len(concordance.Words) == 0 {
		fmt.Fprintf(os.Stderr, "No words received for processing\n")
		os.Exit(1)
	}

	textLen := len(text)
	maxLen := textLen
	if textLen > 100 {
		maxLen = textLen
	}
	fmt.Fprintf(os.Stdout, "Received the following text: %s\n", text[0:maxLen])
	// Print results to stdout
	fmt.Fprintf(os.Stdout, "Found %d words in %d sentences\n", len(concordance.Words), len(sentences))

	keysSorted := make([]string, len(concordance.Words))
	i := 0
	for key, _ := range concordance.Words {
		keysSorted[i] = key
		i++
	}

	sort.Strings(keysSorted)

	for _, wordKey := range keysSorted {
		word := concordance.Words[wordKey]
		fmt.Fprintf(os.Stdout, "%s\n", word.Value)
		fmt.Fprintf(os.Stdout, "\tCount: %d\n", word.Count)
	}
}
