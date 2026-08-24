package main

import (
	"archive/zip"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: archive-manager <file.zip> [list|extract]")
		os.Exit(1)
	}
	r, err := zip.OpenReader(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer r.Close()
	mode := "list"
	if len(os.Args) > 2 {
		mode = os.Args[2]
	}
	switch mode {
	case "list":
		for _, f := range r.File {
			fmt.Printf("%10d  %s\n", f.UncompressedSize64, f.Name)
		}
		fmt.Printf("\n%d entries\n", len(r.File))
	case "extract":
		for _, f := range r.File {
			fmt.Println("would extract:", f.Name)
		}
	default:
		fmt.Println("unknown mode:", mode)
		os.Exit(1)
	}
}
