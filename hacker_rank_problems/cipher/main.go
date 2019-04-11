package main

import (
	"fmt"
)

func main() {
	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)
	/*
		alphabetLower := "abcdefghijklmnopqrstuvwxyz"
		alphabetUpper := strings.ToUpper("abcdefghijklmnopqrstuvwxyz")

		result := ""

		for _, char := range input {
			switch {
			case strings.IndexRune(alphabetLower, char) >= 0:
				result = result + string(rotate(char, delta, []rune(alphabetLower)))
			case strings.IndexRune(alphabetUpper, char) >= 0:
				result = result + string(rotate(char, delta, []rune(alphabetUpper)))

			default:
				result = result + string(char)
			}
		}

		fmt.Println(result)
		fmt.Println(string(cipher('Z', 3)))
	*/

	var result []rune

	for _, char := range input {
		result = append(result, cipher(char, delta))
	}
	fmt.Println(string(result))
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26

	return rune(tmp + base)
}

/*
func rotate(str rune, delta int, key []rune) rune {
	index := strings.IndexRune(string(key), str)

	if index < 0 {
		panic("index < 0")
	}
	index = (index + delta) % len(key)

	return key[index]
}

*/
