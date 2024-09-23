package main

import (
	"fmt"
	"time"
)

const (
	SYN2     = "SYN"
	SYN_ACK2 = "SYN-ACK"
	ACK2     = "ACK"
)

type Message2 struct {
	Type string
}

func main() {
	ch := make(chan Message2)

	go client2(ch)
	go server2(ch)

	// Allow time for goroutines to finish
	time.Sleep(2 * time.Second)
}

func client2(ch chan Message2) {
	fmt.Println("Client: Sending SYN=100")
	ch <- Message2{Type: SYN2}

	msg := <-ch
	if msg.Type == SYN_ACK2 {
		fmt.Println("Client: Recieved SYN=101 ACK=300")

	}

	ch <- Message2{Type: ACK2}
	fmt.Println("Client: Sending ACK=301")

}

func server2(ch chan Message2) {
	msg := <-ch
	if msg.Type == SYN2 {
		fmt.Println("Server: Recieved SYN=100")
	}

	fmt.Println("Server: Sending SYN=101 ACK=300")
	ch <- Message2{Type: SYN_ACK2}

	msg2 := <-ch
	if msg2.Type == ACK2 {
		fmt.Println("Server: Recieved ACK=301. Connection Established!")
	}

}
