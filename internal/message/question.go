package message

import "strings"

type QTYPE uint16

const (
	QTYPE_A     QTYPE = 1
	QTYPE_NS    QTYPE = 2
	QTYPE_MD    QTYPE = 3
	QTYPE_MF    QTYPE = 4
	QTYPE_CNAME QTYPE = 5
	QTYPE_SOA   QTYPE = 6
	QTYPE_MB    QTYPE = 7
	QTYPE_MG    QTYPE = 8
	QTYPE_MR    QTYPE = 9
	QTYPE_NULL  QTYPE = 10
	QTYPE_WKS   QTYPE = 11
	QTYPE_PTR   QTYPE = 12
	QTYPE_HINFO QTYPE = 13
	QTYPE_MINFO QTYPE = 14
	QTYPE_MX    QTYPE = 15
	QTYPE_TXT   QTYPE = 16
	QTYPE_AXFR  QTYPE = 252
	QTYPE_MAILB QTYPE = 253
	QTYPE_MAILA QTYPE = 254
	QTYPE_ALL   QTYPE = 255
)

type QCLASS uint16

const (
	QCLASS_IN  QCLASS = 1
	QCLASS_CS  QCLASS = 2
	QCLASS_CH  QCLASS = 3
	QCLASS_HS  QCLASS = 4
	QCLASS_ANY QCLASS = 255
)

type Question struct {
	// A domain name represented as a sequence of labels, where each label consists
	// of a length byte followed by that number of bytes. The domain name terminates
	// with the zero length octet for the null label of the root.
	//
	// www.google.com becomes 3www6google3com0
	QName []byte

	// A 16 bit unsigned integer representation of the resource type
	// TYPE            value and meaning
	// A               1 a host address
	// NS              2 an authoritative name server
	// MD              3 a mail destination (Obsolete - use MX)
	// MF              4 a mail forwarder (Obsolete - use MX)
	// CNAME           5 the canonical name for an alias
	// SOA             6 marks the start of a zone of authority
	// MB              7 a mailbox domain name (EXPERIMENTAL)
	// MG              8 a mail group member (EXPERIMENTAL)
	// MR              9 a mail rename domain name (EXPERIMENTAL)
	// NULL            10 a null RR (EXPERIMENTAL)
	// WKS             11 a well known service description
	// PTR             12 a domain name pointer
	// HINFO           13 host information
	// MINFO           14 mailbox or mail list information
	// MX              15 mail exchange
	// TXT             16 text strings
	// AXFR            252 A request for a transfer of an entire zone
	// MAILB           253 A request for mailbox-related records (MB, MG or MR)
	// MAILA           254 A request for mail agent RRs (Obsolete - see MX)
	// *               255 A request for all records
	QType QTYPE

	// A 16 bit unsigned integer representing the class of the query
	// IN              1 the Internet
	// CS              2 the CSNET class (Obsolete - used only for examples in
	//                 some obsolete RFCs)
	// CH              3 the CHAOS class
	// HS              4 Hesiod [Dyer 87]
	// *               255 any class
	QClass QCLASS
}

func ConvertHostnameToQName(hostname string) []byte {
	var b []byte
	labels := strings.Split(hostname, ".")

	for _, label := range labels {
		lbs := []byte(label)
		length := len(lbs)
		b = append(b, byte(length))
		for _, lb := range lbs {
			b = append(b, lb)
		}
	}

	b = append(b, byte(0))

	return b
}
