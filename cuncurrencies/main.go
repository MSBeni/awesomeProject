package main

import (
	"fmt"
	"time"
)

//In Go, each concurrently executing activity is called a goroutine. Consider a program that has
//two functions, one that does some computation and one that writes some output, and assume
//that neither function calls the other. A sequential program may call one function and then call the other, but in a
// concurrent program with two or more goroutines, calls to both functions can be active at the same time. goroutine is
// similar to a thread
// If goroutines are the activities of a con cur rent Go program, channels are the connections between them. A channel
// is a communication mechanism that lets one goroutine send values to another goroutine. Each channel is a conduit for
// values of a particular type, called the channel’s element type.
func main(){
	sq := make(chan bool)
	qs := make(chan bool)
	nqs := make(chan int)
	/// defining a channel periodically listen to a value will be sent over
	go Periodiclisten(nqs)
	for i:= range nqs{
		fmt.Println(i)
	}
	_,ok := <- nqs    /// check whether the channel is open/close
	fmt.Println(ok)

	go func() {
		fmt.Println("saying hello from the anonimous func")
		qs <- true
	}()
	go SayingHello(sq)


	fmt.Println("Hello world")
	//<- sq
	<- qs
	val := <- sq   // channel value and the opening in the main
	fmt.Println(val)
}

// making a function with an channel input to enable the code to run simultaneously
func SayingHello(sq chan bool){
	fmt.Println("Hello from the goroutine")
	sq <- true      // enable the code to go through the channel
}
func Periodiclisten(nqs chan int){

	for i:=0;i<3;{
		nqs <- i
		time.Sleep(1*time.Second)
		i++
	}
	close(nqs)

}