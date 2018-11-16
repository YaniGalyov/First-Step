package main

import "fmt"

func main() {

    names := []string{"Yani", "Vlad", "Krasi", "Ivan", "Kris", "Naso", "Dimi", "Geri", "Lili", "Zoya"}

    for i, c := range names {
        fmt.Println("", i)
        fmt.Println("", c)
    }
}