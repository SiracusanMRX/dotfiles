package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("open", "https://github.com/Siracusant")
	err := cmd.Run()
	if err != nil {
		fmt.Println("FAILED : ", err)
		return
	}
}
