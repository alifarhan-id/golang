package main

import (
	"fmt"
	"strings"
)

type FunctionParameter func(string) string

// function say hallo with filter
func sayHelloWithFilter(name string, convertToLowerCase FunctionParameter, filter FunctionParameter) {
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
