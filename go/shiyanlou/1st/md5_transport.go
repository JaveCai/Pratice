package md5transport

import (
    "net/http"
)

type Transport struct {
    http.RoundTripper
}

func NewTransport(transport http.RoundTripper) http.RoundTripper {
    return &Transport{transport}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {

    // 补全代码

    return t.RoundTripper.RoundTrip(req)
}