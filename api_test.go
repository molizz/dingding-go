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
	AgentId:   int64(testAgentID),
	AppKey:    os.Getenv("DINGDING_KEY"),
	AppSecret: os.Getenv("DINGDING_SECRET"),
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

func TestApi_DepartmentList(t *testing.T) {
	resp, err := testDingding.DepartmentList(0)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(resp.Departments))
}
