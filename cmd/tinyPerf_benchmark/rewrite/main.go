package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	elapsed time.Duration
	result  []int
)

func main() {

	port := flag.String("p", "8888", "port to listen to")                                             // Done
	interval := flag.Int("i", 1, "set interval between periodic bandwidth, jitter, ans loss reports") // Done
	verbose := flag.Bool("V", false, "give more detailed output")                                     // Easy 4)
	server := flag.Bool("s", false, "run in server mode")                                             // Done
	client := flag.Bool("c", false, "run in client mode")
	duration := flag.Int("t", 10, "time in seconds to transmit for")           // Done
	length := flag.Int("l", 128, "length of buffers to read or write (in KB)") // Done
	ip := flag.String("ip", "0.0.0.0", "ip to connect to")

	flag.Parse()

	*length = *length * 1000

	fmt.Println("port:", *port)
	fmt.Println("interval:", *interval)
	fmt.Println("verbose:", *verbose)
	fmt.Println("server:", *server)
	fmt.Println("client:", *client)
	fmt.Println("duration:", *duration)
	fmt.Println("length:", *length)
	fmt.Println("ip:", *ip)

	if *server {
		handleServerMode(ip, port, length, interval, duration)
	}

	if *client {
		handleClientMode(length, ip, port)
	}

	fmt.Println(len(result))
}

func handleServerMode(ip *string, port *string, length *int, interval *int, duration *int) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", *ip, *port))
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	var wg sync.WaitGroup
	wg.Add(1)

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnection(conn, &wg, length, interval, duration)

		wg.Wait()

		break
	}
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup, length *int, interval *int, duration *int) {
	input := make([]byte, *length)

	go doEvery(time.Duration(*interval) * time.Second)
	for start := time.Now(); time.Since(start) < time.Second*(time.Duration(*duration)); {

		startTimer := time.Now()

		_, err := conn.Write(input)
		checkError(err)

		_, err = conn.Read(input[0:])
		checkError(err)

		elapsed = time.Since(startTimer)

		result = append(result, int(elapsed))
	}

	wg.Done()
}

func handleClientMode(length *int, ip *string, port *string) {
	input := make([]byte, *length)

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", *ip, *port))
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
}

func doEvery(d time.Duration) {
	for x := range time.Tick(d) {
		fmt.Println(elapsed)
		_ = x
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
