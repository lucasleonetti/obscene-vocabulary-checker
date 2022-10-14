package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var FileLines []string

func main() {

	//Enter the path and open our file
	file, err := os.Open(readFile())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// we read the sentence
	inputS := readSentence()

	for !(strings.HasPrefix(inputS, "exit")) {
		checkedLine := wordsChecker(file, inputS)
		fmt.Println(checkedLine)
		inputS = readSentence()
	}

	fmt.Println("Bye!")
}

func readFile() string {
	var rF string
	fmt.Println("Enter file's path you wanted to check: ")
	fmt.Scanln(&rF)
	return rF
} // we read the file, return the file.

func readSentence() string {
	reader := bufio.NewReader(os.Stdin)
	rS, _ := reader.ReadString('\n')
	return rS
} // we input the sentence, return the sentence.

func wordsChecker(file *os.File, inputSen string) string {

	inputSenArr := make([]string, len(inputSen)) //we initialize a new slice for our input sentence
	inputSenArr = strings.Fields(inputSen)       //we convert the input in slices of words for processing
	scanner := bufio.NewScanner(file)            //read the file line by line, this file only contains our bad choose words
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		FileLines = append(FileLines, scanner.Text())
	} //we save the censored words in our file for processing after

	for _, line := range FileLines { //we travel our new slice with the censored words
		for i, word := range inputSenArr {
			if strings.EqualFold(line, word) {
				inputSenArr[i] = strings.Repeat("*", len(word)) //replace the bad word by censored symbol
			}
		}
	}

	processedLine := strings.Join(inputSenArr, " ")
	return processedLine

} // verifying line by line, return sentence with censored words specified in the file.
