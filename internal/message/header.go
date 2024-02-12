package message

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
	// 0               a standard query (QUERY)
	// 1               an inverse query (IQUERY)
	// 2               a server status request (STATUS)
	// 3 to 15         reserved for future use
	OPCODE OPCODE

	// Authoritative Answer - this bit is valid in responses,
	// and specifies that the responding name server is an
	// authority for the domain name in question section.
	// 1bit
	AA byte

	// TrunCation - specifies that this message was truncated
	// due to length greater than that permitted on the
	// transmission channel.
	TC byte

	RD      byte
	Z       byte
	RC      RCODE
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}
