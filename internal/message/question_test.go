package message_test

import (
	"github.com/samallen659/ccDNSResolver/internal/message"
	"golang.org/x/exp/slices"
	"testing"
)

func TestQuestion(t *testing.T) {
	t.Run("ConvertHostnameToQName returns correct byte slice", func(t *testing.T) {
		hostname := "www.google.com"
		b := message.ConvertHostnameToQName(hostname)

		if a := slices.Compare(b, []byte{3, 119, 119, 119, 6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0}); a != 0 {
			t.Error("incorret byte slice received")
		}
	})
	t.Run("Marshall returns correct byte slice of question", func(t *testing.T) {
		qname := message.ConvertHostnameToQName("www.google.com")
		question := message.Question{
			QName:  qname,
			QType:  message.QTYPE_A,
			QClass: message.QCLASS_IN,
		}
		b := question.Marshall()

		if len(b) != 20 {
			t.Fatal("Incorrect length byte slice received")
		}

		if a := slices.Compare(b[:16], []byte{3, 119, 119, 119, 6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0}); a != 0 {
			t.Error("Incorrect QName revieced")
		}

		if a := slices.Compare(b[16:18], []byte{0, 1}); a != 0 {
			t.Error("Incorrect QType recieved")
		}

		if a := slices.Compare(b[18:], []byte{0, 1}); a != 0 {
			t.Error("Incorrect QClass received")
		}
	})
}
