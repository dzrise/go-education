package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	f("EqualFold: %v\n", s.EqualFold("Denis", "dEnIs"))
	f("EqualFold: %v\n", s.EqualFold("Denis", "dEnI"))

	f("Index: %v\n", s.Index("Denis", "ni"))
	f("Index: %v\n", s.Index("Denis", "Ni"))

	f("Prefix: %v\n", s.HasPrefix("Denis", "De"))
	f("Prefix: %v\n", s.HasPrefix("Denis", "de"))
	f("Suffix: %v\n", s.HasSuffix("Denis", "is"))
	f("Prefix: %v\n", s.HasSuffix("Denis", "Is"))

	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))

	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	f("TrimFunction %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))

}
