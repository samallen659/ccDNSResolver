package message_test

import (
	"github.com/samallen659/ccDNSResolver/internal/message"
	"reflect"
	"testing"
)

func TestResourceRecord(t *testing.T) {
	expected := message.ResourceRecord{
		Name:     "www.northlincs.gov.uk",
		Type:     message.TYPE_A,
		Class:    message.CLASS_IN,
		TTL:      uint32(111),
		RDLength: uint16(4),
		RData:    []byte{5, 134, 12, 204},
	}
	t.Run("Parse correct ResourceRecord from DNS response with single RR and pointer to full name", func(t *testing.T) {
		b := []byte{47, 228, 129, 128, 0, 1, 0, 1, 0, 0, 0, 0, 3, 119, 119, 119, 10, 110, 111, 114, 116, 104, 108, 105, 110, 99,
			115, 3, 103, 111, 118, 2, 117, 107, 0, 0, 1, 0, 1, 192, 12, 0, 1, 0, 1, 0, 0, 0, 111, 0, 4, 5, 134, 12, 204}

		rrs, err := message.ParseResourceRecords(&b, 39, 1)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(*rrs[0], expected) {
			t.Errorf("ResourcRecord does not match expected, received: %#v, expected: %#v", rrs[0], expected)
		}
	})
	t.Run("Parse correct ResourceRecord from DNS response with single RR and labels to full name", func(t *testing.T) {
		b := []byte{47, 228, 129, 128, 0, 1, 0, 1, 0, 0, 0, 0, 3, 119, 119, 119, 10, 110, 111, 114, 116, 104, 108, 105, 110, 99,
			115, 3, 103, 111, 118, 2, 117, 107, 0, 0, 1, 0, 1, 3, 119, 119, 119, 10, 110, 111, 114, 116, 104, 108, 105, 110, 99,
			115, 3, 103, 111, 118, 2, 117, 107, 0, 0, 1, 0, 1, 0, 0, 0, 111, 0, 4, 5, 134, 12, 204}

		rrs, err := message.ParseResourceRecords(&b, 39, 1)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(*rrs[0], expected) {
			t.Errorf("ResourcRecord does not match expected, received: %#v, expected: %#v", rrs[0], expected)
		}
	})
	t.Run("Parse correct ResourceRecords from DNS response with multiple RR in labels and pointer format", func(t *testing.T) {
		b := []byte{47, 228, 129, 128, 0, 1, 0, 2, 0, 0, 0, 0, 3, 119, 119, 119, 10, 110, 111, 114, 116, 104, 108, 105, 110, 99,
			115, 3, 103, 111, 118, 2, 117, 107, 0, 0, 1, 0, 1, 192, 12, 0, 1, 0, 1, 0, 0, 0, 111, 0, 4, 5, 134, 12, 204, 3, 119,
			119, 119, 10, 110, 111, 114, 116, 104, 108, 105, 110, 99, 115, 3, 103, 111, 118, 2, 117, 107, 0, 0, 1, 0, 1, 0, 0, 0,
			111, 0, 4, 5, 134, 12, 204}

		rrs, err := message.ParseResourceRecords(&b, 39, 2)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(*rrs[0], expected) {
			t.Fatalf("First ResourcRecord does not match expected, received: %#v, expected: %#v", rrs[0], expected)
		}
		if !reflect.DeepEqual(*rrs[0], expected) {
			t.Errorf("Second ResourcRecord does not match expected, received: %#v, expected: %#v", rrs[1], expected)
		}
	})
}
