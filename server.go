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
	ack := make([]byte, 32)

	_, err := conn.Write([]byte("message"))
	if err != nil {
		fmt.Println("connection error during sending message")
		return
	}

	_, err = conn.Read(ack)
	if err != nil {
		fmt.Println("connection error during sending message")
		return
	}

	fmt.Println("ack recieved for message numbner %d", string(ack))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
