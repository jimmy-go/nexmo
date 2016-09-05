// Package sms contains Nexmo SMS Request and Response as in
//
// see: https://docs.nexmo.com/messaging/sms-api/api-reference
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
package sms

const (
	// StatusOK 0 - Delivered.
	StatusOK = "0"

	// StatusUnknown 1 - Unknown - either:
	// * An unknown error was received from the carrier who
	//   tried to send this this message.
	// * Depending on the carrier, that to is unknown.
	// When you see this error, and status is rejected,
	// always check if to in your request was valid.
	StatusUnknown = "1"

	// StatusAbsentSubscriberTemporary 2 - Absent Subscriber
	// Temporary - this message was not delivered because
	// to was temporarily unavailable. For example,
	// the handset used for to was out of coverage or
	// switched off. This is a temporary failure,
	// retry later for a positive result.
	StatusAbsentSubscriberTemporary = "2"

	// StatusAbsentSubscriberPermanent 3 - Absent Subscriber
	// Permanent - to is no longer active, you should
	// remove this phone number from your database.
	StatusAbsentSubscriberPermanent = "3"

	// StatusCallBarredUser 4 - Call barred by user - you should
	// remove this phone number from your database. If the
	// user wants to receive messages from you, they need
	// to contact their carrier directly.
	StatusCallBarredUser = "4"

	// StatusPortabilityError 5 - Portability Error - there
	// is an issue after the user has changed carrier for
	// to. If the user wants to receive messages from you,
	// they need to contact their carrier directly.
	StatusPortabilityError = "5"

	// StatusAntiSpamRejection 6 - Anti-Spam Rejection -
	// carriers often apply restrictions that block messages
	// following different criteria. For example, on
	// SenderID or message content.
	StatusAntiSpamRejection = "6"

	// StatusHandsetBusy 7 - Handset Busy - the handset
	// associated with to was not available when this
	// message was sent. If status is Failed, this is a
	// temporary failure; retry later for a positive result.
	// If status is Accepted, this message has is in the
	// retry scheme and will be resent until it expires
	// in 24-48 hours.
	StatusHandsetBusy = "7"

	// StatusNetworkError 8 - Network Error - a network
	// failure while sending your message. This is a
	// temporary failure, retry later for a positive result.
	StatusNetworkError = "8"

	// StatusIllegalNumber 9 - Illegal Number - you tried
	// to send a message to a blacklisted phone number.
	// That is, the user has already sent a STOP opt-out
	// message and no longer wishes to receive messages
	// from you.
	StatusIllegalNumber = "9"

	// StatusInvalidMessage 10 - Invalid Message - the
	// message could not be sent because one of the
	// parameters in the message was incorrect.
	// For example, incorrect type or udh.
	StatusInvalidMessage = "10"

	// StatusUnroutable 11 - Unroutable - the
	// chosen route to send your message is not available.
	// This is because the phone number is either:
	// * currently on an unsupported network.
	// * On a pre-paid or reseller account that
	//   could not receive a message sent by from.
	// To resolve this issue either email us at
	// support@nexmo.com or create a helpdesk ticket
	// at https://help.nexmo.com.
	StatusUnroutable = "11"

	// StatusDestinationUnreachable 12 - Destination
	// unreachable - the message could not be delivered to
	// the phone number.
	StatusDestinationUnreachable = "12"

	// StatusAgeRestriction 13 - Subscriber Age Restriction
	// - the carrier blocked this message because the
	// content is not suitable for to based on
	// age restrictions.
	StatusAgeRestriction = "13"

	// StatusBlockedByCarrier 14 - Number Blocked by Carrier
	// - the carrier blocked this message. This could be
	// due to several reasons. For example, to's plan does
	// not include SMS or the account is suspended.
	StatusBlockedByCarrier = "14"

	// StatusPrePaidInsufficient 15 - Pre-Paid - Insufficent
	// funds - to’s pre-paid account does not have enough
	// credit to receive the message.
	StatusPrePaidInsufficient = "15"

	// StatusGeneralError 99 - General Error - there is a
	// problem with the chosen route to send your message.
	// To resolve this issue either email us at
	// support@nexmo.com or create a helpdesk ticket
	// at https://help.nexmo.com.
	StatusGeneralError = "99"
)

// Request Nexmo SMS request.
//
// see: https://docs.nexmo.com/messaging/sms-api/api-reference#request
type Request struct {
	From         string `url:"from"`
	To           string `url:"to"`
	Type         string `url:"type"`
	Text         string `url:"text"`
	StatusReport string `url:"status-report-req"`
	ClientRef    string `url:"client-ref"`
	Vcard        string `url:"vcard"`
	Vcal         string `url:"vcal"`
	Callback     string `url:"callback"`
	MessageClass string `url:"message-class"`
	UDH          string `url:"udh"`
	ProtocolID   string `url:"protocol-id"`
	Body         string `url:"body"`
	Title        string `url:"title"`
	URL          string `url:"url"`
	Validity     string `url:"validity"`
}

// Response Nexmo SMS response.
//
// see: https://docs.nexmo.com/messaging/sms-api/api-reference#response
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

// DeliveryReceipt for webhook endpoint.
//
// see: https://docs.nexmo.com/messaging/sms-api/api-reference#delivery_receipt
type DeliveryReceipt struct {
	To               string `json:"to"`
	NetworkCode      string `json:"network-code"`
	MessageID        string `json:"messageId"`
	Msisdn           string `json:"msisdn"`
	Status           string `json:"status"`
	ErrCode          string `json:"err-code"`
	Price            string `json:"price"`
	Scts             string `json:"scts"`
	MessageTimestamp string `json:"message-timestamp"`
	ClientRef        string `json:"client-ref"`
}
