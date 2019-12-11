package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("This is version v0.0.2")

	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
}
