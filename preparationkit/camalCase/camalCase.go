package main

import (
	"fmt"
	"unicode"
	"strings"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 input := "C;M;mouse pad"
 var output []string

 //output[1] = rune("S")
	// Split
	if string(input[0]) == "S" {
		for i, r := range input {
			if i < 4 {
				continue
			}
			if unicode.IsUpper(r) {
				output = append(output, " ", strings.ToLower(string(r)))
				continue
			}
			output = append(output, strings.ToLower(string(r)))
		}
		
	}

	// Split Method
	if string(input[0]) == "S" && string(input[2]) == "M" {
		output = output[:len(output)-2]
	}

	//Combine
	if string(input[0]) == "C" {
		var skip bool
		for i, r := range input {
			if i < 4 {
				continue
			}

			if skip {
				skip = false
				continue
			}

			if string(input[2]) == "C" && i == 4{
				output = append(output, " ", strings.ToUpper(string(r)))
				continue
			}

			if string(input[i]) == " " {
				output = append(output, strings.ToUpper(string(input[i+1])))
				skip = true
				continue
			}
			


			output = append(output, strings.ToLower(string(r)))
		}
		
	}

	if string(input[0]) == "C" && string(input[2]) == "M" {
		output = append(output, "()")
	}
 
 fmt.Println(strings.Join(output,""))
}