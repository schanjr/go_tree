package main

import (
	"flag"
	"fmt"
	"github.com/schanjr/go_tree/src/cmd"
	"os"
	"path/filepath"
)

func main() {
	// Define flags
	var depth int
	flag.IntVar(&depth, "d", -1, "depth of the tree")
	flag.Parse()

	// Get the path argument
	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	root, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(root)
	tree := cmd.NewTree(depth)
	err = tree.Execute(root)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
