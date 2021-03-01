package dingding

import (
	"strconv"
	"strings"
	"time"
)

const (
	URLAccessToken            = "https://oapi.dingtalk.com/gettoken"
	URLSendMessage            = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2"
	URLDepartmentList         = "https://oapi.dingtalk.com/department/list"
	URLDepartmentMemberByPage = "https://oapi.dingtalk.com/user/listbypage"
	URLDepartment             = "https://oapi.dingtalk.com/department/get"
	URLDepartmentMemberIDs    = "https://oapi.dingtalk.com/user/getDeptMember"
	URLUserIDByUnionid        = "https://oapi.dingtalk.com/user/getUseridByUnionid"
	URLUserIDByMoblie         = "https://oapi.dingtalk.com/user/get_by_mobile"
	URLUserByID               = "https://oapi.dingtalk.com/user/get"
	URLAuthScopes             = "https://oapi.dingtalk.com/auth/scopes"
	URLDepartmentMembers      = "https://oapi.dingtalk.com/user/list" // NOTE 这个接口钉钉文档中已不存在，可能未来某个时候就无法使用了
	URLGetUserInfoByCode      = "https://oapi.dingtalk.com/sns/getuserinfo_bycode"
)

type Api struct {
	atm    AccessTokenManager
	client HttpClientor

	// 钉钉配置
	cfg *Config
}

func (a *Api) AccessToken() (string, error) {
	token, err := a.atm.Get(a.cfg.AgentId)
	if err != nil {
		if err == ErrTokenExpired {
			at, err := a.doAccessToken()
			if err != nil {
				return "", err
			}
			a.atm.Set(a.cfg.AgentId, at)
			return at.AccessToken, nil
		}
		return "", err
	}

	return token, nil
}

func (a *Api) doAccessToken() (*AccessToken, error) {
	u, _ := URLParse(URLAccessToken, "appkey", a.cfg.AppKey, "appsecret", a.cfg.AppSecret)

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
		AgentID:    a.cfg.AgentId,
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

// DepartmentList 获取部门列表
func (a *Api) DepartmentList(rootDepID int64) (*DepartmentsResponse, error) {
	if rootDepID <= 0 {
		rootDepID = 1
	}

	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLDepartmentList,
		"access_token", accessToken,
		"id", strconv.FormatInt(rootDepID, 10),
	)

	depResponse := new(DepartmentsResponse)
	_, err = httpGet(u, depResponse)
	return depResponse, err
}

// DepartmentMembers
// depID: 部门id
// offset: 分页的偏移量,从0开始，每次加size
// size: 每次返回的用户数量，最大100
//
func (a *Api) DepartmentMembers(depID int64, offset, size int) (*DepartmentMemberResponse, error) {
	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLDepartmentMemberByPage,
		"access_token", accessToken,
		"department_id", strconv.FormatInt(depID, 10),
		"offset", strconv.FormatInt(int64(offset), 10),
		"size", strconv.FormatInt(int64(size), 10),
	)

	memberResponse := new(DepartmentMemberResponse)
	_, err = httpGet(u, memberResponse)
	return memberResponse, err
}

func (a *Api) DepartmentAllMembers(depID int64) ([]*Member, error) {
	offset := 0
	size := 100
	hasMore := true
	resp := make([]*Member, 0, 100)

	for hasMore {
		result, err := a.DepartmentMembers(depID, offset, size)
		if err != nil {
			return nil, err
		}

		hasMore = result.HasMore
		resp = append(resp, result.UserList...)
		offset += size
	}
	return resp, nil
}

func (a *Api) FetchAllMembersByDepartmentID(depID int64) (*DepartmentMemberResponse, error) {
	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLDepartmentMembers,
		"access_token", accessToken,
		"department_id", strconv.FormatInt(depID, 10),
	)

	memberResponse := new(DepartmentMemberResponse)
	_, err = httpGet(u, memberResponse)
	return memberResponse, err
}

// Department 部门详情
func (a *Api) Department(depID int64) (*DepartmentResponse, error) {
	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLDepartment,
		"access_token", accessToken,
		"id", strconv.FormatInt(depID, 10),
	)

	depResponse := new(DepartmentResponse)
	_, err = httpGet(u, depResponse)
	return depResponse, err
}

// DepartmentMembersCount 获取部门成员数量
func (a *Api) DepartmentMembersCount(depID int64) (int, error) {
	resp, err := a.DepartmentMemberIDs(depID)
	if err != nil {
		return 0, err
	}
	return len(resp.UserIds), nil
}

func (a *Api) DepartmentMemberIDs(depID int64) (*DepartmentMemeberCountResponse, error) {
	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLDepartmentMemberIDs,
		"access_token", accessToken,
		"deptId", strconv.FormatInt(depID, 10),
	)

	depResponse := new(DepartmentMemeberCountResponse)
	_, err = httpGet(u, depResponse)
	if err != nil {
		return nil, err
	}
	return depResponse, nil
}

func (a *Api) UserIDByUnionID(unionID string) (*UserIDResponse, error) {
	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLUserIDByUnionid,
		"access_token", accessToken,
		"unionid", unionID,
	)

	userIDResponse := new(UserIDResponse)
	_, err = httpGet(u, userIDResponse)
	if err != nil {
		return nil, err
	}
	return userIDResponse, nil
}

func (a *Api) UserIDByMobile(mobile string) (*UserIDResponse, error) {
	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLUserIDByMoblie,
		"access_token", accessToken,
		"mobile", mobile,
	)

	userIDResponse := new(UserIDResponse)
	_, err = httpGet(u, userIDResponse)
	if err != nil {
		return nil, err
	}
	return userIDResponse, nil
}

func (a *Api) UserByID(userID string) (*UserResponse, error) {
	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLUserByID,
		"access_token", accessToken,
		"userid", userID,
	)

	userResponse := new(UserResponse)
	_, err = httpGet(u, userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}

func (a *Api) AuthScopes() (*AuthScopesResponse, error) {
	accessToken, err := a.AccessToken()
	if err != nil {
		return nil, err
	}

	u, _ := URLParse(
		URLAuthScopes,
		"access_token", accessToken,
	)

	authScopesResponse := new(AuthScopesResponse)
	_, err = httpGet(u, authScopesResponse)
	if err != nil {
		return nil, err
	}
	return authScopesResponse, nil
}

// accessKey：登录开发者后台，选择应用开发 > 移动接入应用 > 登录所看到应用的appId
//
func (a *Api) GetUserInfoByCode(authCode, accessKey, appSecret string) (*UserInfoResponse, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000/1000, 10)
	signature := Signature(timestamp, appSecret)

	u, _ := URLParse(
		URLGetUserInfoByCode,
		"accessKey", accessKey,
		"timestamp", timestamp,
		"signature", signature,
	)

	req := map[string]string{
		"tmp_auth_code": authCode,
	}

	userInfoResponse := new(UserInfoResponse)

	_, err := httpPost(u, req, userInfoResponse)
	if err != nil {
		return nil, err
	}
	return userInfoResponse, nil
}
