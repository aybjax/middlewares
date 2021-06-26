package main

import "fmt"

func main() {
	generator := generateGenerator()

	for i := 0; i<5; i++ {
		fmt.Print(generator(), "\t")
	}

	fmt.Println()
}

func generateGenerator() func() int {
	i := 0

	return func () int  {
		i++
		return i
	}
}