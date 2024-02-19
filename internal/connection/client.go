package connection

import (
	"errors"
	"fmt"
	"github.com/samallen659/ccDNSResolver/internal/message"
	"net"
	"time"
)

type Client struct {
	addr *net.UDPAddr
}

func NewClient(addr string) (*Client, error) {
	if addr == "" {
		return nil, errors.New("addr string cannot be empty")
	}

	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr, 53))
	if err != nil {
		return nil, fmt.Errorf("Failed to resolve address: %s", err.Error())
	}

	return &Client{addr: udpAddr}, nil
}

func (c *Client) Resolve(hostname string) error {

	id := message.NewHeaderID()
	m := message.Message{
		Header: message.Header{
			ID:      id,
			QR:      byte(0),
			OPCode:  message.OPCODEQuery,
			RD:      byte(1),
			QDCOUNT: uint16(1),
		},
		Question: message.Question{
			QName:  message.ConvertHostnameToQName(hostname),
			QType:  message.QTYPE_A,
			QClass: message.QCLASS_IN,
		}}

	mBytes := m.Marshall()
	fmt.Println(mBytes)

	fmt.Println("setting up connection")
	conn, err := net.DialUDP("udp", nil, c.addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	_, err = conn.Write(mBytes)
	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return err
	}

	fmt.Println(buf)

	return nil
}
