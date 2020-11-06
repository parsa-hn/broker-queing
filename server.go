package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("wrong input format")
		os.Exit(1)
	}
	brokerAddres := os.Args[1]
	serverType := os.Args[2]

	var input string

	broker, err := net.ResolveTCPAddr("tcp4", brokerAddres)
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	for i := 1; ; i++ {

		fmt.Scanln(&input)
		if input == "exit" {
			break
		}

		if serverType == "sync" {
			sendMessage(i, broker)
		}
		if serverType == "async" {
			go sendMessage(i, broker)
		}
	}
	os.Exit(0)
}

func sendMessage(messageNumber int, broker *net.TCPAddr) {
	ack := make([]byte, 32)
	conn, err := net.DialTCP("tcp", nil, broker)
	if err != nil {
		fmt.Println("connection error when sending message number %d", messageNumber)
		return
	}

	_, err = conn.Write([]byte("message"))
	if err != nil {
		fmt.Println("connection error when sending message number %d", messageNumber)
		return
	}

	_, err = conn.Read(ack)
	if err != nil {
		fmt.Println("connection error when sending message number %d", messageNumber)
		return
	}

	fmt.Println("ack for '%s' recieved for message numbner %d", string(ack), messageNumber)
}
