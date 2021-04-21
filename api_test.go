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
var AppID = os.Getenv("DINGDING_APPID")
var AppSecret = os.Getenv("DINGDING_APPSECRET")

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

func TestApi_DepartmentWithID(t *testing.T) {
	members, err := testDingding.DepartmentAllMembers(1)
	assert.Equal(t, nil, err)
	assert.Greater(t, len(members), 1)
}

func TestApi_FetchAllMembersByDepartmentID(t *testing.T) {
	resp, err := testDingding.FetchAllMembersByDepartmentID(1)
	assert.Equal(t, nil, err)
	assert.Greater(t, len(resp.UserList), 1)
}

func TestApi_GetUserInfoByCode(t *testing.T) {
	got, err := testDingding.GetUserInfoByCode(
		"code",
		AppID,
		AppSecret,
	)

	assert.Equal(t, nil, err)
	assert.NotEmpty(t, got)
}

func TestApi_UserIDByUnionID(t *testing.T) {
	got, err := testDingding.UserIDByUnionID("fUqL5R7MCN2iiWgMbPolsngiEiE")
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, got)
}

func TestApi_SendTextMessage(t *testing.T) {
	err := testDingding.SendTextMessage("test", []string{"manager9961"})
	assert.Equal(t, nil, err)
}
