package cmd

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestTreeCreateFileTree(t *testing.T) {
	// Create a temporary directory and some files and directories for testing
	tmpdir, err := ioutil.TempDir("", "tree-test1")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	err = os.Mkdir(filepath.Join(tmpdir, "dir1"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Mkdir(filepath.Join(tmpdir, "dir2"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(tmpdir, "file1.txt"), []byte("Hello, world!"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Test the CreateFileTree method
	tree := NewTree(1)
	err = tree.createFileTree(tmpdir, "", 0)
	if err != nil {
		t.Fatal(err)
	}

	expected := " ├─ dir1\n ├─ dir2\n └─ file1.txt\n"
	assert.Equal(t, expected, tree.fileTree)
	assert.Equal(t, 2, tree.folderCount)
	assert.Equal(t, 1, tree.fileCount)
}

func TestTreeExecuteCallsAll(t *testing.T) {
	// Create a temporary directory and some files and directories for testing
	tmpdir, err := ioutil.TempDir("", "tree-test2")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	err = os.MkdirAll(filepath.Join(tmpdir, "dir3", "dir4", "dir5"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(tmpdir, "dir3", "dir4", "dir5", "file1.txt"), []byte("Hello, world!"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	tree := NewTree(3)
	err = tree.Execute(tmpdir)
	if err != nil {
		return
	}
	expected := " └─ dir3\n   └─ dir4\n     └─ dir5\n       └─ file1.txt\n"
	assert.Equal(t, expected, tree.fileTree)
	assert.Equal(t, 3, tree.folderCount)
	assert.Equal(t, 1, tree.fileCount)
}
