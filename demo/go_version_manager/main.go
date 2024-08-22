package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	switch os.Args[1] {
	case "1.16":
		if err := exec.Command("zsh", "-c", "sudo ln -snf /usr/local/go1.16 /usr/local/go").Run(); err != nil {
			fmt.Println("switch failed, check if you has the specific version go1.16")
		}
	case "1.17":
		if err := exec.Command("zsh", "-c", "sudo ln -snf /usr/local/go1.17 /usr/local/go").Run(); err != nil {
			fmt.Println("switch failed, check if you has the specific version go1.17")
		}
	case "latest":
		if err := exec.Command("zsh", "-c", "sudo ln -snf /usr/local/go_latest /usr/local/go").Run(); err != nil {
			fmt.Println("switch failed, check if you has the go source code package")
		}
	default:
		fmt.Println("unsupported version")
	}
}
