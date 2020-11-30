package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"sync"

	"github.com/valyala/fastjson"
)

// DecodeJSONSumResult decodes JSON sum results
type DecodeJSONSumResult struct {
	SumResult []float64 `json:"sumResult"`
	MyCount   int       `json:"myCount"`
}

// EncodeJSONSumInput encodes JSON sum input
type EncodeJSONSumInput struct {
	InputArray []float64 `json:"inputArray"`
	IonCount   int       `json:"ionCount"`
	MyCount    int       `json:"myCount"`
}

// EncodeJSONSoftmaxInput encodes JSON softmax input
type EncodeJSONSoftmaxInput struct {
	InputArray []float64 `json:"inputArray"`
	IonCount   int       `json:"ionCount"`
	MyCount    int       `json:"myCount"`
	Sum        float64   `json:"sum"`
}

// DecodeJSONSoftmaxResult decodes JSON softmax result
type DecodeJSONSoftmaxResult struct {
	SoftmaxResult []float64 `json:"softmaxResult"`
	MyCount       int       `json:"myCount"`
}

// Softmax provides variables for the softmax calculation
type Softmax struct {
	sumResultArray     []float64
	softmaxResultArray []float64
	sumResult          float64
	inputArray         []float64
	ionCount           int
}

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:3333")
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	inputArray := []float64{1, 1, 3}
	var data = Softmax{make([]float64, len(inputArray)), make([]float64, len(inputArray)), 0, inputArray, 0}

	var wgSum sync.WaitGroup
	var wgSoftmax sync.WaitGroup
	var wgStart sync.WaitGroup
	var wgStart2 sync.WaitGroup

	id := 0

	go manager(&wgSum, &wgSoftmax, &data, &wgStart, &wgStart2)

	for {
		conn, err := ln.Accept()
		checkError(err)

		wgSum.Add(1)
		wgSoftmax.Add(1)

		go handleConnection(conn.(*net.TCPConn), &wgSum, &wgSoftmax, id, &data, &wgStart, &wgStart2)

		id++
		data.ionCount++

	}

}

func manager(wgSum *sync.WaitGroup, wgSoftmax *sync.WaitGroup, data *Softmax, wgStart *sync.WaitGroup, wgStart2 *sync.WaitGroup) {

	reader := bufio.NewReader(os.Stdin)

	wgStart.Add(1)
	wgStart2.Add(1)

	fmt.Println("[INFO] Press ENTER to start calculation")
	input, _ := reader.ReadString('\n')
	_ = input

	wgStart.Done()

	wgSum.Wait()

	for i := 0; i < len(data.sumResultArray); i++ {
		data.sumResult += data.sumResultArray[i]
	}

	wgStart2.Done()

	wgSoftmax.Wait()

	fmt.Println(data.softmaxResultArray)

}

func handleConnection(conn *net.TCPConn, wgSum *sync.WaitGroup, wgSoftmax *sync.WaitGroup, id int, data *Softmax, wgStart *sync.WaitGroup, wgStart2 *sync.WaitGroup) {
	var input [512]byte
	var JSONArena fastjson.Arena

	n, err := conn.Read(input[0:])
	checkError(err)

	wgStart.Wait()

	output := JSONArena.NewObject()

	inputArray := JSONArena.NewArray()

	for i := 0; i < len(data.inputArray); i++ {
		inputArray.SetArrayItem(i, JSONArena.NewNumberFloat64(data.inputArray[i]))
	}

	output.Set("inputArray", inputArray)
	output.Set("ionCount", JSONArena.NewNumberInt(data.ionCount))
	output.Set("myCount", JSONArena.NewNumberInt(id))

	outputEncoded := output.MarshalTo([]byte{})

	_, err = conn.Write(outputEncoded)
	checkError(err)

	n, err = conn.Read(input[0:])
	checkError(err)

	sumResultChunk := decodeJSON(string(input[0:n]))

	for i := 0; i < len(sumResultChunk.GetArray("sumResult")); i++ {
		data.sumResultArray[i+(int(math.Ceil(float64(len(data.inputArray))/float64(data.ionCount)))*sumResultChunk.GetInt("myCount"))] = sumResultChunk.GetFloat64("sumResult", fmt.Sprintf("%v", i))
	}

	wgSum.Done()

	wgStart2.Wait()

	output2 := JSONArena.NewObject()

	output2.Set("inputArray", inputArray)
	output2.Set("ionCount", JSONArena.NewNumberInt(data.ionCount))
	output2.Set("myCount", JSONArena.NewNumberInt(id))
	output2.Set("sum", JSONArena.NewNumberFloat64(data.sumResult))

	outputDecoded2 := output2.MarshalTo([]byte{})

	_, err = conn.Write([]byte(outputDecoded2))
	checkError(err)

	o, err := conn.Read(input[0:])
	checkError(err)

	softmaxResultChunk := decodeJSON(string(input[0:o]))

	for i := 0; i < len(softmaxResultChunk.GetArray("softmaxResult")); i++ {
		data.softmaxResultArray[i+(int(math.Ceil(float64(len(data.inputArray))/float64(data.ionCount)))*softmaxResultChunk.GetInt("myCount"))] = softmaxResultChunk.GetFloat64("softmaxResult", fmt.Sprintf("%v", i))
	}

	wgSoftmax.Done()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func decodeJSON(input string) *fastjson.Value {

	var p fastjson.Parser
	v, err := p.Parse(input)
	checkError(err)

	return v
}
