/*
	Channels

	When writing concurrent code (maybe parallel when using multiple cpu's),
	where multiple foroutines are in action and the program doesn't have
	a single running thread, then there´s a need of controlling how the data
	flows and communicate form goroutine to goroutine.

	We introduce channels to control how this data is shared between different goroutines.

	Go Proverb: "Don not communicate by sharing memory; instead, share memory by communicating."

*/
package main

import (
	"fmt"
)

func main() {

	// crate a channel
	fmt.Printf("[Create channel] ch1 := make(chan int)\n")
	ch1 := make(chan int)
	fmt.Printf("ch1:\t\t%v\n", ch1)
	fmt.Printf("ch1.Type:\t%T\n", ch1)

	// :::::::::::::::::::::::
	// :: BLOCKING CHANNELS ::
	// :::::::::::::::::::::::
	//
	// [!] Most important Key Tip: CHANNELS CAN BLOCK [1]
	// error: "fatal error: all goroutines are asleep - deadlock!"
	// meaning: send and receive may be able to happen at the **SAME** time
	// the following ch1 <- will create a deadlock, why?
	// 	ch1 <- 42
	// 1. we created the ch1
	// 2. program code execution follows the normal sequency
	// 3. it tries to send 42 over the channel
	// 4. send/receive CANNOT happen at the SAME time
	// 5. program returns fatal deadlock error
	//
	// Use goroutines to branch into a new thread of execution the send

	fmt.Printf("\n[Send value] go func(){ch1 <- 42}()\n")

	// By running the send to channel in a goroutine
	// we split the send and receive so they can happen
	// at the same time
	go func() {
		ch1 <- 42
	}()
	fmt.Printf("ch1:\t\t%v\n", <-ch1)

	//  :::::::::::::::::::::::
	//	:: Buffered Channels ::
	//  :::::::::::::::::::::::
	//
	//	A Buffered channels allow n values to be buffered
	//	so it will no longer block if the passed value
	//	fits into the buffer, even if the send/receive cannot
	//	happen at the same time.
	//
	fmt.Printf("\n[Create buffered channel] ch2 := make(chan int, 1)\n")
	ch2 := make(chan int, 1)
	fmt.Printf("ch2:\t\t%v\n", ch2)
	fmt.Printf("ch2.Type:\t%T\n", ch2)

	fmt.Printf("\n[Send value directly to buffer] ch2 <- 42\n")
	ch2 <- 42

	// Buffer is full now
	// A new value being sent would mean a deadlock

	fmt.Printf("ch1:\t\t%v\n", <-ch2)

	// Now it's safe to send a new value, since the first one
	// is already outside the channel, as we output if in the
	// above Printf
	// In general, using buffered channels is a **bad idea**.

	//
	// ::::::::::::::::::::::::::
	// :: Directional Channels ::
	// ::::::::::::::::::::::::::
	//
	// By default channels can send and receive data, but:
	// - Channels can also be created send-only and receive-only.
	// - Send/receive channels can be passed and returned as send/receive only to/from funcs.
	//
	// Directions:
	// <-chan is a receive-only channel; data can't be send into it, only received from it
	// chan<- is a send-only channel; data can't be retrieved from it, only sent to it.
	//

	fmt.Printf("\n[Create receive-only channel] ch3 := make(<-chan int)\n")
	ch3 := make(<-chan int)
	fmt.Printf("ch3:\t\t%v\n", ch3)
	fmt.Printf("ch3.Type:\t%T\n", ch3)

	/*
		The following code snippet would fail with:
		invalid operation: cannot send to receive-only channel ch3 (variable of type <-chan int)

		go func() {
			ch3 <- 42
		}()
	*/

	fmt.Printf("\n[Key point!] A bi-directional channel can become more specific\n")
	fmt.Printf("ch4 := (<-chan int)(ch1)\n")
	fmt.Printf("ch5 := (chan<- int)(ch1)\n")
	ch4 := (<-chan int)(ch1)
	ch5 := (chan<- int)(ch1)
	fmt.Printf("ch4.Type:\t%T\n", ch4)
	fmt.Printf("ch5.Type:\t%T\n", ch5)

	/*
		chSender will send an int to the channel and the channel will be send-only
		chReceiver will receive the int from the channel
	*/

	fmt.Printf("\n[Sending/Receiving data from channels] Spawm two func: chSender / chReceiver\n")
	fmt.Printf("signature: chSender(c chan<- int, n int)\n")
	fmt.Printf("signature: chReceiver(c <-chan int) int\n")

	// chSender is sent to its own goroutine
	// remember that in order to not block the channel
	// the program has to be able to send/receive at the same time
	fmt.Printf("\nsend data in a goroutine:\n\tgo chSender(ch1, 10)\n")
	go chSender(ch1, 10)

	// chReceiver is invoked in the normal control sequency
	// in this case we don't need waitgroups, the func will simply wait
	// for the int to be sent to the channel and receive it
	// no more send/receive controls needed
	fmt.Printf("receive data in the normal program sequence:\n\tnum := chReceiver(ch1)\n")
	num := chReceiver(ch1)
	fmt.Printf("received num:\n\t%v\n", num)

	fmt.Printf("\n[Sending/Receiving data in Loops] Spawm two func: chSenderN / chReceiverN\n")
	fmt.Printf("signature: chSenderN(c chan<- int, n int)\n")
	fmt.Printf("signature: chReceiverN(c <-chan int) int\n")

	// chSenderN is sent to its own goroutine
	// remember that in order to not block the channel
	// the program has to be able to send/receive at the same time
	//
	// chSenderN sends data to the chan from a for loop:
	// [!] In order to receive chan from chReceiverN we have to:
	//	- Before exiting chSenderN, we want to close the chan with close(c)
	//	- Without closing it, we risk a deadlock!
	//	- Because the goroutine would still be processing data, and the main func wouldn't be capable of retrieving it.
	//	- REMEMBER: Send/Receive has to happen at the SAME time!
	fmt.Printf("\nsend data in a goroutine:\n\tgo chSenderN(ch1, 10)\n")
	go chSenderN(ch1, 10)

	// chReceiverN is invoked in the normal control sequency
	// in this case we don't need waitgroups, the func will simply wait
	// for the int to be sent to the channel and receive it
	// no more send/receive controls needed
	fmt.Printf("receive data in the normal program sequence:\n\tnum := chReceiverN(ch1)\n")
	sum := chReceiverN(ch1)
	fmt.Printf("received sum:\n\t%v\n", sum)

	// Making decisions based on data received over a channel: select
	// For this we'll create three channels:
	// - First channel will send even integers
	// - Second channel will send odd integers
	// - Third channel will send a signal to shutdown
	fmt.Printf("\n[Sending/Receiving data from channels] Select statement\n")
	fmt.Printf("signature sndr: channelRcvSignal(e, o <-chan int, s <-chan bool)\n")
	fmt.Printf("signature recv: channelRcvSignal(e, o <-chan int, s <-chan bool)\n")
	evenCh := make(chan int)
	oddCh := make(chan int)
	signalCh := make(chan bool)
	defer close(evenCh)
	defer close(oddCh)
	defer close(signalCh)

	// Sender function
	// send only channels, narrow scope
	go channelSndrSignal(evenCh, oddCh, signalCh)

	// Receiver function
	// recv only channels, narrow scope
	channelRcvSignal(evenCh, oddCh, signalCh)

	// comma ok idiom is available with channels
	// this is useful to control when to close channels
	fmt.Printf("\ncomma-ok idiom is available in channels: value, ok := <- c\n")
	ch2 <- 42
	value, ok := <-ch2
	fmt.Printf("[Received data] comma-ok reports value: %v and ok: %v\n", value, ok)
	value, ok = <-ch1
	fmt.Printf("[No new data] comma-ok reports value: %v and ok: %v\n", value, ok)
}

// chSender sends integer n into channel c
func chSender(c chan<- int, n int) {
	c <- n
}

// chSender sends n integers into channel c
func chSenderN(c chan<- int, n int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}

// chReceiver receives data from channel c
func chReceiver(c <-chan int) int {
	return <-c
}

// chReceiver receives n integers from channel c
func chReceiverN(c <-chan int) int {
	sum := 0
	for value := range c {
		sum += value
	}
	return sum
}

// channelSndrSignal sends even and odd numbers over a channel, and close it at the end
func channelSndrSignal(e, o chan<- int, s chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// send a signal to shutdown
	s <- false
}

// channelRcvSignal receives even/odd numbers over channels, and signal to shutdown the func
func channelRcvSignal(e, o <-chan int, s <-chan bool) {
	// infinite for
	// we'll wait the signaling to shutdown
	for {
		// select syntax is similar to switch
		select {
		case i := <-e:
			fmt.Println("evenCh sends:\t", i)
		case i := <-o:
			fmt.Println("oddCh sends:\t", i)
		case s := <-s:
			fmt.Println("signalCh sends:\t", s)
			fmt.Println("Shutting down.")
			return
		}
	}
}
