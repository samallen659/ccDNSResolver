package message

import "bytes"

type Message struct {
	Header   Header
	Question Question
}

func (m *Message) Marshall() []byte {
	var b bytes.Buffer
	b.Write(m.Header.Marshall())
	b.Write(m.Question.Marshall())
	return b.Bytes()
}
