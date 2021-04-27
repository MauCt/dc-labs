// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
//
// Crawl3 adds support for depth limiting.
//
package main

import (
	"flag"

	"log"
	"os"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

type data struct {
	depth int
	urls  []string
}

func crawl(url string) []string {
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	depth := flag.Int("depth", 1, "Depth of crawling")
	results := flag.String("results", "results.txt", "name of the file")
	flag.Parse()

	if *depth == 0 || *results == "none" {
		log.Fatal("ERROR")
	}

	var n int // number of pending sends to worklist
	n++

	file, err := os.Create(*results)
	if err != nil {

		log.Print(err)
	}

	defer file.Close()
	// Start with the command-line arguments.
	worklist := make(chan data)

	go func() {
		worklist <- data{depth: 0, urls: flag.Args()}
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		currentDepth := list.depth
		for _, link := range list.urls {

			if currentDepth <= *depth {
				if !seen[link] {
					file.Write([]byte(link + "\n"))
					seen[link] = true

					go func(link string) {
						worklist <- data{urls: crawl(link), depth: currentDepth + 1}
					}(link)

					n++
				}

			}

		}
	}
}

//!-
