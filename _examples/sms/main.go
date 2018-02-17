// Package main contains Nexmo SMS example usage.
package main

import (
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/jimmy-go/nexmo"
	"github.com/jimmy-go/nexmo/sms"
)

func main() {
	key := flag.String("api-key", "", "Nexmo API KEY.")
	secret := flag.String("api-secret", "", "Nexmo API SECRET.")
	to := flag.String("to", "", "Nexmo phone destination.")
	from := flag.String("from", "", "Your Nexmo phone number.")
	text := flag.String("text", "", "SMS message content.")
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
