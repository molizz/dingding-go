package dingding

import "fmt"

type resultError interface {
	HasError() bool
	ErrorMsg() string
}

type ApiResultBase struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (a *ApiResultBase) HasError() bool {
	return a.Errcode != 0
}

func (a *ApiResultBase) ErrorMsg() string {
	if !a.HasError() {
		return ""
	}

	msg := DingdingErrorSet[a.Errcode]
	return fmt.Sprintf("%d:%s", a.Errcode, msg)
}

type AccessToken struct {
	ApiResultBase
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type TextMessageResponse struct {
	ApiResultBase
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

type departmentInner struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	ParentID        int    `json:"parentid"`
	CreateDeptGroup bool   `json:"createDeptGroup"`
	AutoAddUser     bool   `json:"autoAddUser"`
}
type DepartmentsResponse struct {
	ApiResultBase
	Departments []*departmentInner `json:"department"`
}

type DepartmentResponse struct {
	ApiResultBase
	departmentInner
	Order int `json:"order"`
}

type DepartmentMemeberCountResponse struct {
	ApiResultBase
	UserIds []string `json:"userIds"`
}

type memberInner struct {
	UserID      string `json:"userid"`
	Name        string `json:"name"`
	Departments []int  `json:"department"`
	UnionID     string `json:"unionid"`
	Email       string `json:"email"`
}
type DepartmentMemberResponse struct {
	ApiResultBase
	UserList []*memberInner `json:"userlist"`
}

type UserIDResponse struct {
	ApiResultBase
	UserID string `json:"userid"`
}

type UserResponse struct {
	ApiResultBase

	UserID      string `json:"userid"`
	Name        string `json:"name"`
	Departments []int  `json:"department"`
	UnionID     string `json:"unionid"`
	Email       string `json:"email"`
}

type authOrgScopesInner struct {
	AuthedUser []string `json:"authed_user"`
	AuthedDept []int    `json:"authed_dept"`
}
type AuthScopesResponse struct {
	ApiResultBase
	AuthOrgScopes *authOrgScopesInner `json:"auth_org_scopes"`
	AuthUserField []string            `json:"auth_user_field"`
}
