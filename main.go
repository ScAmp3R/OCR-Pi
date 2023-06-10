package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("/bin/zsh", "-c", "wget -o boiler.jpg http://192.168.2.245/capture")
	cmd := exec.Command("/bin/zsh", "-c", "/opt/homebrew/bin/ssocr -d -1 -f black boiler.jpg")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(string(out))

}
