package main

/*
A word counter.  Splits a corpus into words and counts occurrances of unique words.

*/

// TODO: figure out the sorting better

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type pair struct {
	Word  string
	Count int
}

type by_freq []pair

func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].Count > a[j].Count }

type by_word []pair

func (a by_word) Len() int           { return len(a) }
func (a by_word) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_word) Less(i, j int) bool { return a[i].Word < a[j].Word }

type by_len []pair

func (a by_len) Len() int      { return len(a) }
func (a by_len) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a by_len) Less(i, j int) bool {
	l := len(a[i].Word) - len(a[j].Word)
	if l == 0 {
		return (a[i].Count > a[j].Count)
	} else {
		return l > 0
	}
}

var word_re = regexp.MustCompile("\\b[[:alnum:]']*\\b")

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	loc := word_re.FindIndex(data)
	if loc != nil {
		if loc[1] == len(data) && !atEOF {
			return 0, nil, nil
		}
		return loc[1], data[loc[0]:loc[1]], nil
	}
	return len(data), nil, nil
}

func main() {

	count := make(map[string]int)
	s := bufio.NewScanner(os.Stdin)
	s.Split(split) //bufio.ScanWords)
	for s.Scan() {
		count[strings.ToLower(s.Text())]++
	}

	words := make([]pair, len(count))
	pw := new(pair)
	i := 0
	for k, v := range count {
		pw.Word = k
		pw.Count = v
		words[i] = *pw
		i++
	}

	sort.Sort(by_len(words))
	for _, p := range words {
		fmt.Printf("%s\t%d\n", p.Word, p.Count)
	}
}
