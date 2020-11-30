package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pojntfx/webassembly-berkeley-sockets-via-webrtc/examples/go/pkg/tinynet"
)

var (
	RADDR  = "10.0.0.240:1234"
	BUFLEN = 1038
)

func main() {
	conn, err := tinynet.Dial("tcp", RADDR)
	if err != nil {
		fmt.Println("could not listen", err)

		os.Exit(1)
	}

	fmt.Println("Connected to", RADDR)

	reader := bufio.NewReader(os.Stdin)

	for {
		out, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("could not read from stdin", err)

			os.Exit(1)
		}

		if n, err := conn.Write([]byte(out)); err != nil {
			if n == 0 {
				break
			}

			fmt.Println("could not write from connection, removing connection", err)

			break
		}

		buf := make([]byte, BUFLEN)
		if n, err := conn.Read(buf); err != nil {
			if n == 0 {
				break
			}

			fmt.Println("could not read from connection, removing connection", err)

			break
		}

		fmt.Print(string(buf))
	}

	fmt.Println("Disconnected")

	if err := conn.Close(); err != nil {
		fmt.Println("could not close connection", err)
	}
}
