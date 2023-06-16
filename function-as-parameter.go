package main

import (
	"fmt"
	"strings"
)

// function say hallo with filter
func sayHelloWithFilter(name string, convertToLowerCase func(string) string, filter func(string) string) {
	lowerCaseName := convertToLowerCase(name)
	nameFiltered := filter(lowerCaseName)

	fmt.Println("Hello", nameFiltered)
}

func safeToLowerCase(str string) string {
	return strings.ToLower(str)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "...."
	} else {
		return name
	}
}
