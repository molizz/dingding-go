package dingding

import (
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
	var body io.Reader
	if in != nil {
		body = in.(io.Reader)
	}
	req, err := http.NewRequest("POST", uri.String(), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := doHttp(req, in, out)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

func httpGet(uri *url.URL, out interface{}) (*http.Response, error) {
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := doHttp(req, nil, out)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return resp, nil
}

func doHttp(req *http.Request, in, out interface{}) (*http.Response, error) {
	resp, err := DefaultHttpClient.Do(req)
	if err != nil {
		return nil, err
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
