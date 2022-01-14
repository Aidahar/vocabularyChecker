package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var f, word string
	var bad int
	//answer := map[string]bool{}
	var answer []string
	// read from input
	fmt.Scan(&f)

	// open file
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read word by word
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read word
		str := scanner.Text()
		// fill map words as key
		answer = append(answer, str)
	}

	// check error in scan
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Scan(&word)
		if word == "exit" {
			fmt.Println("Bye!")
			return
		}
		//fmt.Println("bad in start ", bad)
		for i := 0; i < len(answer); i++ {
			wordToLow := strings.ToLower(word)
			if answer[i] == wordToLow {
				fmt.Println(strings.Repeat("*", len(word)))
				bad += 1
			}
		}
		//fmt.Println("bad after check ", bad)
		//fmt.Println("word after check", word)
		if bad == 0 {
			fmt.Println(word)
		}
		bad = 0
		//fmt.Println("bad after zeroed ", bad)
	}
}
