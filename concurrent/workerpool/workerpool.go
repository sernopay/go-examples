package main

import (
	"fmt"
	"sync"
	"time"
)

func expensiveOperation(workerID int) {
	fmt.Printf("WorkerID %d Run Expensive Operation \n", workerID)
}

func scrapWeb(link string, workerID int) {
	time.Sleep(1 * time.Second) // to simulate the process
	fmt.Printf("workerID %d finish scrap link %s \n", workerID, link)
}

func workerWebScraper(linkChan chan string, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	expensiveOperation(workerID)
	for link := range linkChan {
		scrapWeb(link, workerID)
	}
}

func runWorkerpoolWebScraper(numberOfWorker int, links []string) {
	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < numberOfWorker; i++ {
		wg.Add(1)
		go workerWebScraper(ch, &wg, i)
	}

	for _, link := range links {
		ch <- link
	}

	close(ch)
	wg.Wait()
}

func main() {
	links := []string{
		"https://www.example.com/1", "https://www.example.com/2", "https://www.example.com/3", "https://www.example.com/4",
		"https://www.example.com/5", "https://www.example.com/6", "https://www.example.com/7", "https://www.example.com/8",
	}
	runWorkerpoolWebScraper(3, links)
}
