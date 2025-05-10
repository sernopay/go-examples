# Implementing Worker Pool Pattern in Go
Worker Pool Pattern is a concurrency design pattern that allows you to manage a pool of worker goroutines to perform tasks concurrently. This pattern is commonly used to handle a large number of tasks efficiently. It prevents overwhelming system resources by limiting the number of concurrent tasks. Additionally, by reusing workers (goroutines), you can avoid the overhead of creating and destroying goroutines for each task.

In this example, we will implement a simple worker pool pattern in Go to simulate the web scraping tasks. Let's say we have many links to scrape, and we want to scrape them concurrently to speed up the process. If we try to scrape them all at once, it may lead to resource exhaustion or network congestion. Instead, we can use a worker pool to limit the number of concurrent scraping operations to prevent these issues.

## Learning Objectives
- Understand the worker pool pattern and its benefits.
- Implement a worker pool in Go to handle concurrent tasks.

## Implementation
First, let's create a function that will create a pool of workers, pass tasks to them, and wait for all tasks to complete.
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
In this code, we create a channel `ch` to send links to the workers. We also use a `sync.WaitGroup` to wait for all workers to finish their tasks. Next, we create the specified number of workers based on the `numberOfWorker` parameter. Then we track the number of workers using `wg.Add(1)` and start each worker in a goroutine. Each worker will receive links from the channel and process them. \
In the second loop, we send each link to the channel `ch` gradually depending on the available workers until all links are sent. After all links are sent, we close the channel by calling `close(ch)` to signal that no more links will be sent. Finally, we call `wg.Wait()` to block the main goroutine until all workers have finished processing the links. \

Now, let's implement the worker function that will process the links. The worker function will receive links from the channel and perform the scraping operation.
```go
func workerWebScraper(linkChan chan string, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	expensiveOperation(workerID)
	for link := range linkChan {
		scrapWeb(link, workerID)
	}
}
```
In this code, we define the `workerWebScraper` function that takes a channel of links, a wait group, and the worker ID as parameters. The `defer wg.Done()` statement ensures that the wait group counter is decremented when the worker finishes its task. The `expensiveOperation(workerID)` simulates an expensive operation that each worker needs to perform before starting to scrape links. Then, we loop through the channel of links, this will block until a link is available and stop when the channel is closed. For each link received, we call the `scrapWeb(link, workerID)` function to perform the scraping operation. \

Lastly, we will implement the dummy `expensiveOperation` and `scrapWeb` functions to simulate the scraping process.
```go
func expensiveOperation(workerID int) {
	fmt.Printf("WorkerID %d Run Expensive Operation \n", workerID)
}

func scrapWeb(link string, workerID int) {
	time.Sleep(1 * time.Second) // to simulate the process
	fmt.Printf("workerID %d finish scrap link %s \n", workerID, link)
}
```
In `scrapWeb`, we simulate the scraping process by sleeping for 1 second to mimic the time taken to scrape a link. After that, we print the worker ID and the link that was scraped.

## Run the Code
To run the code, you can create a main function that initializes the worker pool and provides a list of links to scrape.
```go
func main() {
	links := []string{
		"https://www.example.com/1", "https://www.example.com/2", "https://www.example.com/3", "https://www.example.com/4",
		"https://www.example.com/5", "https://www.example.com/6", "https://www.example.com/7", "https://www.example.com/8",
	}
	runWorkerpoolWebScraper(3, links)
}
```
In this example, we create a list of links to scrape and call the `runWorkerpoolWebScraper` function with 3 workers. The pool will process the links concurrently, with a maximum of 3 workers running at any given time. This will help to speed up the scraping process while preventing resource exhaustion. \
Here is the output of the code:
```bash
WorkerID 2 Run Expensive Operation 
WorkerID 0 Run Expensive Operation 
WorkerID 1 Run Expensive Operation
workerID 2 finish scrap link https://www.example.com/1 
workerID 1 finish scrap link https://www.example.com/3 
workerID 0 finish scrap link https://www.example.com/2
workerID 1 finish scrap link https://www.example.com/5 
workerID 2 finish scrap link https://www.example.com/4 
workerID 0 finish scrap link https://www.example.com/6
workerID 1 finish scrap link https://www.example.com/7 
workerID 2 finish scrap link https://www.example.com/8
```
The output can vary depending on the order in which the workers finish initializing and processing the links. The important thing to note is that the workers are processing the links concurrently, and the expensive operation is performed only once for each worker.

## Summary
In this example, we implemented the worker pool pattern in Go to handle concurrent tasks efficiently. By using a worker pool, we can limit the number of concurrent tasks and reuse workers, which helps to improve efficiency and prevent resource exhaustion.