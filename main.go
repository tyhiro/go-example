package main

import (
	"flag"
	"fmt"
	"log"

	"go-example/worker"
)

func main() {
	workerCount := flag.Int("parallel", 10, "limit the number of parallel requests")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("Usage: myhttp [-parallel N] URL1 [URL2...]")
		return
	}

	processUrls(*workerCount, urls)
}

func processUrls(workerCount int, urls []string) {
	log.SetFlags(0)
	wp := worker.New(workerCount)

	go wp.GenerateFromUrls(urls)
	go wp.Run()

	for {
		select {
		case r, ok := <-wp.Results():
			if !ok {
				continue
			}

			log.Println(r.String())
		case <-wp.Done:
			return
		default:
		}
	}
}
