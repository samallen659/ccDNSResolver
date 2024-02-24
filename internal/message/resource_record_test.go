package message_test

import (
	"testing"

	"github.com/samallen659/ccDNSResolver/internal/message"
)

func TestResourceRecord(t *testing.T) {
	t.Run("Parse correct ResourceRecord from DNS response with single RR and pointer to full name", func(t *testing.T) {
		b := []byte{47, 228, 129, 128, 0, 1, 0, 1, 0, 0, 0, 0, 3, 119, 119, 119, 10, 110, 111, 114, 116, 104, 108, 105, 110, 99,
			115, 3, 103, 111, 118, 2, 117, 107, 0, 0, 1, 0, 1, 192, 12, 0, 1, 0, 1, 0, 0, 1, 44, 0, 4, 5, 134, 12, 204, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		message.ParseResourceRecords(&b, 39, 1)
		t.Fail()
	})
}
