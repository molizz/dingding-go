package dingding

import "fmt"

type resultError interface {
	HasError() bool
	ErrorMsg() string
}

type apiResultBase struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	ExpiresIn int    `json:"expires_in"`
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
	AccessToken string `json:"access_token"`
}
