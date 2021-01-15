package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var site string = `https://www.aozora.gr.jp`

/*
* use http.Get(URL) to get bytes of file from url
* create an empty file using os.Create
* use io.Copy to copy downloaded bytes to file created.
*
* unzip -O shift-jis fire.zip
 */
func main() {
	// Go to author page.
	resp, err := http.Get("https://www.aozora.gr.jp/index_pages/person20.html")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body := respToString(resp)
	mymap := makeMap(body)
	for k, v := range mymap {
		fmt.Println(k, v)
	}

	// zip, err := os.Create("blep.html")
	// if err != nil {
	// 	panic(err)
	// }
	// defer zip.Close()

	// b, err := io.Copy(zip, resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File size: ", b)
}

func respToString(resp *http.Response) string {
	// Convert the response body to a string.
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodystr := string(bodyBytes)
		return bodystr
	}
	return ""
}

func makeMap(ps string) map[string]string {
	sakuhin := make(map[string]string)

	// Get the names of the works: <a href....>WN</a>
	regWN := regexp.MustCompile(`\">.*?\</a>`)
	resWN := regWN.FindAllString(ps, -1)
	// Get the corresponding links to those works: <a href="../LINK">
	regL := regexp.MustCompile(`\".*?\"`)
	resL := regL.FindAllString(ps, -1)
	for i := 0; i < len(resWN); i++ {
		wn := strings.Trim(resWN[i], "\">") // Get rid of tag elements.
		wn = strings.Trim(wn, "</a>")

		addr := strings.Trim(resL[i], "\"") // Strip trailing "
		addr = strings.Trim(addr, "..\"")   // Strip leading .."

		sakuhin[wn] = site + addr
	}
	return sakuhin
}
