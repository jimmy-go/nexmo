####Nexmo client written in Go.

Not tested with credentials.

[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/jimmy-go/nexmo.svg?branch=master)](https://travis-ci.org/jimmy-go/nexmo)
[![Go Report Card](https://goreportcard.com/badge/github.com/jimmy-go/nexmo)](https://goreportcard.com/report/github.com/jimmy-go/nexmo)
[![GoDoc](http://godoc.org/github.com/jimmy-go/nexmo?status.png)](http://godoc.org/github.com/jimmy-go/nexmo)
[![Coverage Status](https://coveralls.io/repos/github/jimmy-go/nexmo/badge.svg?branch=master)](https://coveralls.io/github/jimmy-go/nexmo?branch=master)

#####Install:
```
go get gopkg.in/jimmy-go/nexmo.v0
```

#####Usage:
Call `nexmo.New` method to get a new Nexmo client.
For every Nexmo feature there is a package with that name
that contains his Request and Response type.

E.g.: For method SMS you need a nexmo.sms.Request{} that returns
nexmo.sms.Response{}

```
# declare a new client
client, err := nexmo.New("APIKEY", "APISECRET")
// check errors...

# use it
msg := nexmo.NewSMS("5215522334455", "NexmoTest", "Hello world!")
resp, err := client.SMS(msg)
// resp = nexmo.sms.Response

req := nexmo.NewCall("5215522334455", "http://someurl/answer.xml")
resp, err := client.Call(req)
// resp = nexmo.call.Response

t2s := nexmo.NewText2Speech("5215522334455", "NexmoTest", "Hello my world!", "en-us", "female")
resp, err := client.Text2Speech(t2s)
// resp = nexmo.text2speech.Response
```

#####Credits:

* [NEXMO](https://www.nexmo.com)
* [NEXMO DOC](https://docs.nexmo.com)

#####License:

MIT License

Copyright (c) 2016 Angel Del Castillo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
