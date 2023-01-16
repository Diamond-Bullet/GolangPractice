package main

import "fmt"

func Error(message ...interface{}) {
	fmt.Printf("%v\n", message)
}

func Fatal(message ...interface{}) {
	panic(fmt.Sprintf("%v\n", message))
}
