// Package call contains Nexmo Call Request and Response.
//
// see: https://docs.nexmo.com/voice/call
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
