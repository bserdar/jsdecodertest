package main

import (
	"bytes"
	"fmt"
	//"github.com/bserdar/jsdecodertest/json"
	"encoding/json"
	"reflect"
	"runtime"
)

func main() {
	str := `{
   "1234568790":"",
   "0987654321":"",
   "3456218736":""}`

	out := bytes.Buffer{}
	out.WriteRune('[')
	for i := 0; i < 1000000; i++ {
		if i > 0 {
			out.WriteRune(',')
		}
		out.WriteString(str)
	}
	out.WriteRune(']')

	var m interface{}
	dec := json.NewDecoder(bytes.NewReader(out.Bytes()))
	dec.Decode(&m)
	var stats runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats)
	fmt.Printf("%+v\n", stats)
	// Make sure there's still reference to m
	reflect.ValueOf(m)
}
