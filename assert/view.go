package assert

import (
	"bufio"
	"testing"
	"time"

	"github.com/atreugo/mock"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func View(t *testing.T, req *fasthttp.Request, fnView atreugo.View, assertFn func(resp *fasthttp.Response)) {
	s := atreugo.New(atreugo.Config{})
	s.Path(string(req.Header.Method()), string(req.URI().PathOriginal()), fnView)

	if len(req.URI().Host()) == 0 {
		req.SetHost("http://http-server.test")
	}

	conn := new(mock.Conn)
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
