// Package nexmo contains Nexmo client using oficial documentation.
//
// see: https://docs.nexmo.com
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
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/jimmy-go/nexmo/nexmo.call"
	"github.com/jimmy-go/nexmo/nexmo.sms"
	"github.com/jimmy-go/nexmo/nexmo.text2speech"
)

var (
	// ErrInvalidKey returned when API KEY is empty.
	ErrInvalidKey = errors.New("nexmo: invalid key length")

	// ErrInvalidSecret returned when API SECRET is empty.
	ErrInvalidSecret = errors.New("nexmo: invalid key length")

	// ErrEmptyResponse returned if response has no messages.
	ErrEmptyResponse = errors.New("nexmo: response is empty")

	// ErrBadRequest is return when http response status
	// is not 200.
	ErrBadRequest = errors.New("nexmo: bad request")

	// ErrSupportNotFound returned when this package has no
	// implementation for some feature. See supportmap map.
	ErrSupportNotFound = errors.New("nexmo: feature not supported")
)

const (
	// EndpointSMS Nexmo API endpoint.
	EndpointSMS = "https://rest.nexmo.com/sms/json?"

	// EndpointCall Nexmo API endpoint.
	EndpointCall = "https://rest.nexmo.com/call/json?"

	// EndpointText2Speech Nexmo API endpoint.
	EndpointText2Speech = "https://api.nexmo.com/tts/json?"
)

// Nexmo client
type Nexmo struct {
	key    string
	secret string
	client *http.Client
	sync.RWMutex
}

// New returns a new Nexmo client with timeout.
func New(key, secret string, timeout time.Duration) (*Nexmo, error) {
	if len(key) < 1 {
		return nil, ErrInvalidKey
	}
	if len(secret) < 1 {
		return nil, ErrInvalidSecret
	}
	n := &Nexmo{
		key:    key,
		secret: secret,
		client: &http.Client{
			Timeout: timeout,
		},
	}
	return n, nil
}

// Must calls New func or panic.
func Must(key, secret string, timeout time.Duration) *Nexmo {
	nex, err := New(key, secret, timeout)
	if err != nil {
		panic(err)
	}
	return nex
}

// Support struct
type Support struct {
	DocURL string
	Method string
	URL    string
}

var (
	// supportmap contains general endpoints for nexmo client
	// using oficial nexmo docs.
	supportmap = map[string]*Support{
		"sms": &Support{
			DocURL: "https://docs.nexmo.com/messaging/sms-api/api-reference",
			Method: "POST",
			URL:    EndpointSMS,
		},
		"call": &Support{
			DocURL: "https://docs.nexmo.com/voice/call",
			Method: "POST",
			URL:    EndpointCall,
		},
		"text2speech": &Support{
			DocURL: "https://docs.nexmo.com/voice/text-to-speech",
			Method: "POST",
			URL:    EndpointText2Speech,
		},
	}
)

// do internal client request doer.
func (x *Nexmo) do(p url.Values, supportType string, dst interface{}) error {
	x.RLock()
	defer x.RUnlock()
	resource, ok := supportmap[supportType]
	if !ok {
		return ErrSupportNotFound
	}
	// clean empty keys
	for k, val := range p {
		if len(val) < 1 {
			p.Del(k)
			continue
		}
		if len(val[0]) < 1 {
			p.Del(k)
			continue
		}
	}
	// force credentials
	p.Set("api_key", x.key)
	p.Set("api_secret", x.secret)
	uri := resource.URL + p.Encode()
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}
	resp, err := x.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return ErrBadRequest
	}
	err = json.NewDecoder(resp.Body).Decode(dst)
	if err != nil {
		buf := bytes.NewBuffer([]byte{})
		_, _ = buf.ReadFrom(resp.Body)
		log.Printf("Nexmo : do : err body [%v]", buf.String())
		return err
	}
	return nil
}

// NewSMS returns a new SMS request only with required fields.
// see: https://docs.nexmo.com/messaging/sms-api/api-reference#request
func NewSMS(to, from, text string) *sms.Request {
	req := &sms.Request{
		To:   to,
		From: from,
		Text: text,
	}
	return req
}

// SMS use SMS API to send and receive a high volume of SMS
// anywhere in the world.
//
// see: https://docs.nexmo.com/messaging/sms-api
func (x *Nexmo) SMS(r *sms.Request) (*sms.Response, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, err
	}
	var res *sms.Response
	err = x.do(v, "sms", &res)
	if err != nil {
		return res, err
	}
	if len(res.Messages) < 1 {
		return res, ErrEmptyResponse
	}
	return res, nil
}

// NewCall returns a new call request with only required fields.
// see: https://docs.nexmo.com/voice/call/request
func NewCall(to, answerURL string) *call.Request {
	req := &call.Request{
		To:        to,
		AnswerURL: answerURL,
	}
	return req
}

// Call You use Call API to make outbound calls from Nexmo
// virtual numbers to other phone numbers.
func (x *Nexmo) Call(r *call.Request) (*call.Response, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, err
	}
	var res *call.Response
	err = x.do(v, "call", &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// NewText2Speech returns a new text2speech request with only
// required fields.
//
// see: https://docs.nexmo.com/voice/text-to-speech/request
func NewText2Speech(to, from, text, lang, voice string) *text2speech.Request {
	req := &text2speech.Request{
		To:       to,
		From:     from,
		Text:     text,
		Language: lang,
		Voice:    voice,
	}
	return req
}

// Text2Speech You use Text-To-Speech API to send
// synthesized speech or recorded sound files to a phone number
func (x *Nexmo) Text2Speech(r *text2speech.Request) (*text2speech.Response, error) {
	v, err := query.Values(r)
	if err != nil {
		return nil, err
	}
	var res *text2speech.Response
	err = x.do(v, "text2speech", &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
