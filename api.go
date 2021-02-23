package dingding

import "strings"

const (
	URLAccessToken = "https://oapi.dingtalk.com/gettoken"
	URLSendMessage = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2"
)

type Api struct {
	atm    AccessTokenManager
	client HttpClientor

	// 钉钉配置
	cfg *Config
}

func (a *Api) AccessToken() (*AccessToken, error) {
	u, _ := URLParse(URLAccessToken, "appkey", a.cfg.appKey, "appsecret", a.cfg.appSecret)

	var result = new(AccessToken)
	_, err := httpGet(u, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *Api) SendTextMessage(msg string, toUserIDs []string) error {
	accessToken, err := a.atm.Get(a.cfg.agentId)
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
