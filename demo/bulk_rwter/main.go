package main

import "flag"

var (
	mod    = flag.String("mod", "", "read/write file with random string")
	method = flag.String("method", "", "forward/block in file")
	file   = flag.String("file", "test.txt", "file path")
)

func main() {
	flag.Parse()

	var modMask uint64

	switch *mod {
	case "read":
		modMask |= 1 << 32
	case "write":
		modMask |= 1 << 33
	}

	switch *method {
	case "forward":
		modMask |= 1 << 0
	case "block":
		modMask |= 1 << 1
	}

	switch modMask {
	case ModReadForward:
		ReadForward(*file)
	}
}
