package dingding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLParse(t *testing.T) {
	uri := "https://oapi.dingtalk.com/gettoken"

	str, err := URLParse(uri, []string{"a", "1", "b", "2"}...)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://oapi.dingtalk.com/gettoken?a=1&b=2", str)
}
