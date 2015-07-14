/*
 * DAILY PROGRAMMER - 7/06/2015
 * Balancing Words
 * Author: Muhan Guclu
 * Description: https://www.reddit.com/r/dailyprogrammer/comments/3c9a9h/20150706_challenge_222_easy_balancing_words/
 */

package main

import (
	"fmt"
	"strconv"
)

func getAlphaPos(l rune) int {
	return int(l) - 64
}

func sumSides(a string, b string) (int, int) {
	sumA, sumB := 0, 0
	for idxA, charA := range a {
		sumA += getAlphaPos(charA) * (len(a) - idxA)
	}
	for idxB, charB := range b {
		sumB += getAlphaPos(charB) * (idxB + 1)
	}
	return sumA, sumB
}

func getCenter(w string) string {
	result := ""
	for idx, letter := range w {
		sideA, sideB := w[:idx], w[(idx+1):]
		sumA, sumB := sumSides(sideA, sideB)
		if sumA-sumB == 0 {
			result = fmt.Sprint(sideA, " ", strconv.QuoteRuneToASCII(letter), " ", sideB, " - ", sumA)
			break
		}
	}
	if result == "" {
		result = fmt.Sprint(w, " DOES NOT BALANCE")
	}

	return result
}

func main() {
	challengeInput := []string{
		"STEAD",
		"CONSUBSTANTIATION",
		"WRONGHEADED",
		"UNINTELLIGIBILITY",
		"SUPERGLUE",
	}
	for _, input := range challengeInput {
		fmt.Println(getCenter(input))
	}
}
