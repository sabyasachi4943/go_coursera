package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var sourceString string
	fmt.Println("Please inupt string")
	//reading input from command promot. Input may contain spaces.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sourceString = scanner.Text()
		break
	}

	sourceString = strings.ToLower(sourceString)

	if strings.HasPrefix(sourceString, "i") &&
		strings.HasSuffix(sourceString, "n") &&
		strings.Contains(sourceString, "a") {

		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
