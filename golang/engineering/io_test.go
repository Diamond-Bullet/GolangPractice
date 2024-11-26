package engineering

import (
	"GolangPractice/lib/logger"
	"archive/zip"
	"bufio"
	"bytes"
	"github.com/gookit/color"
	"io"
	"os"
	"testing"
)

// here are some frequently used package provided by the official team.
// io: basic `reader`, `writer`, etc. pipe,
// os: file
// bufio: reading and writing with buffer.
// ioutil: read and write files. It's about to be obsolete. previous functions will be translocated to package `io`,`os`.
// archive/zip: read and write compressed files with buffer.
// io/fs: file system.

func TestWorkDir(t *testing.T) {
	// Get current work directory
	workDir, err := os.Getwd()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("workDir:", workDir)

	// change current directory
	err = os.Chdir("/root")
	if err != nil {
		logger.Error(err)
		return
	}
}

func TestFile(t *testing.T) {
	// get user home dir
	homeDir, _ := os.UserHomeDir()

	// flag `os.O_CREATE` used for creating a file when not existing. you can use os.Create() instead.
	f, err := os.OpenFile(homeDir+"test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logger.Error("open file err:", err)
		return
	}
	defer f.Close()

	n, err := f.Write([]byte(`you are the beast!`))
	if err != nil {
		logger.Errorf("writing does not finish. %d bytes writed\n", n)
	}

	os.Remove("test1.txt")
}

func TestReadFile(t *testing.T) {
	// METHOD 1: os.Open just for Reading.
	file, err := os.Open("test.txt")
	if err != nil {
		logger.Error("open file err: ", err)
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
			logger.Error("read file err:", err1)
			return
		}
		content.Write(buf[:n])
	}
	logger.Info("content:", content.String())

	// METHOD 2
	content1, err := os.ReadFile("test.txt")
	logger.Infof("content1: %s, 				err: %s\n", content1, err.Error())
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
	logger.Infof("input: %s, 					err: %s", input, err.Error())

	_, err = os.Stdout.Write(input)
	logger.Error("Write err:", err)

	// TODO doesn't work well
	// redirect output to particular file
	redirectFile, err := os.OpenFile("/tmp/redirectFile", os.O_RDWR|os.O_CREATE|os.O_SYNC, 0766)
	if err != nil {
		color.Redln("failed to open redirectFile file, reason: ", err.Error())
		return
	}
	defer redirectFile.Close()
	os.Stdout = redirectFile
	os.Stderr = redirectFile

	logger.Info("output after redirect")
}

func TestBufIO(t *testing.T) {
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logger.Error("open file err: ", err)
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
