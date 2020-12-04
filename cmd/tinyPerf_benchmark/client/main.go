package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {

	port := flag.String("p", "8888", "port to connect/listen to")
	length := flag.Int("l", 128, "length of the buffers to send in Kb")
	ip := flag.String("ip", "0.0.0.0", "ip to connect to")

	*length = *length * 1000

	input := make([]byte, *length)

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", *ip, *port))
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
