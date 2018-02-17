// Package nexmo contains tests for nexmo client.
package nexmo

import (
	"testing"
	"time"
)

type T struct {
	Input    Input
	Expected error
}

type Input struct {
	Key       string
	Secret    string
	Timeout   time.Duration
	To        string
	From      string
	Text      string
	Lang      string
	Voice     string
	AnswerURL string
}

func TestNew(t *testing.T) {
	table := []T{
		T{
			Input: Input{
				Key:     "",
				Secret:  "",
				Timeout: time.Second,
			},
			Expected: ErrInvalidKey,
		},
		T{
			Input: Input{
				Key:     "123",
				Secret:  "",
				Timeout: time.Second,
			},
			Expected: ErrInvalidSecret,
		},
	}
	for i := range table {
		x := table[i]
		_, err := New(x.Input.Key, x.Input.Secret, x.Input.Timeout)
		if err != x.Expected {
			t.Logf("expected [%v] actual [%v]", x.Expected, err)
			t.Fail()
			continue
		}
	}
	Must("123", "456", time.Second)
}

func TestTableSMS(t *testing.T) {
	table := []T{
		T{
			Input: Input{
				Key:     "123",
				Secret:  "123",
				Timeout: time.Second * 10,
			},
			Expected: nil,
		},
		T{
			Input: Input{
				Key:     "123",
				Secret:  "123",
				Timeout: time.Second * 10,
			},
			Expected: nil,
		},
	}
	for i := range table {
		x := table[i]
		client, err := New(x.Input.Key, x.Input.Secret, x.Input.Timeout)
		if err != nil {
			t.Logf("new : err [%v]", err)
			t.Fail()
			continue
		}
		msg := NewSMS(x.Input.To, x.Input.From, x.Input.Text)
		_, err = client.SMS(msg)
		if err != x.Expected {
			t.Logf("expected [%v] actual [%v]", x.Expected, err)
			t.Fail()
			continue
		}
	}
}

// TODO; review, for now sms method is priority.
func TestTableText2Speech(t *testing.T) {
	table := []T{
		T{
			Input: Input{
				Key:     "123",
				Secret:  "123",
				Timeout: time.Second * 10,
			},
			// Expected: ErrBadRequest,
			Expected: nil,
		},
		T{
			Input: Input{
				Key:     "123",
				Secret:  "123",
				Timeout: time.Second * 10,
			},
			// Expected: ErrBadRequest,
			Expected: nil,
		},
	}
	for i := range table {
		x := table[i]
		client, err := New(x.Input.Key, x.Input.Secret, x.Input.Timeout)
		if err != nil {
			t.Errorf("new : err [%v]", err)
			continue
		}
		msg := NewText2Speech(x.Input.To, x.Input.From, x.Input.Text, x.Input.Lang, x.Input.Voice)
		_, err = client.Text2Speech(msg)
		if err != x.Expected {
			t.Errorf("expected [%v] actual [%v]", x.Expected, err)
			continue
		}
	}
}

func TestTableCall(t *testing.T) {
	table := []T{
		T{
			Input: Input{
				Key:       "123",
				Secret:    "123",
				Timeout:   time.Second,
				To:        "5215522334455",
				AnswerURL: "http://localhost/somexml.xml",
			},
			Expected: nil,
		},
		T{
			Input: Input{
				Key:       "123",
				Secret:    "123",
				Timeout:   time.Second,
				To:        "5215522334455",
				AnswerURL: "http://localhost/somexml.xml",
			},
			Expected: nil,
		},
	}
	for i := range table {
		x := table[i]
		client, err := New(x.Input.Key, x.Input.Secret, x.Input.Timeout)
		if err != nil {
			t.Logf("new : err [%v]", err)
			t.Fail()
			continue
		}
		msg := NewCall(x.Input.To, x.Input.AnswerURL)
		_, err = client.Call(msg)
		if err != x.Expected {
			t.Logf("expected [%v] actual [%v]", x.Expected, err)
			t.Fail()
			continue
		}
	}
}
