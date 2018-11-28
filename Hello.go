package main

import "fmt"

func main() {
	printNames()
}

func printNames() {

	fmt.Println("First-Step")

	names := []string{"Yani", "Vlad", "Krasi", "Ivan", "Kris", "Naso", "Dimi", "Geri", "Lili", "Zoya"}

	for _, c := range names {
		fmt.Println("Name:", c)
	}
}
