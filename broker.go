package main

import (
	"fmt"
	"net"
	"os"
)

type Message struct {
	Content string
}

var channel chan string = make(chan string, 10)
var serverCon net.Conn

func runBroker(conn net.Conn) {
	message := <-channel
	conn.Write([]byte(message))
	serverCon.Write([]byte("ack"))
	conn.Close()
	fmt.Println("Message delivered")
}

func connectToClient() {
	fmt.Println("Broker is listening for new clients...")
	service := ":8001"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err == nil {
			fmt.Println("New Client")
			go runBroker(conn)
		}
	}
}

func connectToServer() net.Conn {
	fmt.Println("Broker is listening for a server...")
	service := ":8000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err == nil {
			return conn
		}
	}
}

func readMessage() {
	for {
		message := make([]byte, 512)
		serverCon.Read(message)
		channel <- string(message)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	serverCon = connectToServer()

	go readMessage()

	connectToClient()

	close(channel)
}
