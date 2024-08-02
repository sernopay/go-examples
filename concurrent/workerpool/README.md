# Worker Pool Web Scraper

This Go program demonstrates a worker pool pattern to perform web scraping tasks concurrently. The program creates a pool of worker goroutines to scrape multiple web links efficiently.

## Overview

The program defines several functions to simulate web scraping and manage worker goroutines. It uses channels and wait groups to coordinate the work among multiple workers.

## Code Explanation
### Expensive Operation Simulation
```go
func expensiveOperation(workerID int) {
    fmt.Printf("WorkerID %d Run Expensive Operation \n", workerID)
}
```
This function simulates an expensive operation performed by a worker. It prints a message indicating the worker ID.

### Web Scraping Simulation
```go
func scrapWeb(link string, workerID int) {
    time.Sleep(1 * time.Second) // to simulate the process
    fmt.Printf("workerID %d finish scrap link %s \n", workerID, link)
}
```
This function simulates the web scraping process. It takes a link and a worker ID, sleeps for 1 second to simulate the scraping process, and then prints a message indicating the worker ID and the link.

### Worker Function
```go
func workerWebScraper(linkChan chan string, wg *sync.WaitGroup, workerID int) {
    defer wg.Done()
    expensiveOperation(workerID)
    for link := range linkChan {
        scrapWeb(link, workerID)
    }
}
```
This function defines the worker goroutine. It performs the following steps:
* Calls expensiveOperation to simulate an initial setup.
* Iterates over the linkChan channel to receive links and calls scrapWeb to process each link.
* Uses wg.Done() to signal the wait group that the worker has finished its task.

### Worker Pool Runner
```go
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
```
This function sets up and runs the worker pool. It performs the following steps:

* Creates a channel `ch` to send links to workers.
* Initializes a wait group `wg` to wait for all workers to finish.
* Starts `numberOfWorker` worker goroutines.
* Sends each link from the `links` slice to the `ch` channel.
* Closes the channel after sending all links.
* Waits for all workers to finish using `wg.Wait()`.

### Main Function
```go
func main() {
    links := []string{
        "https://www.example.com/1", "https://www.example.com/2", "https://www.example.com/3", "https://www.example.com/4",
        "https://www.example.com/5", "https://www.example.com/6", "https://www.example.com/7", "https://www.example.com/8",
    }
    runWorkerpoolWebScraper(3, links)
}
```
The main function defines a list of links to be scraped and calls runWorkerpoolWebScraper with 3 workers to process the links.

## Running the Program
To run the program, use the following command:
```bash
go run workerpool.go
```

### Example Output
```bash
WorkerID 2 Run Expensive Operation
WorkerID 1 Run Expensive Operation
WorkerID 0 Run Expensive Operation
workerID 0 finish scrap link https://www.example.com/3
workerID 1 finish scrap link https://www.example.com/2
workerID 2 finish scrap link https://www.example.com/1
workerID 1 finish scrap link https://www.example.com/5
workerID 2 finish scrap link https://www.example.com/6
workerID 0 finish scrap link https://www.example.com/4
workerID 2 finish scrap link https://www.example.com/8
workerID 1 finish scrap link https://www.example.com/7
```
