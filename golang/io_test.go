package golang

import (
	"archive/zip"
	"os"
	"testing"
)

// golang common IO libraries: https://www.cnblogs.com/zhichaoma/p/12509984.html

// here are some frequently used package provided by the official group.
// io: basic `reader`, `writer`, etc. pipe,
// os: file
// bufio: reading and writing with buffer.
// ioutil: read and write files. It's about to be obsolete. previous functions will be translocated to package `io`,`os`.
// archive/zip: read and write compressed files with buffer.
// io/fs: file system.

func TestArchiveZip(t *testing.T) {
	f, _ := os.Open("demo")
	defer f.Close()

	z := zip.NewWriter(f)
	z.Close() // `Flush` will execute in the function.
}
