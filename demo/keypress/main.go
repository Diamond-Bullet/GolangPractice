package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"time"
)

// TODO the pkg robotgo can not be compiled successfully.
func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(3 * time.Second)

	robotgo.TypeStr(text, 1000)
}
