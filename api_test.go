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
	AesKey:    os.Getenv("DINGDING_AESKEY"),
	Token:     os.Getenv("DINGDING_TOKEN"),
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

func TestApi_GetUserInfoByCode(t *testing.T) {
	got, err := testDingding.GetUserInfoByCode(
		"04412de768a538329c64c75278b14df6",
		os.Getenv("DINGDING_ACCESSKEY"),
		os.Getenv("DINGDING_APPSECRET"),
	)

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, got)
}

func TestApi_UserIDByUnionID(t *testing.T) {
	got, err := testDingding.UserIDByUnionID("manager9961")
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, got)
}
