package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var wgParallel sync.WaitGroup
	var wgFin sync.WaitGroup

	wg.Add(1)
	wgFin.Add(10)

	go tcpServer(&wg, &wgParallel)

	for i := 0; i < 10; i++ {
		go tcpClient(&wg, &wgParallel, &wgFin)
		wgParallel.Add(1)
	}

	wgFin.Wait()
}

func tcpServer(wg *sync.WaitGroup, wgParallel *sync.WaitGroup) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:8888"))
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	wg.Done()

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnection(conn, wgParallel)
	}
}

func tcpClient(wg *sync.WaitGroup, wgParallel *sync.WaitGroup, wgFin *sync.WaitGroup) {
	var input [512]byte

	wg.Wait()

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:8888"))
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	wgParallel.Done()
	wgParallel.Wait()

	_, err = conn.Write([]byte("Hello World!"))
	checkError(err)

	n, err := conn.Read(input[0:])
	checkError(err)
	fmt.Println(string(input[0:n]))

	wgFin.Done()
}

func handleConnection(conn net.Conn, wgParallel *sync.WaitGroup) {
	var index [512]byte

	n, err := conn.Read(index[0:])
	checkError(err)

	_, err = conn.Write(index[0:n])
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
