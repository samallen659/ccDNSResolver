package message

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
