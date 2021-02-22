package dingding

import "net/url"

func URLParse(uri string, querys ...string) (*url.URL, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	values := make(url.Values)

	for i := 0; i < len(querys)-1; i += 2 {
		values[querys[i]] = []string{querys[i+1]}
	}
	u.RawQuery = values.Encode()
	return u, nil
}
