package dingding

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDingding *Dingding
var testAgentID, _ = strconv.Atoi(os.Getenv("DINGDING_AGENTID"))
var testCfg = &Config{
	agentId:   int64(testAgentID),
	appKey:    os.Getenv("DINGDING_KEY"),
	appSecret: os.Getenv("DINGDING_SECRET"),
}

func init() {
	mg := NewDefaultAccessTokenManager()
	testDingding = New(testCfg, mg)
}

func TestApi_AccessToken(t *testing.T) {
	at, err := testDingding.AccessToken()
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, at)
}
