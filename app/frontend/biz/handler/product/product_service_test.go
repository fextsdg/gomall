package product

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetProductById(t *testing.T) {
	h := server.Default()
	h.GET("/product", GetProductById)
	path := "/product"                                        // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}

func TestSearchProduct(t *testing.T) {
	h := server.Default()
	h.GET("/search", SearchProduct)
	path := "/search"                                         // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	// assert.DeepEqual(t, 200, resp.StatusCode())
	// assert.DeepEqual(t, "null", string(resp.Body()))
}
