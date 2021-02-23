package dingding

import "fmt"

type resultError interface {
	HasError() bool
	ErrorMsg() string
}

type apiResultBase struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (a *apiResultBase) HasError() bool {
	return a.Errcode != 0
}

func (a *apiResultBase) ErrorMsg() string {
	if !a.HasError() {
		return ""
	}

	msg := DingdingErrorSet[a.Errcode]
	return fmt.Sprintf("%d:%s", a.Errcode, msg)
}

type AccessToken struct {
	apiResultBase
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type TextMessageResponse struct {
	apiResultBase
	RequestID string `json:"request_id"`
	TaskID    int64  `json:"task_id"`
}

type textInner struct {
	Content string `json:"content"`
}

type messageInner struct {
	MsgType string    `json:"msgtype"`
	Text    textInner `json:"text"`
}

type TextMessageRequest struct {
	AgentID    int64        `json:"agent_id"`
	UseridList string       `json:"userid_list"`
	Msg        messageInner `json:"msg"`
}
