# go_tree

# A simple CLI that shows a tree of folder/files. 

# Usage
go run ./src/main.go <file_path> -d <depth of folders> 

Example output for this repository: 
```bash
go run src/main.go -d 0
/Users/schanjr/Github/go_tree
 ├─ .git
 ├─ .gitignore
 ├─ .idea
 ├─ LICENSE
 ├─ Makefile
 ├─ README.md
 ├─ executable
 ├─ go.mod
 ├─ go.sum
 └─ src

4 Directories, 6 Files 

```

# build 
`make build`

# test 
`make test`


### References
https://michaelsoolee.com/tree-tool/