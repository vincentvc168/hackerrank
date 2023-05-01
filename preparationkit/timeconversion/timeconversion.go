package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	//fmt.Println(s)
	//fmt.Println(string(s[1:len(s)-1]))
	// fmt.Println(string(s[len(s)-1:len(s)]))
	// fmt.Println(string(s[0:2]))

	if string(s[len(s)-2:len(s)]) == "AM" && string(s[0:2]) == "12" {
		s = strings.Replace(s, "12", "00", 1)
		s = string(s[:len(s)-2])
		temp, _ := strconv.Atoi(string(s[0:2]))
		fmt.Println(temp)
		return s
	}
	if string(s[len(s)-2:len(s)]) == "PM" && string(s[0:2]) == "12" {
		s = string(s[:len(s)-2])
		return s
	}
	if string(s[len(s)-2:len(s)]) == "AM" && !(string(s[0:2]) == "12") {
		s = string(s[:len(s)-2])

		return s
	}
	if string(s[len(s)-2:len(s)]) == "PM" && !(string(s[0:2]) == "12") {
		hour_int, _ := strconv.Atoi(string(s[0:2]))
		hour_int += 12
		hour_string := strconv.Itoa(hour_int)
		s = strings.Replace(s, string(s[0:2]), hour_string, 1)
		s = string(s[:len(s)-2])
		return s
	}

	return s
}

func main() {
	fmt.Println(timeConversion("07:05:45AM"))

}
