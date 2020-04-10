package httptest

import "bytes"

type mockConn struct {
	r bytes.Buffer
	w bytes.Buffer
}
