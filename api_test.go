package dingding

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDingding *Dingding
var testCfg = &Config{
	agentId:   os.Getenv("DINGDING_AGENTID"),
	appKey:    os.Getenv("DINGDING_KEY"),
	appSecret: os.Getenv("DINGDING_SECRET"),
}

func init() {
	testDingding = New(testCfg, &DefaultAccessTokenManager{})
}

func TestApi_AccessToken(t *testing.T) {
	at, err := testDingding.AccessToken()
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, at)
}
