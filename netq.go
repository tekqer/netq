package netq

import (
	"context"
	"errors"
	"io"
	"net"
)

type Listener interface {
	// Accept blocks until receiving a connection from the transport layer.
	Accept(context.Context) (Conn, error)
	// Addr returns the address bound to the listener.
	Addr() Addr
	// Close closes the listener.
	Close() error
}

type Conn interface {
	net.Conn
	// Open opens a stream from the connection.
	Open(context.Context) (Stream, error)
	// Accept blocks until receiving a stream from the connection.
	Accept(context.Context) (Stream, error)
	// Alive indicates if the connection is alive.
	Alive() bool
}

type Stream interface {
	io.ReadWriteCloser
}

type Addr = net.Addr

// Dial connects to an address on the named network that must be
// `tcpx` or `quic`.
// For streams over TCP, use `tcpx`.
// For streams over UDP, use `quic`.
func Dial(network, addr string) (Conn, error) {
	return dialContext(context.Background(), network, addr)
}

func dialContext(ctx context.Context, network, addr string) (Conn, error) {
	switch network {
	case "tcpx":
		return DialTCPX(ctx, addr)
	case "quic":
		return DialQUIC(ctx, addr)
	}
	return nil, errors.New("bad network name")
}

// Listen announces on a local network address. Network must be
// `tcpx` or `quic`.
// For streams over TCP, use `tcpx`.
// For streams over UDP, use `quic`.
func Listen(network, addr string) (Listener, error) {
	switch network {
	case "tcpx":
		return ListenTCPX(addr)
	case "quic":
		return ListenQUIC(addr)
	}
	return nil, errors.New("bad network name")
}
