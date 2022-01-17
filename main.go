package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"math"
)

const fileName string = "input.txt"

// "a??b" -> "abab"
// "a??a" -> "how long is the shortest sequence for both sides?" -> shortest + 1
// "a?b" -> "how long is the shortest sequence for both sides?" -> shortest + 1
// "a?a" -> "aba"

func main() {

	content := readFile()

	fmt.Printf(string(content))
	fmt.Printf("%d\n", findMinPossibleSequence(string(content)))
	printEvenOddQuestionMarkNumber(string(content))
}

func Solution(S string) int {

	minPossibleLength := findMinPossibleSequence(S)
	prevChar := ' '
	sequenceLength := 1

	for idx, s := range S {

		if s == prevChar {
			sequenceLength++
		} else {

			if prevChar == '?' {

				prevLetterDistance := sequenceLength + 1
				prevLetterIdx := idx - prevLetterDistance
				areRunesSameBothSides := idx >= prevLetterDistance && S[prevLetterIdx] == byte(s)

				if sequenceLength % 2 == 0 {
					fmt.Printf("even ? number found\n")
					if areRunesSameBothSides {prevLetterDistance
						fmt.Printf("same runes both sides\n")

						for i := 1; i < minPossibleLength; i++ {
							currentComparedPrecedingLetter
							if prevLetterIdx - i > 0 && idx + 1 < len(S) - 1 {
								if 
							}
						}
					}
				} else {
					fmt.Printf("odd ? number found\n")
					if !areRunesSameBothSides {
						fmt.Printf("runes are not same both sides\n")
					}
				}
			}

			sequenceLength = 1
		}
		
		prevChar = s
	}
	
	return minPossibleLength
}


func findMinPossibleSequence(S string) (result int) {
	result = 0
	prevChar := ' '
	sequenceLength := 1

	for _, s := range S {
		if s == prevChar && s != '?' {
			sequenceLength++
		} else {
			result = int(math.Max(float64(result), float64(sequenceLength)))
			sequenceLength = 1
		}
		prevChar = s
	}

	return
}

func readFile() []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return content
}