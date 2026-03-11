package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	buah := []string{"Apel", "Jeruk", "Pisang"}
	for index, value := range buah {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}
