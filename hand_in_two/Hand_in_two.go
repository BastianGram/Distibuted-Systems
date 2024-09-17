package main

import (
	"fmt"
)

//"net"

var client_chan chan string
var server_chan chan string
var client_ack, server_ack chan int
var client_syn, server_syn chan int
var ACK int
var seq int
var amountDoneEating int

func main() {

	//net.Listen("tcp", )

	//Listener.Accept()

	client_chan = make(chan string, 10)
	server_chan = make(chan string, 10)
	client_ack = make(chan int, 1)
	server_ack = make(chan int, 1)
	client_syn = make(chan int, 1)
	server_syn = make(chan int, 1)

	client_syn <- 0

	server_ack <- 0

	go start()

	for ACK < 10 {
		if ACK > 5 {
			fmt.Println("done!")
			return
		}

	}

}

func start() {

	var syn = <-client_syn
	ACK = <-server_ack

	server_syn <- syn
	fmt.Println("Syn sent from client")

	client_ack <- ACK
	client_syn <- syn
	fmt.Println("Syn and ack sent from server")

	seq = 1

	go sendFromClient()
	go sendFromServer()

}

func sendFromClient() {

	for {

		ACK = <-client_ack + 1
		server_ack <- ACK
		fmt.Println("ACK: ", ACK, ", Sent from client")

	}

}
func sendFromServer() {
	for {
		seq++
		ACK = <-server_ack
		client_ack <- ACK
		fmt.Println("ACK: ", ACK, ", sent from server")
	}

}
