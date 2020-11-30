package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/ugjka/messenger"
)

func main() {

	var wg sync.WaitGroup

	m := messenger.New(0, false)

	wg.Add(1)

	go test(m, &wg)

	var i interface{}
	m.Broadcast(i)

	wg.Wait()
}

func test(m *messenger.Messenger, wg *sync.WaitGroup) {

	start, err := m.Sub()
	if err != nil {
		log.Fatal()
	}

	msg := <-start
	_ = msg

	fmt.Println("klappt")

	wg.Done()
}
