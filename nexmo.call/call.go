// Package call contains Nexmo Call Request and Response.
//
// see: https://docs.nexmo.com/voice/call
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
package call

// Request Nexmo TextToSpeech request.
//
// see: https://docs.nexmo.com/voice/call/request
type Request struct {
	To               string `url:"to"`
	AnswerURL        string `url:"answer_url"`
	From             string `url:"from"`
	MachineDetection string `url:"machine_detection"`
	MachineTimeout   string `url:"machine_timeout"`
	AnswerMethod     string `url:"answer_method"`
	ErrorURL         string `url:"error_url"`
	ErrorMethod      string `url:"error_method"`
	StatusURL        string `url:"status_url"`
	StatusMethod     string `url:"status_method"`
}

// Response Nexmo TextToSpeech response.
//
// see: https://docs.nexmo.com/voice/call/response
type Response struct {
	CallID    string `json:"call-id"`
	To        string `json:"to"`
	Status    int    `json:"status"`
	ErrorText string `json:"error-text"`
}
