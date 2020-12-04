package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	input := make([]byte, 100000)

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:8888"))
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	for {
		n, err := conn.Read(input[0:])
		checkError(err)

		_, err = conn.Write(input[0:n])
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
