// Package main contains Nexmo CALL example usage.
//
// MIT License
//
// Copyright (c) 2016 Angel Del Castillo
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/jimmy-go/nexmo"
)

var (
	key    = flag.String("api-key", "", "Nexmo API KEY.")
	secret = flag.String("api-secret", "", "Nexmo API SECRET.")
	to     = flag.String("to", "", "Nexmo phone destination.")
	answer = flag.String("answer", "", "Answer URL.")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.Printf("Nexmo Key [%s]", *key)
	log.Printf("Nexmo Secret [%s]", *secret)

	client, err := nexmo.New(*key, *secret, time.Second)
	if err != nil {
		panic(err)
	}

	// new CALL with only required parameters.
	req := nexmo.NewCall(*to, *answer)
	resp, err := client.Call(req)
	if err != nil {
		panic(err)
	}

	log.Printf("CALL : response [%v]", marshal(resp))
}

func marshal(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "	")
	return string(b)
}
