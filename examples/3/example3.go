package main

import (
	"fmt"

	"github.com/sunshibao/parallelizer"
)

func main() {
	group := parallelizer.NewGroup(parallelizer.WithPoolSize(10))
	defer group.Close()

	for i := 1; i <= 10; i++ {
		i := i
		group.Add(func() {
			fmt.Print(i, " ")
		})
	}

	err := group.Wait()

	fmt.Println()
	fmt.Println("Done")
	fmt.Printf("Error: %v", err)
}
