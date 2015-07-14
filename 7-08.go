/*
 * DAILY PROGRAMMER - 7/08/2015
 * RC4 Cipher
 * https://www.reddit.com/r/dailyprogrammer/comments/3chvxy/20150708_challenge_222_intermediate_simple_stream/
 * Author: Muhan Guclu
 * Description: String encryption and decryption via linear congruential generator (LCG)
 */

package main

import "fmt"

func lcg(seed uint32) func() uint32 {
	r := seed
	a, c, m := uint32(214013), uint32(2531011), uint32(1<<31)
	return func() uint32 {
		r = (a*r + c) % m
		return r
	}
}

func parse(msg string, key uint32) string {
	ret := ""
	rng := lcg(key)
	for _, char := range msg {
		ret += string((uint32(char) ^ rng()) & uint32(0xFF))
	}
	return ret
}

func main() {
	input := "Attack at dawn"
	key := uint32(234)

	fmt.Println("MESSAGE: ", input, " - KEY: ", key)
	fmt.Println("")

	ciphertext := parse(input, key)
	fmt.Println("ENCRYPTED: ", ciphertext)

	orig := parse(ciphertext, key)
	fmt.Println("DECRYPTED: ", orig)
}
