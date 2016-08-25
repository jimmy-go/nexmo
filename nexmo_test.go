// Package nexmo contains tests for nexmo client.
package nexmo

import (
	"testing"
	"time"
)

func TestSMS(t *testing.T) {
	nmo, err := New("123", "123", time.Second)
	if err != nil {
		t.Logf("new : err [%s]", err)
		t.Fail()
		return
	}

	res, err := nmo.BasicSMS("5215522334455", "nexmotest", "Hi nexmo test")
	// we can take this as a natural error in tests. FIXME;
	if err == ErrInvalidRequest {
		t.Logf("basicSMS : expected error : err [%s]", err)
		return
	}
	if err != nil {
		t.Logf("basicSMS : err [%s]", err)
	}
	if res == nil {
		t.Logf("basicSMS : res nil")
		return
	}
	if len(res.Messages) < 1 {
		t.Logf("messages [%v]", len(res.Messages))
		t.Fail()
	}
}

type T struct {
	Key      string
	Secret   string
	Timeout  time.Duration
	Expected error
}

func TestSMStable(t *testing.T) {
	table := []T{
		T{
			Key:      "",
			Secret:   "123",
			Timeout:  time.Second,
			Expected: ErrInvalidKey,
		},
		T{
			Key:      "123",
			Secret:   "",
			Timeout:  time.Second,
			Expected: ErrInvalidSecret,
		},
	}
	for i := range table {
		x := table[i]
		_, err := New(x.Key, x.Secret, time.Second)
		if err != x.Expected {
			t.Logf("expected [%v] actual [%v]", x.Expected, err)
			t.Fail()
		}
	}
	Must("123", "123", time.Second)
}
