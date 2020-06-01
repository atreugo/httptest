package httptest

import "net"

func (m *MockListener) Accept() (net.Conn, error) {
	m.AcceptCalled = true

	if m.AcceptError != nil {
		return nil, m.AcceptError
	}

	return m.LN.Accept()
}

func (m *MockListener) Close() error {
	m.CloseCalled = true

	if m.CloseError != nil {
		return m.CloseError
	}

	return m.LN.Close()
}

func (m *MockListener) Addr() net.Addr {
	m.AddrCalled = true

	return m.LN.Addr()
}
