package dingding

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

var ErrInvalidEventType = errors.New("invalid event type.")

type EncryptBody struct {
	MsgSignature string `json:"msg_signature"`
	TimeStamp    string `json:"timeStamp"`
	Nonce        string `json:"nonce"`
	Encrypt      string `json:"encrypt"`
}

type EventBody struct {
	EventType string `json:"EventType"`
}

type EventProcessor interface {
	Type() string
	Process(decryptMsg []byte) error
}

type EventHub struct {
	appKey, aesKey, token string

	hub map[string]EventProcessor
}

func NewEventHub(appKey, aesKey, token string) *EventHub {
	hub := &EventHub{
		appKey: appKey,
		aesKey: aesKey,
		token:  token,
		hub:    map[string]EventProcessor{},
	}

	checkURL := &EventCheckURL{
		appKey: appKey,
		aesKey: aesKey,
		token:  token,
	}
	hub.Register(checkURL)
	return hub
}

func (e *EventHub) Register(p EventProcessor) {
	e.hub[p.Type()] = p
}

func (e *EventHub) Do(requestBody io.Reader) (interface{}, error) {
	var encryptBody = new(EncryptBody)
	err := json.NewDecoder(requestBody).Decode(encryptBody)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	signature := encryptBody.MsgSignature
	timestamp := encryptBody.TimeStamp
	nonce := encryptBody.Nonce

	decryptMsg, err := NewDingTalkCrypto(e.token, e.aesKey, e.appKey).
		GetDecryptMsg(signature, timestamp, nonce, encryptBody.Encrypt)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body := new(EventBody)
	err = json.Unmarshal([]byte(decryptMsg), body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if p, ok := e.hub[body.EventType]; ok {
		err = p.Process([]byte(decryptMsg))
		if err != nil {
			return nil, err
		} else {
			// 返回 success
			return NewDingTalkCrypto(e.token, e.aesKey, e.appKey).
				GetDecryptMsg(signature, timestamp, nonce, "success")
		}
	}
	return nil, ErrInvalidEventType
}
