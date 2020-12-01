package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pojntfx/tinynet/pkg/tinynet"
)

func main() {

	port := flag.String("p", "3333", "port to listen to")
	format := flag.String("f", "M", "specify the format of bandwidth numbers. (k = Kbits/sec, K = KBytes/sec, m = Mbits/sec, M = MBytes/sec)")
	interval := flag.Int("i", 0, "set interval between periodic bandwidth, jitter, ans loss reports")
	verbose := flag.Bool("V", false, "give more detailed output")
	server := flag.Bool("s", false, "run in server mode")
	//client := flag.Bool("c", false, "run in client mode")
	time := flag.Int("t", 10, "time in seconds to transmit for")
	length := flag.Int("l", 128, "length of buffers to read or write (in KB)")
	parallel := flag.Int("P", 1, "number of simultaneous connections to make to the server")
	reverse := flag.Bool("R", false, "run in reverse mode (server sends, client receives)")

	flag.Parse()

	fmt.Println("port:", *port)
	fmt.Println("format:", *format)
	fmt.Println("interval:", *interval)
	fmt.Println("verbose:", *verbose)
	fmt.Println("server:", *server)
	//	fmt.Println("client:", *client)
	fmt.Println("time:", *time)
	fmt.Println("length:", *length)
	fmt.Println("parallel:", *parallel)
	fmt.Println("reverse:", *reverse)

}

func tcpServer(port *string) {

	tcpAddr, err := tinynet.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%v", *port))
	checkError(err)

	ln, err := tinynet.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnection(conn.(*tinynet.TCPConn))
	}
}

func handleConnection(conn *tinynet.TCPConn) {
	var input []byte

	_, err := conn.Write([]byte("Hello World!"))
	checkError(err)

	n, err := conn.Read(input[0:])
	checkError(err)

	echo := string(input[0:n])

	fmt.Println(echo)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
