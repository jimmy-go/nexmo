// Package main contains Nexmo SMS example usage.
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
	"github.com/jimmy-go/nexmo/nexmo.sms"
)

var (
	key    = flag.String("api-key", "", "Nexmo API KEY.")
	secret = flag.String("api-secret", "", "Nexmo API SECRET.")
	to     = flag.String("to", "", "Nexmo phone destination.")
	from   = flag.String("from", "", "Your Nexmo phone number.")
	text   = flag.String("text", "", "SMS message content.")
)

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile)

	nclient, err := nexmo.New(*key, *secret, time.Second)
	if err != nil {
		panic(err)
	}

	// new SMS with only required parameters.
	// msg := nexmo.NewSMS(*to, *from, *text)

	// complete sms usage.
	// see: https://docs.nexmo.com/messaging/sms-api/api-reference#request
	msg := &sms.Request{
		From:         *from,
		To:           *to,
		Type:         "",
		Text:         *text,
		StatusReport: "",
		ClientRef:    "",
		Vcard:        "",
		Vcal:         "",
	}
	resp, err := nclient.SMS(msg)
	if err != nil {
		panic(err)
	}

	log.Printf("SMS : response [%v]", marshal(resp))
}

func marshal(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "	")
	return string(b)
}
