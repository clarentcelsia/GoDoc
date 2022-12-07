package functions

import "fmt"

/**
GOROUTINES AND CHANNEL

Goroutines is a lightweight thread.
A function marked/started by "go" will run independently.

Channel is a pipeline for sending & receiving data.
Channel provides a goroutine to send a data to another goroutine as receiver.
*/

// CONCURRENCY & PARALLELISM
// https://www.golangprograms.com/go-language/concurrency.html
/**
Concurrency is about to handle numerous tasks at once but
only be doing a single tasks at a time.

Parallelism is about doing lots of tasks at once.
This means that even if we have two tasks,
they are continuously working without any breaks in between them
*/

func Channels() {
	// In buffered channel there is a capacity to hold one or more values
	// before they're received. In this types of channels don't force goroutines
	// to be ready at the same instant to perform sends and receives.
	mug := make(chan string, 2)

	// channel 'mug' is able to hold 2 types of drink. [goroutine receive str data]
	println("ADD MUG 1 MILK")
	mug <- "milk"
	println("ADD MUG 2 TEA")
	mug <- "tea"

	// if I input one drink to mug, error will be occured.
	// mug <- "coffee" >>> error

	// [goroutine send the data to Printfunc]
	fmt.Println(<-mug) // >>> milk
	fmt.Println(<-mug) // >>> tea

	//---------released---------

	// Channel mug doesnt hold anything, then below code is valid
	mug <- "coffee"
	go fmt.Println(<-mug) // >>> coffee (see Goroutines(). [how to fetch goroutines value])

	mug <- "fanta"
	fmt.Println(<-mug) // >>> fanta
}

func Goroutines() {
	meatball := func(num int, temp chan int) {
		temp <- num
	}

	bowl := make(chan int, 1)
	// illustrate.
	// channel Bowl holds max 1 meatball
	meatball(1, bowl) // 1 meatball has been received, goroutines ready to execute

	//meatball(5, bowl) // add another meatball, error occurs because the bowl has been reached max meatball num.

	go meatball(2, bowl) // runs independently,
	go meatball(3, bowl) // runs independently

	fmt.Printf("%d", <-bowl) // >>> 1
	fmt.Printf("%d", <-bowl) // >>> 2
	fmt.Printf("%d", <-bowl) // >>> 3

	close(bowl)
}
