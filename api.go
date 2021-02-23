package dingding

import (
	"strconv"
	"strings"
)

const (
	URLAccessToken            = "https://oapi.dingtalk.com/gettoken"
	URLSendMessage            = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2"
	URLDepartmentList         = "https://oapi.dingtalk.com/department/list"
	URLDepartmentMemberByPage = "https://oapi.dingtalk.com/user/listbypage"
)

type Api struct {
	atm    AccessTokenManager
	client HttpClientor

	// 钉钉配置
	cfg *Config
}

func (a *Api) AccessToken() (string, error) {
	token, err := a.atm.Get(a.cfg.agentId)
	if err != nil {
		if err == ErrTokenExpired {
			at, err := a.doAccessToken()
			if err != nil {
				return "", err
			}
			a.atm.Set(a.cfg.agentId, at)
			return at.AccessToken, nil
		}
		return "", err
	}

	return token, nil
}

func (a *Api) doAccessToken() (*AccessToken, error) {
	u, _ := URLParse(URLAccessToken, "appkey", a.cfg.appKey, "appsecret", a.cfg.appSecret)

	var result = new(AccessToken)
	_, err := httpGet(u, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *Api) SendTextMessage(msg string, toUserIDs []string) error {
	accessToken, err := a.AccessToken()
	if err != nil {
		return err
	}

	u, _ := URLParse(URLSendMessage, "access_token", accessToken)
	req := &TextMessageRequest{
		AgentID:    a.cfg.agentId,
		UseridList: strings.Join(toUserIDs, ","),
		Msg: messageInner{
			MsgType: "text",
			Text: textInner{
				Content: msg,
			},
		},
	}

	textMessageResp := new(TextMessageResponse)
	_, err = httpPost(u, req, textMessageResp)
	return err
}

func (a *Api) DepartmentList(rootDepID int64) (*DepartmentResponse, error) {
	if rootDepID <= 0 {
		rootDepID = 1
	}

	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}
	u, _ := URLParse(URLDepartmentList, "access_token", accessToken, "id", strconv.Itoa(int(rootDepID)))

	depResponse := new(DepartmentResponse)
	_, err = httpGet(u, depResponse)
	return depResponse, err
}
