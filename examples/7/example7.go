package main

import (
	"fmt"
	"time"

	"github.com/sunshibao/parallelizer"
)

func main() {
	group := parallelizer.NewGroup()
	defer group.Close()

	group.Add(func() {
		fmt.Println("Finished work")
	})

	fmt.Println("We did not wait!")

	time.Sleep(time.Second)
}
