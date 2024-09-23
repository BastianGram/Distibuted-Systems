package main

import(
	"fmt"
	"time"
)

const (
	SYN = "SYN"
	SYN_ACK = "SYN-ACK"
	ACK = "ACK"
)

type Message struct{
	message string
}

func main(){
	ch1 := make(chan Message,2)
	ch2 := make(chan Message,2)

    go client(ch1, ch2)
    go server(ch1, ch2)

    // Allow time for goroutines to finish
    time.Sleep(1 * time.Second)
}

func client(ch1 chan Message, ch2 chan Message) {
	ch1 <- Message{SYN}

	var recieve Message
	for {
		recieve = <-ch2
		if (recieve == Message{SYN_ACK}) {
			fmt.Println("Connection established requesting another segment")
			ch1 <- Message{SYN_ACK}
			break
		} else {
			fmt.Println("Connection failed. Sending anouther SYN")
			ch1 <- Message{SYN}
		}
	}

}

func server(ch1 chan Message, ch2 chan Message) {
	var recieve Message
	for {
		recieve = <-ch1
		if (recieve == Message{SYN}) {
			ch2 <- Message{SYN_ACK}
		} else if (recieve == Message{SYN_ACK}) {
			fmt.Println("Success")
			break
		} else {
			ch2 <- Message{"fail"}
		}
	}
}

	
