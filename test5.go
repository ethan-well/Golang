package main

import (
	"fmt"
	"strings"
)

func main() {
	name1 := "we"
	name2 := "win"
	name := name1 + name2
	fmt.Printf("name1: %s, name2: %s, name: %s\n", name1, name2, name)

	phrase := "val og tt 魏巍"
	fmt.Printf("String: \"%s\"\n", phrase);
	for index, char := range phrase {
		fmt.Printf("%-2d    %U  %v  '%c'    %X\n",
				index, char, char, char, []byte(string(char)))
	}

	line := "xxx ddd ttt"
	i := strings.Index(line, " ")
	firstWord := line[0:]
	fmt.Printf("first word: %s\n", firstWord)
	fmt.Printf("first blank index: %d\n", i)
}
