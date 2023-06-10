package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	//fileUrl := "http://192.168.2.89/capture"
	fileUrl := "https://i.stack.imgur.com/wOtFx.jpg"
	err := DownloadFile("/tmp/boiler.jpg", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded /tmp/boiler.jpg")

	cmd := exec.Command("/bin/zsh", "-c", "/opt/homebrew/bin/ssocr -d -1 -f black /tmp/boiler.jpg")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println("Found Text: " + string(out))

}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err

}
