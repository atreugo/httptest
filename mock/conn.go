package mock

import (
	"net"
	"time"
)

var zeroTCPAddr = &net.TCPAddr{
	IP: net.IPv4zero,
}

func (conn *MockConn) Close() error {
	return conn.ErrClose
}

func (conn *MockConn) Read(b []byte) (int, error) {
	if conn.ErrRead != nil {
		return 0, conn.ErrRead
	}

	return conn.RBuff.Read(b)
}

func (conn *MockConn) Write(b []byte) (int, error) {
	if conn.ErrWrite != nil {
		return 0, conn.ErrWrite
	}

	return conn.WBuff.Write(b)
}

func (conn *MockConn) RemoteAddr() net.Addr {
	return zeroTCPAddr
}

func (conn *MockConn) LocalAddr() net.Addr {
	return zeroTCPAddr
}

func (conn *MockConn) SetDeadline(t time.Time) error {
	return conn.ErrSetDeadline
}

func (conn *MockConn) SetReadDeadline(t time.Time) error {
	return conn.ErrSetReadDeadline
}

func (conn *MockConn) SetWriteDeadline(t time.Time) error {
	return conn.ErrSetWriteDeadline
}
