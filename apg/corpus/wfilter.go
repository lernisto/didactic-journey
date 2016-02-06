package main

/*
This program reads a list of words (1 per line, e.g. /usr/share/dict/words)
and selects a given number of words.

Filters:
	acronyms (i.e. all capitals)
	posessives ending in "'s"
	words longer than a certain length (e.g. 7)

I created this to make a "diceware-like" word list
*/

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var word_re = regexp.MustCompile("['A-Za-z0-9]['a-z0-9]*")
var possessive = regexp.MustCompile(".*'s$")

func main() {

	maxlen := 7
	maxwords := 16384
	args := os.Args[1:]
	if len(args) >= 1 {
		ml, e := strconv.Atoi(args[0])
		if e == nil {
			maxlen = ml
		}
		if len(args) >= 2 {
			c, e := strconv.Atoi(args[1])
			if e == nil {
				maxwords = c
			}
		}
	}
	fmt.Fprintf(os.Stderr, "maxlen: %d\nmaxwords: %d\n", maxlen, maxwords)

	//words := make([]string, maxlen)
	count := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := strings.ToLower(s.Text())
		if len(word) > maxlen {
			continue
		}
		if word_re.MatchString(word) && !possessive.MatchString(word) {
			//words[i] = word
			fmt.Println(word)
			count++
			if count >= maxwords {
				break
			}
		}
	}

}
