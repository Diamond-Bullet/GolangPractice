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

func TestFile(t *testing.T) {
	// flag `os.O_CREATE` used for creating a file when not existing. you can use os.Create() instead.
	f, _ := os.OpenFile("test.txt", os.O_RDWR | os.O_CREATE, 0666)
	defer f.Close()

	f.Write([]byte(`you are the beast!`))
}

func TestZip(t *testing.T) {
	f, _ := os.OpenFile("demo.zip", os.O_RDWR | os.O_CREATE, 0666)
	defer f.Close()

	z := zip.NewWriter(f)
	defer z.Close() // `Flush` will execute in the function.
}

func TestStd(t *testing.T) {
	input := make([]byte, 0, 20)
	os.Stdin.Read(input)

	os.Stdout.Write(input)
}