package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("test 2")
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("halo")
	}
}

