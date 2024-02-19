package message

type Message struct {
	header   Header
	question Question
}

func (m *Message) Marshall() []byte {
	hBytes, _ := m.header.Marshall()
	qBytes := m.question.Marshall()
	var b []byte
	b = append(b, hBytes...)
	b = append(b, qBytes...)
	return b
}
