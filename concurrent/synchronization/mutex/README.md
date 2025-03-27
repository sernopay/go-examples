# Mutex
We've seen how channels are great for communication among goroutines. But what if we don't need communication?  What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?
This concept is called mutual exclusion. and the conventional name for the data structure that provides it is mutex.

Go's standard library provides mutual exclusion with sync.Mutex and its two methods:
- Lock
- Unlock