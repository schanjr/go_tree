package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Tree struct {
	MaxDepth    int
	fileCount   int
	folderCount int
	fileTree    string
}

func NewTree(depth int) *Tree {
	return &Tree{MaxDepth: depth}
}

func (t *Tree) Execute(root string) error {
	err := t.createFileTree(root, "", 0)
	if err == nil {
		fmt.Printf(t.fileTree)
		fmt.Printf("\n%d Directories, %d Files \n", t.folderCount, t.fileCount)
	}
	return err
}

func (t *Tree) createFileTree(root, prefix string, depth int) error {
	if t.MaxDepth >= 0 && depth > t.MaxDepth {
		return nil
	}
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %s", root, err)
	}
	for i, file := range files {
		filename := file.Name()
		if i == len(files)-1 {
			t.fileTree += prefix + " └─ " + filename + "\n"
		} else {
			t.fileTree += prefix + " ├─ " + filename + "\n"
		}
		if file.IsDir() {
			t.folderCount += 1
			err := t.createFileTree(filepath.Join(root, filename), prefix+"  ", depth+1)
			if err != nil {
				return err
			}
		} else {
			t.fileCount += 1
		}
	}

	return nil
}
