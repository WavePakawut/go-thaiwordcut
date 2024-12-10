package gothaiwordcut

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/armon/go-radix"
)

// Segmenter : Segmenter main class
type Segmenter struct {
	Tree *radix.Tree

	minLength int
}

// Option : Option for Segmenter
type Option func(*Segmenter)

func (w *Segmenter) loadFileIntoTrie(filePath string) {
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		w.Tree.Insert(scanner.Text(), 1)
	}

	check(scanner.Err())
}

func (w *Segmenter) findSegment(c string) []string {
	i := 0
	N := len(c)
	arr := make([]string, 0)
	for i < N {
		// search tree
		j := w.searchTrie(c[i:N])
		if j == "" {
			i = i + 1
		} else {
			arr = append(arr, j)
			i = i + len(j)
		}
	}

	return arr
}

func (w *Segmenter) searchTrie(s string) string {
	// check if the word is latin
	latinResult := simpleRegex("[A-Za-z\\d]*", s)
	if latinResult != "" {
		return latinResult
	}

	// check if its number
	numberResult := simpleRegex("[\\d]*", s)
	if numberResult != "" {
		return numberResult
	}

	// loop word character, trying to find longest word
	longestWord, _, _ := w.Tree.LongestPrefix(s)

	return longestWord
}

func simpleRegex(expr string, s string) string {
	r, err := regexp.Compile(expr)
	check(err)
	return r.FindString(s)
}

func (w *Segmenter) Segment(txt string) []string {
	return w.findSegment(txt)
}

// Wordcut : main wordcut function
func Wordcut(options ...Option) *Segmenter {
	segmenter := &Segmenter{}
	segmenter.Tree = radix.New()
	return segmenter
}

// LoadDefaultDict : load dictionary into trie
//
//	func (w *Segmenter) LoadDefaultDict() {
//		_, filename, _, _ := runtime.Caller(0)
//		w.loadFileIntoTrie(path.Dir(filename) + "/dict/lexitron.txt")
//	}
func (w *Segmenter) LoadDefaultDict() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %v", err)
	}
	dictPath := filepath.Join(filepath.Dir(exePath), "dict", "lexitron.txt")
	if _, err := os.Stat(dictPath); os.IsNotExist(err) {
		return fmt.Errorf("dictionary file not found at %s", dictPath)
	}
	w.loadFileIntoTrie(dictPath)
	return nil
}

/*
 * If error, then we should PANIC!
 */
func check(e error) {
	if e != nil {
		panic(e)
	}
}
