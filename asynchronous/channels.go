package asynchronous

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/michael-steinert/GoEssentials/calculator"
)

func Channels() {
	messages := make(chan string)
	go func() {
		time.Sleep(time.Microsecond * 500)
		messages <- "ping"
		time.Sleep(time.Microsecond * 500)
		messages <- "pong"
	}()
	message := <-messages
	fmt.Println(message)
	message = <-messages
	fmt.Println(message)
}

func DoneChannels() {
	doneChannel := make(chan bool)
	go asyncFunc(doneChannel)
	// Waiting until a Value is written in (Unbuffered)Channel
	<-doneChannel
}

func asyncFunc(doneChannel chan<- bool) {
	fmt.Println("Start Processing")
	time.Sleep(time.Second * 2)
	fmt.Println("Finished Processing")
	// Writing Value in Directed Channel
	doneChannel <- true
}

func ManyChannels() {
	// Buffered Channel i.e. Channel with Size bigger then one
	primesCandidatesChannel := make(chan int, 100)
	doneChannel := make(chan bool)
	// Runs in Background Thread
	go func() {
		for {
			candidate, nextCandidate := <-primesCandidatesChannel
			if nextCandidate {
				if calculator.IsPrime(candidate) {
					fmt.Println(candidate)
				}
			} else {
				doneChannel <- true
				return
			}
		}
	}()
	// Runs in Main Thread
	for i := 999_999_000; i < 1_000_000_000; i++ {
		primesCandidatesChannel <- i
	}
	close(primesCandidatesChannel)
	// Writing to DoneChannel to signal that the Function is finished
	<-doneChannel
}

func ManyChannelsWaiting() {
	firstResultChannel := make(chan int)
	secondResultChannel := make(chan int)
	go getFirstResult(firstResultChannel)
	go getSecondResult(secondResultChannel)
	for i := 0; i < 2; i++ {
		// Select waits in parallel on many Channels until they are finished
		select {
		case firstResult := <-firstResultChannel:
			fmt.Printf("First Result %d\n", firstResult)
		case secondResult := <-secondResultChannel:
			fmt.Printf("Second Result %d\n", secondResult)
		}
	}
}

func getFirstResult(firstResultChannel chan<- int) {
	randomNumber := rand.Intn(5)
	milliseconds := time.Duration(randomNumber) * time.Second
	time.Sleep(milliseconds)
	firstResultChannel <- 1
}

func getSecondResult(secondResultChannel chan<- int) {
	randomNumber := rand.Intn(5)
	milliseconds := time.Duration(randomNumber) * time.Second
	time.Sleep(milliseconds)
	secondResultChannel <- 2
}

func ManyChannelsTimeout() {
	firstCandidate := 111111111
	secondCandidate := 111111113
	firstResultChannel := make(chan bool)
	secondResultChannel := make(chan bool)
	// Context is an Object that is shared between Functions
	ctx := context.Background()
	context, cancel := context.WithCancel(ctx)
	// A Cancel Signal can be sent to interrupt all Executions of the Function that use the Context
	defer cancel()
	go getResultWithContext(context, firstCandidate, firstResultChannel)
	go getResultWithContext(context, secondCandidate, secondResultChannel)
	// Select waits in parallel on many Channels until they are finished
	select {
	case firstResult := <-firstResultChannel:
		fmt.Println(firstCandidate, "is Prime:", firstResult)
	case secondResult := <-secondResultChannel:
		fmt.Println(secondCandidate, "is Prime:", secondResult)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
		cancel()
	}
}

func getResultWithContext(context context.Context, candidate int, resultChannel chan<- bool) {
	isNumberPrime := false
	select {
	case <-context.Done():
		fmt.Println("Cancelled")
		return
	default:
		isNumberPrime = calculator.IsPrime(candidate)
	}
	resultChannel <- isNumberPrime
}
