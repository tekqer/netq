package netq

import (
	"context"
	"io"
	"net"
)

type TCPXListener struct {
}

func (l *TCPXListener) Accept(ctx context.Context) (Conn, error) {
	return nil, nil
}

func (l *TCPXListener) Addr() Addr {
	return nil
}

func (l *TCPXListener) Close() error {
	return nil
}

type TCPXConn struct {
	net.TCPConn
}

func (c *TCPXConn) Open(ctx context.Context) (Stream, error) {
	return nil, nil
}

func (c *TCPXConn) Accept(ctx context.Context) (Stream, error) {
	return nil, nil
}

func (c *TCPXConn) Alive() bool {
	return false
}

type TCPXStream struct {
	io.ReadWriteCloser
}

type TCPXAddr struct{}

func (a TCPXAddr) Network() string {
	return "tcpx"
}

func (a TCPXAddr) String() string {
	return ""
}

func DialTCPX(ctx context.Context, addr string) (Conn, error) {
	return nil, nil
}

func ListenTCPX(addr string) (*TCPXListener, error) {
	return nil, nil
}
