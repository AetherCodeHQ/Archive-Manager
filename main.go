package main

import (
	"fmt"
	"os"
)

// archive_manager - Manage archive formats
func archive_manager(path string) {
	fmt.Println("========================================")
	fmt.Println("  Archive-Manager")
	fmt.Println("  Manage archive formats")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	archive_manager(path)
}
