package httptest

import (
	"bufio"
	"testing"
	"time"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func AssertView(t *testing.T, req *fasthttp.Request, fnView atreugo.View, assertFn func(resp *fasthttp.Response)) {
	s := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			actx := atreugo.AcquireRequestCtx(ctx)
			fnView(actx) // nolint:errcheck
			atreugo.ReleaseRequestCtx(actx)
		},
	}

	conn := new(MockConn)

	if len(req.Header.Host()) == 0 {
		req.Header.SetHost("http-server.test")
	}

	if _, err := req.WriteTo(&conn.RBuff); err != nil {
		panic(err)
	}

	ch := make(chan error, 1)

	go func() {
		ch <- s.ServeConn(conn)
	}()

	select {
	case err := <-ch:
		if err != nil {
			t.Fatalf("Serve connection error: %v", err)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Serve connection timeout")
	}

	br := bufio.NewReader(&conn.WBuff)
	resp := new(fasthttp.Response)

	if err := resp.Read(br); err != nil {
		panic(err)
	}

	assertFn(resp)
}
