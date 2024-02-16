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
}
