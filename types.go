package httptest

import (
	"bytes"
	"net"
)

type MockConn struct {
	RBuff bytes.Buffer
	WBuff bytes.Buffer

	ErrClose            error
	ErrRead             error
	ErrWrite            error
	ErrSetDeadline      error
	ErrSetReadDeadline  error
	ErrSetWriteDeadline error

	net.Conn
}

type MockListener struct {
	LN net.Listener

	AcceptError error
	CloseError  error

	AcceptCalled bool
	CloseCalled  bool
	AddrCalled   bool
}
