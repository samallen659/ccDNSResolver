package message_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/samallen659/ccDNSResolver/internal/message"
)

func TestHeader(t *testing.T) {
	t.Run("Test Marshall", func(t *testing.T) {
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
}
