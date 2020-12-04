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

	port := flag.String("p", "8888", "port to connect/listen to")
	interval := flag.Int("i", 1, "report intervals in seconds")
	server := flag.Bool("s", false, "run as server")
	client := flag.Bool("c", false, "run as client")
	duration := flag.Int("t", 10, "time to test in s")
	length := flag.Int("l", 128, "size of the buffer to transfer in Kb")
	ip := flag.String("ip", "0.0.0.0", "ip to connect to")

	flag.Parse()

	*length = *length * 1000

	if *server {
		handleServerMode(ip, port, length, interval, duration)
	}

	if *client {
		handleClientMode(length, ip, port, interval)
	}

	sum := 0

	for i := 0; i < len(result); i++ {
		sum += result[i]
	}

	if *server {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("Server mode")
		fmt.Println(fmt.Sprintf("Packets of length %v Kb have been received for %v s", *length/1000, *duration))
		fmt.Println(fmt.Sprintf("Number of requests: %v", len(result)))
		fmt.Println(fmt.Sprintf("Average transfer speed: %v Mb/s", float64(len(result)*(*length))/float64(10*1000*1000)))
		fmt.Println("-----------------------------------------------------")
	}

	if *client {
		fmt.Println("-----------------------------------------------------")
		fmt.Println(fmt.Sprintf("Connection to: %v:%v", *ip, *port))
		fmt.Println(fmt.Sprintf("Packets of length %v Kb have been sent for %v s", *length/1000, *duration))
		fmt.Println(fmt.Sprintf("Number of requests: %v", len(result)))
		fmt.Println(fmt.Sprintf("Average transfer speed: %v Mb/s", float64(len(result)*(*length))/float64(10*1000*1000)))
		fmt.Println("-----------------------------------------------------")
	}
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

		result = append(result, int(time.Since(startTimer)))

		elapsed = time.Since(startTimer)

	}

	wg.Done()
}

func handleClientMode(length *int, ip *string, port *string, interval *int) {
	input := make([]byte, *length)

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", *ip, *port))
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	go doEvery(time.Duration(*interval) * time.Second)
	for start := time.Now(); time.Since(start) < time.Second*time.Duration(10); {

		startTimer := time.Now()

		_, err = conn.Write(input)
		checkError(err)

		_, err := conn.Read(input[0:])
		checkError(err)

		result = append(result, int(time.Since(startTimer)))

		elapsed = time.Since(startTimer)

	}
}

func doEvery(d time.Duration) {
	for x := range time.Tick(d) {
		fmt.Println(fmt.Sprintf("Current response time: %v", elapsed))
		_ = x
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
