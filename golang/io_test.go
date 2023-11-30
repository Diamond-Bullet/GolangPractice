package golang

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// here are some frequently used package provided by the official group.
// io: basic `reader`, `writer`, etc. pipe,
// os: file
// bufio: reading and writing with buffer.
// ioutil: read and write files. It's about to be obsolete. previous functions will be translocated to package `io`,`os`.
// archive/zip: read and write compressed files with buffer.
// io/fs: file system.

func TestFile(t *testing.T) {
	// flag `os.O_CREATE` used for creating a file when not existing. you can use os.Create() instead.
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()

	n, err := f.Write([]byte(`you are the beast!`))
	if err != nil {
		fmt.Printf("writing does not finish. %d bytes writed\n", n)
	}

	os.Remove("test1.txt")
}

func TestReadFile(t *testing.T) {
	// os.Open just for Reading.
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer file.Close()

	content := bytes.NewBuffer(nil)
	var buf [128]byte
	for {
		n, err1 := file.Read(buf[:])
		// complete reading file.
		if err1 == io.EOF {
			break
		}
		if err1 != nil {
			fmt.Println("read file err: ", err1)
			return
		}
		content.Write(buf[:n])
	}
	fmt.Println("content:", content.String())

	content1, err := os.ReadFile("test.txt")
	fmt.Printf("content1: %s, 				err: %s\n", content1, err.Error())
}

func TestZip(t *testing.T) {
	// TODO
	f, _ := os.OpenFile("demo.zip", os.O_RDWR|os.O_CREATE, 0666)
	defer f.Close()

	z := zip.NewWriter(f)
	defer z.Close() // `Flush` will execute in the function.
}

func TestStd(t *testing.T) {
	input := make([]byte, 0, 20)
	_, err := os.Stdin.Read(input)
	fmt.Printf("input: %s, 					err: %s", input, err.Error())

	_, err = os.Stdout.Write(input)
	fmt.Println("err:", err)
}

func TestBufIO(t *testing.T) {
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer f.Close()

	bufWriter := bufio.NewWriter(f)
	for i := 0; i < 10; i++ {
		bufWriter.Write([]byte("123\n"))
		bufWriter.Flush()
	}

	// bufio.NewReader(f) = bufio.NewReaderSize(f, 4096)
	reader := bufio.NewReaderSize(f, 4096)

	reader.ReadBytes('\n')
}
