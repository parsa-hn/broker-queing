package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	brokerAddres := os.Args[1]

	var sleep_delay int
	task := make([]byte, 512)

	broker, err := net.ResolveTCPAddr("tcp4", brokerAddres)
	checkError(err)
	for true {
		conn, err := net.DialTCP("tcp", nil, broker)
		checkError(err)

		_, err = conn.Write([]byte("woker ready"))
		checkError(err)

		_, err = conn.Read(task)
		if err != nil {
			continue
		}

		if string(task) == "close" {
			fmt.Println("cleint closed")
			break
		}

		sleep_delay = rand.Intn(10) + 1
		fmt.Println("message '%s' recieved and now sleeping for %d seconds", string(task), sleep_delay)
		time.Sleep(time.Duration(sleep_delay) * time.Second)
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
