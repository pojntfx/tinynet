package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func main() {

	port := flag.String("p", "8888", "port to listen to")                                                                                      // Done
	format := flag.String("f", "M", "specify the format of bandwidth numbers. (k = Kbits/sec, K = KBytes/sec, m = Mbits/sec, M = MBytes/sec)") // Easy
	interval := flag.Int("i", 0, "set interval between periodic bandwidth, jitter, ans loss reports")                                          // Easy
	verbose := flag.Bool("V", false, "give more detailed output")                                                                              // Easy
	server := flag.Bool("s", false, "run in server mode")                                                                                      // we are server and client echoes
	client := flag.Bool("c", false, "run in client mode")                                                                                      // we are client and server echoes
	time := flag.Int("t", 10, "time in seconds to transmit for")                                                                               // infinite for loop which sends to echo server for TIME seconds, each iteration have timer and print timer output and store result in array
	length := flag.Int("l", 128, "length of buffers to read or write (in KB)")                                                                 // Easy
	parallel := flag.Int("P", 1, "number of simultaneous connections to make to the server")                                                   // * needs to be done in a different way, we need concurrent clients
	reverse := flag.Bool("R", false, "run in reverse mode (server sends, client receives)")                                                    // * needs to be done in a different way

	flag.Parse()

	fmt.Println("port:", *port)
	fmt.Println("format:", *format)
	fmt.Println("interval:", *interval)
	fmt.Println("verbose:", *verbose)
	fmt.Println("server:", *server)
	fmt.Println("server:", *client)
	fmt.Println("time:", *time)
	fmt.Println("length:", *length)
	fmt.Println("parallel:", *parallel)
	fmt.Println("reverse:", *reverse)

	// server mode
	var wg sync.WaitGroup

	wg.Add(1)

	go tcpClient(port, &wg)
	tcpServer(port, &wg)
}

func tcpServer(port *string, wg *sync.WaitGroup) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	wg.Done()

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	var input [512]byte

	start := time.Now()
	_, err := conn.Write([]byte("Hello World!"))
	checkError(err)

	n, err := conn.Read(input[0:])
	checkError(err)
	elapsed := time.Since(start)
	log.Printf("[INFO] Process time: %s", elapsed)

	fmt.Println(string(input[0:n]))
}

func tcpClient(port *string, wg *sync.WaitGroup) {
	var input [512]byte

	wg.Wait()

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	n, err := conn.Read(input[0:])
	checkError(err)

	_, err = conn.Write(input[0:n])
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
