package main

import (
	"fmt"
	"log"

	"github.com/valyala/fastjson"
)

var (
	JSONParser fastjson.Parser
	JSONArena  fastjson.Arena
)

func encode() {

}

func decode() {

}

func main() {
	//arr := []float64{1, 1, 3}
	// encode
	arr := []float64{1, 1, 3}
	output := JSONArena.NewObject()

	sum := JSONArena.NewArray()

	for i := 0; i < len(arr); i++ {
		sum.SetArrayItem(i, JSONArena.NewNumberFloat64(arr[i]))
	}

	// Hier kann man dann aber schon den ganzen Output zusammensetzen
	// Also hier setzen wir dann den Output zusammen
	output.Set("arr", sum)

	outputDecoded := output.MarshalTo([]byte{})

	fmt.Println(string(outputDecoded))
	// start with string and return object
	// var p fastjson.Parser
	// v, err := p.Parse(`{
	// 	"str": "bar",
	// 	"int": 123,
	// 	"float": 1.23,
	// 	"bool": true,
	// 	"arr": [1, "foo", {}]
	// }`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(v)

	// end with object and get bytes or string out -- brauchen wir vielleicht nicht. Ich glaube wir koennen auch einfach das Object v aus dem Tutorial writen
	newJSON := stringToJSON(`{
		"str": "bar",
		"int": 123,
		"float": 1.23,
		"bool": true,
		"arr": [1.2, 2.0, 1.3]
}`)
	fmt.Println(newJSON.GetStringBytes("str"))
	fmt.Println((newJSON.GetFloat64("arr", fmt.Sprintf("%v", 0))))

}

func stringToJSON(input string) *fastjson.Value {
	var p fastjson.Parser
	v, err := p.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	return v
}
