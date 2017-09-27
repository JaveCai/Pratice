package md5transport

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Transport struct {
	http.RoundTripper
}

func NewTransport(transport http.RoundTripper) http.RoundTripper {
	return &Transport{transport}
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {

	// 补全代码
	if r.Body == nil {
		return t.RoundTripper.RoundTrip(r)
	}
	b, err := ioutil.ReadAll(r.Body)
	//r.Body.Close()
	//ast.Nil(err)
	if err != nil {
		fmt.Println("[RoundTrip] ioutil read fail")
	} else {
		if len(b) == 0 {
			md5exp := ""
			r.Header.Set("X-Md5", md5exp)
		} else {
			hexB := md5.Sum(b)
			md5exp := hex.EncodeToString(hexB[:])
			r.Header.Set("X-Md5", md5exp)

			//bug fixed: 由于 ioutil.ReadAll 方法会读取到 EOF，所以需要重置 Body
			r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		}

	}

	return t.RoundTripper.RoundTrip(r)
}
