package main

import (
	"fmt"
)

func main() {
	fmt.Println("======== Basci GO ========")
	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)

	fmt.Println("======== variadic funciton ========")

	total := sumAll(10, 20, 89)

	fmt.Println(total)

	fmt.Println("======= function as parameter =========")
	sayHelloWithFilter("ALI", safeToLowerCase, spamFilter)
	sayHelloWithFilter("AnjINg", safeToLowerCase, spamFilter)

}
