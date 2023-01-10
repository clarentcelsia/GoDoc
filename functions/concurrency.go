package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"simple-go/model"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

/**
GOROUTINES AND CHANNEL

Goroutines is a lightweight thread.
A function marked/started by "go" will run independently.
Are functions which executes """concurrently along with""" the main program flow.

Goroutine interacts with other goroutines using a communication mechanism called Channels.

Channel is a pipeline for sending & receiving data.
Channel provides a goroutine to send a data to another goroutine as receiver,
	which means [One goroutine writes the data into the channel and the other goroutine can read from it]

Buffered Channel is a channel with capacity.
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

func BasicChannels() {
	//var c chan int >>> has to be initialized or
	c := make(chan int) // >>> Go allocates memory for c (recommended) or
	//c := make(chan int, 1) >>> channel with buffer cap

	c <- 5       // send(write) value to c
	res := <-c   // retrieve value from c
	println(res) // >>> 5

	// note. channel that left open will be collected by garbage collector if it's no longer used.
	close(c)
}

// By default, channel is bidirectional (r/w).
// The unidirectional channel is used to provide the type-safety of the program
// so, that the program gives less error.
func UnidirectionalChannel() {
	// read-only channel/receiving
	// r := make(<- chan int)

	// write-only channel/sending
	// w := make(chan <- int)

	// example==========
	// Params 's' in sendfunc will be written.
	sendfunc := func(s chan<- string) {
		s <- "This s is written by x"
	}

	// This bidirectional channel...
	write_read_me := make(chan string)
	// will be converted to write-only channel '''inside goroutine'''...
	go sendfunc(write_read_me)
	// '''outside goroutine''' will be back to bidirectional channel
	fmt.Println(<-write_read_me)

}

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

func ChannelBlockingWork() {
	s1 := make(chan string)
	s2 := make(chan string)

	go func() {
		s1 <- "1 has been received data"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		s2 <- "2 has been received data"
	}()

	for i := 0; i < 2; i++ {
		/* * as go documentation about select:*/
		select {
		// [ If one or more of the communications can proceed,
		// 		a single one that can proceed is chosen via a uniform pseudo-random selection.]

		// technically, this is likely to be in race condition
		// data that will be chosen/printed first is the data that has been finished first.

		case data1 := <-s1:
			fmt.Println(data1)

		case data2 := <-s2:
			fmt.Println(data2)

			// Otherwise, '''if there is a default case, that case is chosen'''.

			// If there is no default case, the "select" statement
			// '''blocks until at least one of the communications can proceed'''.

			// default:
			// 	fmt.Println("default")
		}
	}

	// Channel will block this line until all (channel) data is consumed by another goroutine.
	fmt.Println("FINISHED")
}

func BasicGoroutines() {
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

// *model.Job refers to the same memory addr
// context tells us about this method is about to be compiled
// with this ctx, we able to tell golang when to stop getting result from
// network call.
func Goroutines(*gin.Context) <-chan *model.Job {
	res := make(chan *model.Job, 1)

	go func() {
		resp, err := http.Get("http://localhost:9000/pagination")
		if err != nil {
			res <- &model.Job{
				Responses: nil,
				Errors:    fmt.Errorf("unable to fetch data, %s", err.Error()),
			}
			return
		}

		defer resp.Body.Close()

		bytresp, _ := ioutil.ReadAll(resp.Body)
		var respmap interface{}
		json.Unmarshal(bytresp, &respmap)

		res <- &model.Job{
			Responses: respmap,
			Errors:    nil,
		}

	}()

	return res

}

// To wait for multiple goroutines to finish, we can use a wait group.
func WaitGroup() {
	var w sync.WaitGroup

	doJob := func(i int) {
		defer w.Done() // tell if the job has finished
		fmt.Printf("Worker %d is doing a job \n", i)
	}

	var workers = 10
	for i := 0; i < workers; i++ {
		fmt.Println("A new job is coming!")
		w.Add(1) // tell group there's a new job doing by worker now
		doJob(i + 1)
	}

	w.Wait() // wait till all jobs has finished by worker then go to the next command
	fmt.Println("All task has finished!")
}

// This occurs when 2 or more goroutines have shared data and interact it simultaneously.
// By example, goroutine funct A and go func B need the same resource at the same time
// in that case, both (go) func in a Race Condition in which while the app is running
// both function will do their job concurrently,
// while funct A is doing read from resource X it's the same as the funct B
// that causes data/resource they use will be updated once where it's supposed to be (updated) twice
// cause both func initially read the same value of the resource.

// note.
// basically, go tries to avoid this race condition on its own
// and the possibility of its success when the workers work behind relatively small.
// read. cooperative multithreading.

func RaceCondition() {
	var (
		workers = 1000

		w   sync.WaitGroup
		res model.Resource
	)

	w.Add(workers) // call n workers in their separate goroutines
	for i := 0; i < workers; i++ {
		go func() {
			runtime.Gosched()

			res.Add() // n workers doing the same job at the same time
			w.Done()
		}()
	}

	w.Wait()

	fmt.Println(res.Sum())
}

// to handle race condition
func Mutex() {
	// var mtx sync.Mutex
	var w sync.WaitGroup

	var totalJobs = 1000
	var additionalJobReq = 100

	// mutex already defined in struct
	var req model.Resource

	for i := 0; i < totalJobs; i++ {
		w.Add(1)
		go func() {
			for j := 0; j < additionalJobReq; j++ {
				//Lock for other goroutines access the shared variable
				// mtx.Lock()
				req.Add()
				// mtx.Unlock()
			}
			w.Done()
		}()
	}

	w.Wait()

	fmt.Println(req.Sum())
}
