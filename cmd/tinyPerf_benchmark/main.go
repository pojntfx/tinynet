package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {

	port := flag.String("p", "8888", "port to listen to")
	format := flag.String("f", "M", "specify the format of bandwidth numbers. (k = Kbits/sec, K = KBytes/sec, m = Mbits/sec, M = MBytes/sec)")
	interval := flag.Int("i", 0, "set interval between periodic bandwidth, jitter, ans loss reports")
	verbose := flag.Bool("V", false, "give more detailed output")
	server := flag.Bool("s", false, "run in server mode")
	times := flag.Int("t", 10, "time in seconds to transmit for")
	length := flag.Int("l", 128, "length of buffers to read or write (in KB)")
	parallel := flag.Int("P", 1, "number of simultaneous connections to make to the server")
	reverse := flag.Bool("R", false, "run in reverse mode (server sends, client receives)")

	flag.Parse()

	fmt.Println("port:", *port)
	fmt.Println("format:", *format)
	fmt.Println("interval:", *interval)
	fmt.Println("verbose:", *verbose)
	fmt.Println("server:", *server)
	fmt.Println("time:", *times)
	fmt.Println("length:", *length)
	fmt.Println("parallel:", *parallel)
	fmt.Println("reverse:", *reverse)

	var wg sync.WaitGroup

	wg.Add(1)

	go tcpServer(port, &wg)
	tcpClient(port, &wg)
}

func tcpServer(port *string, wg *sync.WaitGroup) {

	fmt.Println(*port)

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

	_, err := conn.Write([]byte("Hello World!"))
	checkError(err)
	fmt.Println("handling Connection...")
}

func tcpClient(port *string, wg *sync.WaitGroup) {
	var input [512]byte

	wg.Wait()

	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	checkError(err)
	fmt.Println("TCP address resolved")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	fmt.Println("TCP address dialed")

	n, err := conn.Read(input[0:])
	checkError(err)
	fmt.Println("Read from server...")

	fmt.Println(string(input[0:n]))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
