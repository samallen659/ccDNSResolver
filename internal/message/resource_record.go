package message

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
	Name []byte

	// Contains one of the RR type codes. This field
	// specifies the meaning of the data in RData
	Type uint16

	// Specifies the class of the data in RData
	Class uint16

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

func ParseResourceRecords(b []byte) ([]ResourceRecord, error) {
	return nil, nil
}
