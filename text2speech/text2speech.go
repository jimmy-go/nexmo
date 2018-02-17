// Package text2speech contains Nexmo Text-To-Speech
// Request and Response.
//
// see: https://docs.nexmo.com/voice/text-to-speech
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
