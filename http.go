package dingding

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

var (
	DefaultHttpClient HttpClientor
	HTTPTimeout       = 10 * time.Second // 10s
)

func getHttpClient() HttpClientor {
	if DefaultHttpClient == nil {
		c := &http.Client{}
		c.Timeout = HTTPTimeout

		DefaultHttpClient = c
	}
	return DefaultHttpClient
}

var _ HttpClientor = (*http.Client)(nil)

// for http mock
type HttpClientor interface {
	Do(req *http.Request) (*http.Response, error)
}

func httpPost(uri *url.URL, in, out interface{}) (*http.Response, error) {
	resp, err := doHttp("POST", uri, in, out)
	return resp, errors.WithStack(err)
}

func httpGet(uri *url.URL, out interface{}) (*http.Response, error) {
	resp, err := doHttp("GET", uri, nil, out)
	return resp, errors.WithStack(err)
}

func doHttp(method string, uri *url.URL, in, out interface{}) (*http.Response, error) {
	var body io.Reader
	if in != nil {
		if bodyReader, ok := in.(io.Reader); ok {
			body = bodyReader
		} else {
			bodyBuff := bytes.NewBufferString("")
			err := json.NewEncoder(bodyBuff).Encode(in)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			body = bodyBuff
		}
	}

	req, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := DefaultHttpClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if out != nil {
		if w, ok := out.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(out)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}
	}

	if e, ok := out.(resultError); ok && e.HasError() {
		return nil, errors.Errorf(e.ErrorMsg())
	}

	return resp, nil
}
