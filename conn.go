package httptest

import (
	"net"
	"time"
)

var zeroTCPAddr = &net.TCPAddr{
	IP: net.IPv4zero,
}

func (conn *mockConn) Close() error {
	return nil
}

func (conn *mockConn) Read(b []byte) (int, error) {
	return conn.r.Read(b)
}

func (conn *mockConn) Write(b []byte) (int, error) {
	return conn.w.Write(b)
}

func (conn *mockConn) RemoteAddr() net.Addr {
	return zeroTCPAddr
}

func (conn *mockConn) LocalAddr() net.Addr {
	return zeroTCPAddr
}

func (conn *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (conn *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (conn *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
