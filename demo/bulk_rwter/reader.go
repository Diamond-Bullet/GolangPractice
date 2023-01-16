package main

import (
	"io"
	"os"
	"syscall"
)

func ReadForward(filepath string) []byte {
	f, err := os.OpenFile(filepath, syscall.O_RDONLY, 0644)
	if err != nil {
		Fatal("open file error", err)
	}
	defer f.Close() // 文件打开后需要关闭，释放描述符

	b, err := io.ReadAll(f)
	if err != nil {
		Error("read file error", err)
		return nil
	}

	return b
}

func ReadForwardV1(filepath string) []byte {
	b, err := os.ReadFile(filepath)
	if err != nil {
		Error("read file error", err)
		return nil
	}
	return b
}

func ReadForwardV2(filepath string) []byte {
	return nil
}
