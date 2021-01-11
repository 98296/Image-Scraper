package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
* use http.Get(URL) to get bytes of file from url
* create an empty file using os.Create
* use io.Copy to copy downloaded bytes to file created.
 */
func main() {
	resp, err := http.Get("https://www.aozora.gr.jp/cards/000035/files/1578_44923.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	img, err := os.Create("story.html")
	if err != nil {
		panic(err)
	}
	defer img.Close()

	b, err := io.Copy(img, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("File size: ", b)
}
