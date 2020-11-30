package main

import (
	"fmt"
	"os"

	"github.com/pojntfx/webassembly-berkeley-sockets-via-webrtc/examples/go/pkg/tinynet"
)

var (
	LADDR  = "10.0.0.240:1234"
	BUFLEN = 1024
)

func main() {
	laddr, err := tinynet.ResolveTCPAddr("tcp", LADDR)
	if err != nil {
		fmt.Println("could not resolve TCP address", err)

		os.Exit(1)
	}

	lis, err := tinynet.ListenTCP("tcp", laddr)
	if err != nil {
		fmt.Println("could not listen", err)

		os.Exit(1)
	}

	fmt.Println("Listening on", LADDR)

	for {
		conn, err := lis.AcceptTCP()
		if err != nil {
			fmt.Println("could not accept", err)

			os.Exit(1)
		}

		fmt.Println("Client connected")

		go func(innerConn *tinynet.TCPConn) {
			for {
				buf := make([]byte, BUFLEN)
				if n, err := innerConn.Read(buf); err != nil {
					if n == 0 {
						break
					}

					fmt.Println("could not read from connection, removing connection", err)

					break
				}

				out := []byte(fmt.Sprintf("You've sent: %v", string(buf)))
				if n, err := innerConn.Write(out); err != nil {
					if n == 0 {
						break
					}

					fmt.Println("could not write from connection, removing connection", err)

					break
				}
			}

			fmt.Println("Client disconnected")

			if err := innerConn.Close(); err != nil {
				fmt.Println("could not close connection", err)
			}

			return
		}(conn)
	}
}
