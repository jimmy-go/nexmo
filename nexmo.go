// Package nexmo contains Nexmo client API.
//
// see:
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
package nexmo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Nexmo client
type Nexmo struct {
	Key    string
	Secret string

	Client *http.Client
}

// Send sends a sms using nexmo client and returns carrier code if present.
// FIXME; check working scenario. Until nexmo account is set.
func (x *Nexmo) Send(country, phone, text string) (string, error) {
	// validate lada.
	lada, err := k.CountryLada(country)
	if err != nil {
		return "", err
	}
	// generate nexmo API request.
	to := lada + phone
	from := "Tander"
	uri := fmt.Sprintf("%s?api_key=%s&api_secret=%s&to=%s&from=%s&text=%s",
		restEndpoint, x.Key, x.Secret, to, from,
		strings.Replace(text, " ", "+", -1))

	resp, err := x.Client.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res *Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}

	if len(res.Messages) < 1 {
		return "", errNexmoResponseEmpty
	}

	// TODO; validate when there is more Messages
	for i := range res.Messages {
		m := res.Messages[i]
		if m.Status != StatusOK {
			return "", fmt.Errorf("nexmo: sms: %s", m.ErrorText)
		}
	}

	// by default there is 1 Message
	carrier := res.Messages[0].Network
	return carrier, nil
}

// Call calls using nexmo client.
func (x *Nexmo) Call(country, phone, speech string) (string, error) {
	// TODO;
	return "", errInitRequired
}

// Response response from nexmo sms.
type Response struct {
	MessageCount string     `json:"message-count"`
	Messages     []*Message `json:"messages"`
}

// Message inside nexmo response.
type Message struct {
	Status           string `json:"status"`
	MessageID        string `json:"message-id"`
	To               string `json:"to"`
	ClientRef        string `json:"client-ref"`
	RemainingBalance string `json:"remaining-balance"`
	MessagePrice     string `json:"message-price"`
	Network          string `json:"network"`
	ErrorText        string `json:"error-text"`
}
