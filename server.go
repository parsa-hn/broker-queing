package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Wrong input format")
		os.Exit(1)
	}

	brokerAddress := os.Args[1]
	serverType := os.Args[2]
	var input string

	broker, err := net.ResolveTCPAddr("tcp4", brokerAddress)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, broker)
	checkError(err)

	for {

		fmt.Scanln(&input)
		if input == "exit" {
			break
		}

		if serverType == "sync" {
			sendMessage(conn)
		}
		if serverType == "async" {
			go sendMessage(conn)
		}
	}
	os.Exit(0)
}

func sendMessage(conn net.Conn) {
	input := make([]byte, 512)
	ack := make([]byte, 512)

	fmt.Scanln(&input)

	_, err := conn.Write(input)
	if err != nil {
		fmt.Println("Connection error during sending message")
		return
	}

	_, err = conn.Read(ack)
	if (err != nil) {
		fmt.Println("Acknowledge not recived for this message")
		return
	}

	fmt.Println("Ack recieved for message", input)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
