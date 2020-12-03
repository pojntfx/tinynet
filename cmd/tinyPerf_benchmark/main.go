package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"reflect"
	"sync"
	"time"
)

var (
	elapsed  time.Duration
	interval *int
	result   []int
)

func main() {

	// reponse time
	// datenrate bytes / avg response time
	port := flag.String("p", "8888", "port to listen to")                                            // Done
	interval = flag.Int("i", 1, "set interval between periodic bandwidth, jitter, ans loss reports") // Done
	verbose := flag.Bool("V", false, "give more detailed output")                                    // Easy 4)
	server := flag.Bool("s", false, "run in server mode")                                            // Done
	client := flag.Bool("c", false, "run in client mode")
	reverse := flag.Bool("r", false, "run in reverse mode")                                  // Done
	duration := flag.Int("t", 10, "time in seconds to transmit for")                         // Done
	length := flag.Int("l", 128, "length of buffers to read or write (in KB)")               // Done
	parallel := flag.Int("P", 1, "number of simultaneous connections to make to the server") // Done

	flag.Parse()

	*length = *length * 1000

	fmt.Println("port:", *port)
	fmt.Println("interval:", *interval)
	fmt.Println("verbose:", *verbose)
	fmt.Println("server:", *server)
	fmt.Println("client: ", *client)
	fmt.Println("server:", *reverse)
	fmt.Println("time:", *duration)
	fmt.Println("length:", *length)
	fmt.Println("parallel:", *parallel)

	if *server {
		var wg sync.WaitGroup

		wg.Add(1)

		go tcpClient(port, &wg, length)
		tcpServer(port, &wg, duration, length)
	}

	if *reverse {
		var wg sync.WaitGroup
		var wgParallel sync.WaitGroup
		var wgFin sync.WaitGroup

		wg.Add(1)
		wgFin.Add(*parallel)

		go tcpServerClient(port, &wg, &wgParallel, length)

		for i := 0; i < *parallel; i++ {
			go tcpClientClient(port, &wg, duration, &wgParallel, &wgFin, length)
			wgParallel.Add(1)
		}

		wgFin.Wait()
		fmt.Println(fmt.Sprintf("result: %v", len(result)))

	}

	if *client {
		var wgParallel sync.WaitGroup
		var wgFin sync.WaitGroup

		wgFin.Add(*parallel)

		for i := 0; i < *parallel; i++ {
			go tcpClientClientFlag(port, duration, &wgParallel, &wgFin, length)
			wgParallel.Add(1)
		}

		wgFin.Wait()
		fmt.Println(fmt.Sprintf("result: %v", len(result)))
	}

}

func tcpServer(port *string, wg *sync.WaitGroup, duration *int, length *int) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	wg.Done()

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnection(conn, duration, length)
	}

}

func handleConnection(conn net.Conn, duration *int, length *int) {
	input := make([]byte, *length)

	//fmt.Println(input)
	var i int

	go doEvery((time.Duration(*interval) * time.Second))
	for start := time.Now(); time.Since(start) < time.Second*(time.Duration(*duration)); {
		i++
		startTimer := time.Now()
		_, err := conn.Write([]byte("Hello World!"))
		checkError(err)

		_, err = conn.Read(input[0:])
		checkError(err)
		// check if its the same length or if it is received generally
		elapsed = time.Since(startTimer)
		//log.Printf("[INFO] Process time: %s", elapsed)
		result = append(result, int(elapsed))
	}

}

func tcpClient(port *string, wg *sync.WaitGroup, length *int) {
	input := make([]byte, *length)

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

func tcpServerClient(port *string, wg *sync.WaitGroup, wgParallel *sync.WaitGroup, length *int) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	wg.Done()

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnectionClient(conn, length)
	}
}

func tcpClientClient(port *string, wg *sync.WaitGroup, duration *int, wgParallel *sync.WaitGroup, wgFin *sync.WaitGroup, length *int) {
	input := make([]byte, *length)
	var i int

	wg.Wait()

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	wgParallel.Done()
	wgParallel.Wait()

	go doEvery((time.Duration(*interval) * time.Second))
	for start := time.Now(); time.Since(start) < time.Second*(time.Duration(*duration)); {
		i++

		startTimer := time.Now()
		_, err = conn.Write([]byte("Hello World!"))
		checkError(err)

		_, err := conn.Read(input[0:])
		checkError(err)

		elapsed = time.Since(startTimer)
		// Append time to array instead of printing it
		//log.Printf("[INFO] Response time: %s", elapsed)
		result = append(result, int(elapsed))

	}

	wgFin.Done()
}

func tcpClientClientFlag(port *string, duration *int, wgParallel *sync.WaitGroup, wgFin *sync.WaitGroup, length *int) {

	fmt.Println("test")
	input := make([]byte, *length)
	var i int

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	wgParallel.Done()
	wgParallel.Wait()

	go doEvery((time.Duration(*interval) * time.Second))
	for start := time.Now(); time.Since(start) < time.Second*(time.Duration(*duration)); {
		i++

		startTimer := time.Now()
		_, err = conn.Write([]byte("Hello World!"))
		checkError(err)

		_, err := conn.Read(input[0:])
		checkError(err)

		elapsed = time.Since(startTimer)
		// Append time to array instead of printing it
		//log.Printf("[INFO] Response time: %s", elapsed)
		result = append(result, int(elapsed))

	}

	wgFin.Done()

}

func doEvery(d time.Duration) {
	for x := range time.Tick(d) {
		// Das Ergebnis ist in nanosekunden
		fmt.Println(int(elapsed))
		fmt.Println(elapsed)
		fmt.Println(reflect.TypeOf(int(elapsed)))
		_ = x
	}
}

func handleConnectionClient(conn net.Conn, length *int) {

	input := make([]byte, *length)

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
