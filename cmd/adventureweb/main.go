package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "gopher.json", "JSON file with our adventure stories")
	flag.Parse()

	fmt.Printf("Using the story in %s", *filename)
}
