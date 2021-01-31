# aozora-scraper
This is a small concurrent webscraper that downloads all works of a single author from https://www.aozora.gr.jp/

## !! WARNING !!
This was created for the intent of how to use go routines. If you use this program, do NOT use it maliciously.
Excessive requests could possibly result in your IP being blacklisted.

## How to use this program
It has the following command line flags:
- ap: author page, such as https://www.aozora.gr.jp/index_pages/person11.html The program accepts any link
provided to it to be of that format, where each author's url is distinguished by the person number ("/personxx.html")
- dn: directory name. This is the location of where the program will save all the files. Example, "works" would save
the files to a folder named works. Note, the directory name must not exist already, it must be a new directory.

An example of how one would run it:
```
go run ./cmd -ap=https://www.aozora.gr.jp/index_pages/person35.html -dn=osamu
```
