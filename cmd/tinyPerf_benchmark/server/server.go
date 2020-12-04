package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {

	port := flag.String("p", "8888", "port to connect/listen to")
	length := flag.Int("l", 128, "length of the buffer to transfer")
	ip := flag.String("ip", "0.0.0.0", "ip to connect to")

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", *ip, *port))
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnection(conn, length)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleConnection(conn net.Conn, length *int) {
	input := make([]byte, *length)

	for {
		n, err := conn.Read(input[0:])
		checkError(err)

		_, err = conn.Write(input[0:n])
		checkError(err)
	}

}
