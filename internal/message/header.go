package message

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"math/rand"
)

type OPCODE byte

const (
	OPCODEQuery OPCODE = iota
	OPCODEIQuery
	OPCODEStatus
)

type RCODE byte

const (
	RCODENoError RCODE = iota
	RCODEFormatError
	RCODEServerFailure
	RCODENameError
	RCODENotImplemented
	RCODERefused
)

// Header follows the format specifid in RFC 1035 https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.1
//
//					             0  1  2  3  4  5  6  7
//	     0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F
//	   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	   |                      ID                       |
//	   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	   |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
//	   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	   |                    QDCOUNT                    |
//	   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	   |                    ANCOUNT                    |
//	   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	   |                    NSCOUNT                    |
//	   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	   |                    ARCOUNT                    |
//	   +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
type Header struct {
	// A 16 bit identifier assigned by the program that
	// generates any kind of query.  This identifier is copied
	// the corresponding reply and can be used by the requester
	// to match up replies to outstanding queries.
	ID uint16

	// A one bit field that specifies whether this message is a
	// query (0), or a response (1).
	// 1bit
	QR byte

	// A four bit field that specifies kind of query in this
	// message.  This value is set by the originator of a query
	// and copied into the response.  The values are:
	// 0               - a standard query (QUERY)
	// 1               - an inverse query (IQUERY)
	// 2               - a server status request (STATUS)
	// 3 to 15         - reserved for future use
	OPCode OPCODE

	// Authoritative Answer - this bit is valid in responses,
	// and specifies that the responding name server is an
	// authority for the domain name in question section.
	// 1bit
	AA byte

	// TrunCation - specifies that this message was truncated
	// due to length greater than that permitted on the
	// transmission channel.
	TC byte

	// Recursion Desired - If set, it directs the name server to
	// pursue the query recursively
	RD byte

	// Recursion Available - this is set or cleared in a response,
	// and denotes whether recursive query support is available
	RA byte

	// Reserved for future use
	Z byte

	// Response code - this 4 bit field is set as part of
	// responses.  The values have the following
	// interpretation:
	// 0               - No error condition
	// 1               - Format error - The name server was
	//                 unable to interpret the query.
	// 2               - Server failure - The name server was
	//                 unable to process this query due to a
	//                 problem with the name server.
	// 3               - Name Error - Meaningful only for
	//                 responses from an authoritative name
	//                 server, this code signifies that the
	//                 domain name referenced in the query does
	//                 not exist.
	// 4               - Not Implemented - The name server does
	//                 not support the requested kind of query.
	// 5               - Refused - The name server refuses to
	//                 perform the specified operation for
	//                 policy reasons.  For example, a name
	//                 server may not wish to provide the
	//                 information to the particular requester,
	//                 or a name server may not wish to perform
	//                 a particular operation (e.g., zone
	RCode RCODE

	// An unsigned 16 bit integer specifying the number of entries
	// int the question section.
	QDCOUNT uint16

	// An unsigned 16 bit integer specifying the number of resource
	// records in the answer section.
	ANCOUNT uint16

	// An unsigned 16 bit integer specifying the number of name
	// server resource records in the authroity records section.
	NSCOUNT uint16

	// An unsigned 16 bit integer specifying the number of resource
	// records in the additional records section.
	ARCOUNT uint16
}

func (h *Header) Marshall() []byte {
	var b bytes.Buffer
	var h1 byte
	var h2 byte

	h1 |= h.QR << 7
	h1 |= byte(h.OPCode) << 3
	h1 |= h.AA << 2
	h1 |= h.TC << 1
	h1 |= h.RD

	h2 |= h.RA << 7
	h2 |= h.Z << 4
	h2 |= byte(h.RCode)

	binary.Write(&b, binary.BigEndian, h.ID)
	binary.Write(&b, binary.BigEndian, h1)
	binary.Write(&b, binary.BigEndian, h2)
	binary.Write(&b, binary.BigEndian, h.QDCOUNT)
	binary.Write(&b, binary.BigEndian, h.ANCOUNT)
	binary.Write(&b, binary.BigEndian, h.NSCOUNT)
	binary.Write(&b, binary.BigEndian, h.ARCOUNT)

	return b.Bytes()
}

func (h *Header) Unmarshall(b []byte) error {
	if len(b) != 12 {
		return errors.New("Incorrect length byte slice. Expected 12")
	}

	h.ID = binary.BigEndian.Uint16(b[:2])

	h1 := b[2]

	h.QR = (h1 >> 7) & 1
	h.OPCode = OPCODE((h1 >> 3) & 15)
	h.AA = (h1 >> 2) & 1
	h.TC = (h1 >> 1) & 1
	h.RD = h1 & 1

	h2 := b[3]

	h.RA = (h2 >> 7) & 1
	h.Z = (h2 >> 4) & 7
	h.RCode = RCODE(h2 & 15)

	h.QDCOUNT = binary.BigEndian.Uint16(b[4:6])
	h.ANCOUNT = binary.BigEndian.Uint16(b[6:8])
	h.NSCOUNT = binary.BigEndian.Uint16(b[8:10])
	h.ARCOUNT = binary.BigEndian.Uint16(b[10:12])

	return nil
}

func NewHeaderID() uint16 {
	return uint16(rand.Intn(math.MaxUint16 + 1))
}
