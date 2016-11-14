// Package text2speech contains Nexmo Text-To-Speech
// Request and Response.
//
// see: https://docs.nexmo.com/voice/text-to-speech
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
package text2speech

// Request Nexmo TextToSpeech request.
//
// see: https://docs.nexmo.com/voice/text-to-speech/request
type Request struct {
	To               string `url:"to"`
	From             string `url:"from"`
	Text             string `url:"text"`
	Language         string `url:"lg"`
	Voice            string `url:"voice"`
	Repeat           int    `url:"repeat"`
	MachineDetection string `url:"machine_detection"`
	MachineTimeout   string `url:"machine_timeout"`
	Callback         string `url:"callback"`
	CallbackMethod   string `url:"callback_method"`
}

// Response Nexmo TextToSpeech response.
//
// see: https://docs.nexmo.com/voice/text-to-speech/response
type Response struct {
	CallID    string `json:"call_id"`
	To        string `json:"to"`
	Status    string `json:"status"`
	ErrorText string `json:"error_text"`
}
