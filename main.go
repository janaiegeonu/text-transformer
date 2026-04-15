package main

import (
	"fmt"
	"go-reloaded/helpers"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("ERROR: invalid command")
		fmt.Println("USAGE: go run . sample.txt result.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("ERROR: Failed reading file")
		fmt.Println("ERROR =", err)
		return
	}

	result := helpers.Compiler(string(data))

	err = os.WriteFile(outputFile, []byte(result+"\n"), 0644)
	if err != nil {
		fmt.Println("ERROR: Failed Writing file")
		fmt.Println("ERROR =", err)
		return
	}

}
