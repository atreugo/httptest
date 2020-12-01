package assert

import (
	"testing"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func Test_Path(t *testing.T) {
	t.Parallel()

	body := "TEST"

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.SetRequestURI("/test")

	s := atreugo.New(atreugo.Config{})
	s.GET("/test", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse(body)
	})

	Server(t, req, s, func(resp *fasthttp.Response) {
		if string(resp.Body()) != body {
			t.Errorf("Response.Body == %s, want %s", resp.Body(), body)
		}
	})
}
