package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var (
	result  []int
	elapsed time.Duration
)

func main() {
	input := make([]byte, 10000)

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	for start := time.Now(); time.Since(start) < time.Second*time.Duration(10); {

		startTimer := time.Now()
		_, err = conn.Write(input)
		checkError(err)

		_, err := conn.Read(input[0:])
		checkError(err)

		elapsed = time.Since(startTimer)
		result = append(result, int(elapsed))

	}

	fmt.Println(fmt.Sprintf("result: %v", len(result)))

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
