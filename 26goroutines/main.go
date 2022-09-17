package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
Concurrency vs Parallelism

Let's consider an situation where you are eating,you got a notification on phone,also you are feeling hot so need to turn on the AC.

Concurrency :
Put down your plate.
Look at the phone.
Put down the phone.
Pickup the AC remote turn it On.Put down the remote.
Now again pickup the plate and start eating.

Parallelism :
Checking phone while eating then also taking the AC remote to turn it on

In Concurrency you do one thing a time whereas in parallelism you are multitasking.

Go routines is the way how you achieve parallelism.

Diff b/w threads and Go Routines
Threads :
     Managed by OS.
	 fixed stack - 1MB

Goroutines :
	Managed by Go Runtime
	fixed stack - 2KB

	Goroutines :> Do not communicate by shared memory; instead,share memory by communicating.

*/

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    // pointer

func main() {
	// greeter("Hello")
	// greeter("World")
	/*
		The above two first prints the Hello word 5 times then World word 5 times.
	*/
	// go greeter("Hello")
	// greeter("World")
	// this prints world only,it's like we engrossed in phone so much that we forget to eat food or turn on the AC

	websiteList := []string{
		"https://github.com",
		"https://go.dev",
		"https://google.com",
		"https://twitter.com",
		"https://netflix.com",
		"https://amazon.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond) // to avoid getting world only printed
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}
