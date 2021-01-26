package main

import (
	"fmt"
	"os"
)

func main() {
	name := ""
	//_ es para ignorar el indice en el for

	if len(os.Args) > 1 {
		for _, word := range os.Args[1:] {
			name = fmt.Sprint(name, " ", word)

		}
		fmt.Println("Welcome to the jungle", name, "!")
	} else {
		fmt.Println("Error! No name!")
	}

}
