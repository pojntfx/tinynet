package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
)

// DecodeJSONSumInput decodes JSON sum input
type DecodeJSONSumInput struct {
	InputArray []float64 `json:"inputArray"`
	IonCount   int       `json:"ionCount"`
	MyCount    int       `json:"myCount"`
}

// EncodeJSONSumResult encodes JSON sum result
type EncodeJSONSumResult struct {
	SumResult []float64 `json:"sumResult"`
	MyCount   int       `json:"myCount"`
}

// DecodeJSONSoftmaxInput decodes JSON softmax input
type DecodeJSONSoftmaxInput struct {
	InputArray []float64 `json:"inputArray"`
	IonCount   int       `json:"ionCount"`
	MyCount    int       `json:"myCount"`
	Sum        float64   `json:"sum"`
}

// EncodeJSONSoftmaxResult encodes JSON softmax  result
type EncodeJSONSoftmaxResult struct {
	SoftmaxResult []float64 `json:"softmaxResult"`
	MyCount       int       `json:"myCount"`
}

func main() {
	var jsonSumInput [512]byte
	var jsonSoftmaxInput [512]byte

	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:3333")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte(`Connected`))
	checkError(err)

	n, err := conn.Read(jsonSumInput[0:])
	checkError(err)

	a := decodeJSONSumInput(string(jsonSumInput[0:n]))

	var jsonSumResult []float64

	for i := int(math.Ceil(float64(len(a.InputArray))/float64(a.IonCount))) * a.MyCount; i < int(math.Ceil(float64(len(a.InputArray))/float64(a.IonCount)))*a.MyCount+int(math.Ceil(float64(len(a.InputArray))/float64(a.IonCount))) && i < len(a.InputArray); i++ {

		jsonSumResult = append(jsonSumResult, softmaxSum(a.InputArray[i]))
	}

	bytes := encodeJSONSumResult(EncodeJSONSumResult{jsonSumResult, a.MyCount})

	fmt.Println(string(bytes))

	_, err = conn.Write(bytes)
	checkError(err)

	o, err := conn.Read(jsonSoftmaxInput[0:])
	checkError(err)

	b := decodeJSONSoftmaxInput(string(jsonSoftmaxInput[0:o]))

	var jsonSoftmaxResult []float64

	for i := int(math.Ceil(float64(len(b.InputArray))/float64(b.IonCount))) * b.MyCount; i < int(math.Ceil(float64(len(b.InputArray))/float64(b.IonCount)))*b.MyCount+int(math.Ceil(float64(len(b.InputArray))/float64(b.IonCount))) && i < len(b.InputArray); i++ {

		jsonSoftmaxResult = append(jsonSoftmaxResult, softmaxResult(b.Sum, b.InputArray[i]))
	}

	bytes2 := encodeJSONSoftmaxResult(EncodeJSONSoftmaxResult{jsonSoftmaxResult, b.MyCount})

	fmt.Println(string(bytes2))

	_, err = conn.Write(bytes2)
	checkError(err)
}

func softmaxSum(input float64) float64 {
	return math.Exp(input)
}

func softmaxResult(sum float64, input float64) float64 {
	return math.Exp(input) / sum
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func decodeJSONSumInput(input string) DecodeJSONSumInput {

	rawIn := json.RawMessage(input)

	bytes, err := rawIn.MarshalJSON()
	checkError(err)

	var d DecodeJSONSumInput

	err = json.Unmarshal(bytes, &d)
	checkError(err)

	return d
}

func encodeJSONSumResult(s EncodeJSONSumResult) []byte {

	bytes, err := json.Marshal(s)
	checkError(err)

	return bytes
}

func decodeJSONSoftmaxInput(input string) DecodeJSONSoftmaxInput {

	rawIn := json.RawMessage(input)

	bytes, err := rawIn.MarshalJSON()
	checkError(err)

	var d DecodeJSONSoftmaxInput

	err = json.Unmarshal(bytes, &d)
	checkError(err)

	return d
}

func encodeJSONSoftmaxResult(s EncodeJSONSoftmaxResult) []byte {

	bytes, err := json.Marshal(s)
	checkError(err)

	return bytes
}
