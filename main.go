package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var site string = `https://www.aozora.gr.jp`

// fetchHTML takes in a url and returns the responses body.
func fetchHTML(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error when fetching URL: %v\n", err)
	}
	return resp.Body
}

func tokenize(body io.ReadCloser) map[string]string {
	retval := make(map[string]string)
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return retval
		}

		if tt == html.StartTagToken { // I'm looking for first occurence of <ol>
			token := z.Token()
			if token.Data == "ol" {
				for { // Inside <ol></ol>
					// Find me all <a>
					for tt != html.StartTagToken || token.Data != "a" {
						tt = z.Next()
						token = z.Token()
						if tt == html.EndTagToken && token.Data == "ol" {
							fmt.Println(token.Data)
							return retval
						}
					}
					// tt = z.Next() // Text, most likely '\n'.
					// tt = z.Next() // Start tag, <li>
					// tt = z.Next() // Start tag <a>
					// //tt = z.Next() // Is <a>THIS</a>
					// token = z.Token()

					// if token.Data == "a" {
					// 	fmt.Println(token.Data)
					// }
					// ta := token.Attr
					// fmt.Println(ta)
					// return

					// Get the link first.
					link := token.Attr
					// Now move onto get the name of the work, <a>TITLE</a>.
					tt = z.Next()
					token = z.Token()
					title := token.Data
					// Add the title with it's corresponding link to the map.
					if len(link) > 0 {
						if link[0].Key == "href" {
							retval[title] = link[0].Val
						}
					}
				}

			}
		}
	}
}

func main() {
	// Go to author page and get the HTML response.
	body := fetchHTML("https://www.aozora.gr.jp/index_pages/person20.html")
	defer body.Close()

	mm := tokenize(body)
	i := 1
	for key, val := range mm {
		fmt.Println(i, key, val)
		i++
	}
}
