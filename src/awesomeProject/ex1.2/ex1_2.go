package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {

		fmt.Println(i, os.Args[i])
		fmt.Printf("%d: %s\n", i, os.Args[i])
	}
}
