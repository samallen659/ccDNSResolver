package message

import (
	"bytes"
)

type Message struct {
	Header          Header
	Question        Question
	ResourceRecords []*ResourceRecord
}

func (m *Message) Marshall() []byte {
	var b bytes.Buffer
	b.Write(m.Header.Marshall())
	b.Write(m.Question.Marshall())
	return b.Bytes()
}

func (m *Message) Unmarshall(b []byte) error {
	h := Header{}
	if err := h.Unmarshall(b[:12]); err != nil {
		return err
	}
	m.Header = h

	pos := 12
	//find zero byte ending of QNAME for each Question
	for b[pos] != 0 {
		pos++
	}
	//to account for QTYPE & QCLASS
	pos += 4

	q := Question{}
	if err := q.Unmarshall(b[12 : pos+1]); err != nil {
		return err
	}
	m.Question = q
	pos++

	rrs, err := ParseResourceRecords(&b, pos, int(h.ANCOUNT))
	if err != nil {
		return err
	}

	m.ResourceRecords = rrs

	return nil
}
