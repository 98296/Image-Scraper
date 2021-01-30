package main

import (
	"AozoraScraper/scraper"
	"flag"
	"os"
)

func main() {
	ap := flag.String("ap", "https://www.aozora.gr.jp/index_pages/person20.html", "The url to the author's page")
	dn := flag.String("dn", "works", "The directory you want to save the author's work, too. Must be a new folder")
	flag.Parse()

	// Go to author page and get the HTML response.
	body, err := scraper.FetchHTML(*ap)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	// Tokenize the author page into a map of URLs.
	mm := scraper.ParseAP(body)

	// Create a directory (directory name) to save the work to.
	err = os.Mkdir(*dn, 0755)
	if err != nil {
		panic(err)
	}

	// Now download all the zips from that map of links and save to the provided
	// directory name.
	err = scraper.DownloadWorks(*dn, mm)
	if err != nil {
		panic(err)
	}
}
