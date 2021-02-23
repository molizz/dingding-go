package dingding

type eventCheckURL struct {
}

var _ EventProcessor = (*EventCheckURL)(nil)

type EventCheckURL struct {
	appKey, aesKey, token string
}

func (e *EventCheckURL) Type() string {
	return "check_url"
}

func (e *EventCheckURL) Process(session *EventSession, decryptEventBody []byte) (interface{}, error) {
	en := session.encryptBody
	return NewDingTalkCrypto(e.token, e.aesKey, e.appKey).GetDecryptMsg(en.MsgSignature, en.TimeStamp, en.Nonce, "success")
}
