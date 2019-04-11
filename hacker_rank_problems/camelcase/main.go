package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1

	for _, characther := range input {
		str := string(characther)

		if strings.ToUpper(str) == str {
			//its a capital letter!
			answer++
		}

	}

	fmt.Println(answer)

}
