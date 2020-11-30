package main

import (
	"fmt"
	"log"
	"math"
	"net"

	"github.com/valyala/fastjson"
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
	var JSONArena fastjson.Arena

	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:3333")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte(`Connected`))
	checkError(err)

	n, err := conn.Read(jsonSumInput[0:])
	checkError(err)

	a := decodeJSON(string(jsonSumInput[0:n]))

	var jsonSumResult []float64

	for i := int(math.Ceil(float64(len(a.GetArray("inputArray")))/float64(a.GetInt("ionCount")))) * a.GetInt("myCount"); i < int(math.Ceil(float64(len(a.GetArray("inputArray")))/float64(a.GetInt("ionCount"))))*a.GetInt("myCount")+int(math.Ceil(float64(len(a.GetArray("inputArray")))/float64(a.GetInt("ionCount")))) && i < len(a.GetArray("inputArray")); i++ {

		jsonSumResult = append(jsonSumResult, softmaxSum(a.GetFloat64("inputArray", fmt.Sprintf("%v", i))))
	}

	output := JSONArena.NewObject()

	jsonSumResultArray := JSONArena.NewArray()

	for i := 0; i < len(jsonSumResult); i++ {
		jsonSumResultArray.SetArrayItem(i, JSONArena.NewNumberFloat64(jsonSumResult[i]))
	}

	output.Set("sumResult", jsonSumResultArray)
	output.Set("myCount", JSONArena.NewNumberInt(a.GetInt("myCount")))

	outputDecoded := output.MarshalTo([]byte{})

	fmt.Println(string(outputDecoded))

	_, err = conn.Write(outputDecoded)
	checkError(err)

	o, err := conn.Read(jsonSoftmaxInput[0:])
	checkError(err)

	b := decodeJSON(string(jsonSoftmaxInput[0:o]))

	var jsonSoftmaxResult []float64

	for i := int(math.Ceil(float64(len(b.GetArray("inputArray")))/float64(b.GetInt("ionCount")))) * b.GetInt("myCount"); i < int(math.Ceil(float64(len(b.GetArray("inputArray")))/float64(b.GetInt("ionCount"))))*b.GetInt("myCount")+int(math.Ceil(float64(len(b.GetArray("inputArray")))/float64(b.GetInt("ionCount")))) && i < len(b.GetArray("inputArray")); i++ {

		jsonSoftmaxResult = append(jsonSoftmaxResult, softmaxResult(b.GetFloat64("sum"), a.GetFloat64("inputArray", fmt.Sprintf("%v", i))))
	}

	output2 := JSONArena.NewObject()

	jsonSoftmaxResultArray := JSONArena.NewArray()

	for i := 0; i < len(jsonSoftmaxResult); i++ {
		jsonSoftmaxResultArray.SetArrayItem(i, JSONArena.NewNumberFloat64(jsonSoftmaxResult[i]))
	}

	output2.Set("softmaxResult", jsonSoftmaxResultArray)
	output2.Set("myCount", JSONArena.NewNumberInt(a.GetInt("myCount")))

	outputDecoded2 := output2.MarshalTo([]byte{})

	fmt.Println(string(outputDecoded2))

	_, err = conn.Write(outputDecoded2)
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

func decodeJSON(input string) *fastjson.Value {

	var p fastjson.Parser
	v, err := p.Parse(input)
	checkError(err)

	return v
}
