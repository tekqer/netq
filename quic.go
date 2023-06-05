package netq

import (
	"context"
	"io"
	"net"
)

type QUICListener struct {
}

func (l *QUICListener) Accept(ctx context.Context) (Conn, error) {
	return nil, nil
}

func (l *QUICListener) Addr() Addr {
	return nil
}

func (l *QUICListener) Close() error {
	return nil
}

type QUICConn struct {
	net.TCPConn
}

func (c *QUICConn) Open(ctx context.Context) (Stream, error) {
	return nil, nil
}

func (c *QUICConn) Accept(ctx context.Context) (Stream, error) {
	return nil, nil
}

func (c *QUICConn) Alive() bool {
	return false
}

type QUICStream struct {
	io.ReadWriteCloser
}

type QUICAddr struct{}

func (a QUICAddr) Network() string {
	return "QUIC"
}

func (a QUICAddr) String() string {
	return ""
}

func DialQUIC(ctx context.Context, addr string) (Conn, error) {
	return nil, nil
}

func ListenQUIC(addr string) (*QUICListener, error) {
	return nil, nil
}
