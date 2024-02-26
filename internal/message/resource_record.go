package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type TYPE uint16

const (
	TYPE_A     TYPE = 1
	TYPE_NS    TYPE = 2
	TYPE_MD    TYPE = 3
	TYPE_MF    TYPE = 4
	TYPE_CNAME TYPE = 5
	TYPE_SOA   TYPE = 6
	TYPE_MB    TYPE = 7
	TYPE_MG    TYPE = 8
	TYPE_MR    TYPE = 9
	TYPE_NULL  TYPE = 10
	TYPE_WKS   TYPE = 11
	TYPE_PTR   TYPE = 12
	TYPE_HINFO TYPE = 13
	TYPE_MINFO TYPE = 14
	TYPE_MX    TYPE = 15
	TYPE_TXT   TYPE = 16
)

type CLASS uint16

const (
	CLASS_IN CLASS = 1
	CLASS_CS CLASS = 2
	CLASS_CH CLASS = 3
	CLASS_HS CLASS = 4
)

type ResourceRecord struct {
	// The domain name this resorce record pertains
	Name string
	// Contains one of the RR type codes. This field
	// specifies the meaning of the data in RData
	Type TYPE

	// Specifies the class of the data in RData
	Class CLASS

	// Specifies the time interval (in seconds) that the
	// resource record may be cached before it should be
	// discarded. Zero values mean it should not be cached
	TTL uint32

	// Specifies the length in bytes of RData
	RDLength uint16

	// Describes the resource depending of the Type and Class
	// fields
	RData []byte
}

// Takes in a pointer to a full DNS response in b, and an
// int s that denotes the starting byte of the resource records and c that denotes
// the number of resource records in the response. If formatted correctly returns
// a slice of parsed resource records, otherwise returns an error
func ParseResourceRecords(b *[]byte, s int, c int) ([]*ResourceRecord, error) {
	var rrs []*ResourceRecord
	fmt.Println((*b)[s : s+10])
	pos := s

	for i := 0; i < c; i++ {
		var rr ResourceRecord
		name, offset, err := parseResourceRecordName(b, pos)
		if err != nil {
			return nil, err
		}
		rr.Name = name
		pos += offset

		rr.Type = TYPE(binary.BigEndian.Uint16((*b)[pos : pos+2]))
		pos += 2

		rr.Class = CLASS(binary.BigEndian.Uint16((*b)[pos : pos+2]))
		pos += 2

		rr.TTL = binary.BigEndian.Uint32((*b)[pos : pos+4])
		pos += 4

		rr.RDLength = binary.BigEndian.Uint16((*b)[pos : pos+2])
		pos += 2

		rr.RData = (*b)[pos : pos+int(rr.RDLength)]
		pos += int(rr.RDLength)
		rrs = append(rrs, &rr)
	}

	return rrs, nil
}

// Takes in a pointer to a full DNS response in b, and an int s that denotes
// the starting byte of the resource record. Parses the name into a string
// format checking for compression. Compression format described in RFC 1035
// https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.4. Returns
// the byte slice of the name and an int for the offset in the position that should
// now be checked in the message
func parseResourceRecordName(b *[]byte, s int) (string, int, error) {
	var n bytes.Buffer
	var offset int
	pointer := false
	for (*b)[s+offset] != 0 {
		// if a pointer
		if (*b)[s+offset]&192 == 192 {
			pointer = true
			// pulls pointer value while remove pointer flag bits
			ps := uint16(49152) ^ binary.BigEndian.Uint16((*b)[s+offset:s+offset+2])
			pn := 0
			for i := int(ps); i < len((*b)); i++ {
				if (*b)[i] == byte(0) {
					pn = i
					break
				}
			}
			n.Write((*b)[ps : pn+1])
			offset += 2
			// compression only allows for names to end with a pointer
			break
		}

		c := int((*b)[s+offset])
		n.Write((*b)[s+offset : s+offset+c+1])
		offset += c + 1
	}

	//to account for zero byte in a label format
	if !pointer {
		offset++
	}

	name := ConvertQNameToHostname(n.Bytes())

	return name, offset, nil
}
