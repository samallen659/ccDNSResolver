package message_test

import (
	"encoding/hex"
	"fmt"
	"github.com/samallen659/ccDNSResolver/internal/message"
	"testing"
)

func TestHeader(t *testing.T) {
	t.Run("Marshall returns encoded Header", func(t *testing.T) {
		id := message.NewHeaderID()
		h := message.Header{
			ID:      id,
			QR:      byte(0),
			OPCode:  message.OPCODEStatus,
			AA:      byte(0),
			TC:      byte(1),
			RD:      byte(1),
			RA:      byte(0),
			Z:       byte(6),
			RCode:   message.RCODENoError,
			QDCOUNT: uint16(1),
			ANCOUNT: uint16(0),
			NSCOUNT: uint16(1),
			ARCOUNT: uint16(0),
		}

		b, err := h.Marshall()
		if err != nil {
			t.Fatalf("Marshall failed: %s", err.Error())
		}

		expectedHex := fmt.Sprintf("%x13600001000000010000", id)
		encodedString := hex.EncodeToString(b)
		if encodedString != expectedHex {
			t.Fatalf("Incorret value Recieved, expected: %s recieved: %s",
				expectedHex, encodedString)
		}
	})

	t.Run("Unmarshall decodes bytes to Header", func(t *testing.T) {
		b := []byte{1, 255, 147, 128, 0, 1, 0, 1, 0, 1, 0, 1}
		h := message.Header{}
		err := h.Unmarshall(b)
		if err != nil {
			t.Fatalf("Unmarshall failed: %s", err.Error())
		}

		tt := []struct {
			value   any
			expects any
			error   string
		}{
			{
				value:   h.ID,
				expects: uint16(511),
				error:   "ID value incorrectly converted",
			},
			{
				value:   h.QR,
				expects: byte(1),
				error:   "QR value incorrectly converted ",
			},
			{
				value:   h.OPCode,
				expects: message.OPCODE(2),
				error:   "Opcode value incorrectly converted ",
			},
			{
				value:   h.AA,
				expects: byte(0),
				error:   "AA value incorrectly converted ",
			},
			{
				value:   h.TC,
				expects: byte(1),
				error:   "TC value incorrectly converted ",
			},
			{
				value:   h.RD,
				expects: byte(1),
				error:   "RD value incorrectly converted ",
			},
			{
				value:   h.RA,
				expects: byte(1),
				error:   "RA value incorrectly converted ",
			},
			{
				value:   h.Z,
				expects: byte(0),
				error:   "Z value incorrectly converted ",
			},
			{
				value:   h.RCode,
				expects: message.RCODE(0),
				error:   "RCode value incorrectly converted ",
			},
			{
				value:   h.QDCOUNT,
				expects: uint16(1),
				error:   "QDCOUNT value incorrectly converted ",
			},
			{
				value:   h.ANCOUNT,
				expects: uint16(1),
				error:   "ANCOUNT value incorrectly converted ",
			},
			{
				value:   h.NSCOUNT,
				expects: uint16(1),
				error:   "NSCOUNT value incorrectly converted ",
			},
			{
				value:   h.ARCOUNT,
				expects: uint16(1),
				error:   "ARCOUNT value incorrectly converted ",
			},
		}

		for _, te := range tt {
			if te.value != te.expects {
				fmt.Println(h)
				t.Fatalf("%s", te.error)
			}
		}
	})
}
