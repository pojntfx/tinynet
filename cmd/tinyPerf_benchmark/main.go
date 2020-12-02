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
	server := flag.Bool("s", false, "run in server mode")                                                                                      // Done
	client := flag.Bool("c", false, "run in client mode")                                                                                      // Done
	duration := flag.Int("t", 10, "time in seconds to transmit for")                                                                           // Done
	length := flag.Int("l", 128, "length of buffers to read or write (in KB)")                                                                 // Easy
	parallel := flag.Int("P", 1, "number of simultaneous connections to make to the server")                                                   // Make concurrent client requests

	flag.Parse()

	fmt.Println("port:", *port)
	fmt.Println("format:", *format)
	fmt.Println("interval:", *interval)
	fmt.Println("verbose:", *verbose)
	fmt.Println("server:", *server)
	fmt.Println("server:", *client)
	fmt.Println("time:", *duration)
	fmt.Println("length:", *length)
	fmt.Println("parallel:", *parallel)

	if *server {
		var wg sync.WaitGroup

		wg.Add(1)

		go tcpClient(port, &wg)
		tcpServer(port, &wg, duration)
	}

	if *client {
		var wg sync.WaitGroup
		var wgParallel sync.WaitGroup
		var wgFin sync.WaitGroup

		wg.Add(1)
		wgFin.Add(*parallel)

		go tcpServerClient(port, duration, &wg, &wgParallel)

		for i := 0; i < *parallel; i++ {
			go tcpClientClient(port, &wg, duration, &wgParallel, &wgFin)
			wgParallel.Add(1)
		}

		wgFin.Wait()

	}

}

func tcpServer(port *string, wg *sync.WaitGroup, duration *int) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	wg.Done()

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnection(conn, duration)
	}

}

func handleConnection(conn net.Conn, duration *int) {
	var input [512]byte
	var i int
	for start := time.Now(); time.Since(start) < time.Second*(time.Duration(*duration)); {
		i++
		startTimer := time.Now()
		_, err := conn.Write([]byte("Hello World!"))
		checkError(err)

		_, err = conn.Read(input[0:])
		checkError(err)
		elapsed := time.Since(startTimer)
		// Append time to array instead of printing it
		log.Printf("[INFO] Process time: %s", elapsed)
	}

}

func tcpClient(port *string, wg *sync.WaitGroup) {
	var input [512]byte

	wg.Wait()

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
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

func tcpServerClient(port *string, duration *int, wg *sync.WaitGroup, wgParallel *sync.WaitGroup) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	wg.Done()

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnectionClient(conn)
	}
}

func tcpClientClient(port *string, wg *sync.WaitGroup, duration *int, wgParallel *sync.WaitGroup, wgFin *sync.WaitGroup) {
	var input [512]byte
	var i int

	wg.Wait()

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	wgParallel.Done()
	wgParallel.Wait()
	for start := time.Now(); time.Since(start) < time.Second*(time.Duration(*duration)); {
		i++

		startTimer := time.Now()
		_, err = conn.Write([]byte("Hello World!"))
		checkError(err)

		_, err := conn.Read(input[0:])
		checkError(err)

		elapsed := time.Since(startTimer)
		// Append time to array instead of printing it
		log.Printf("[INFO] Process time: %s", elapsed)
	}

	wgFin.Done()
}

func handleConnectionClient(conn net.Conn) {

	var input [512]byte

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
