package connection

import (
	"errors"
	"fmt"
	"net"
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
