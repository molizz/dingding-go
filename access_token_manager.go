package dingding

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrTokenExpired = errors.New("token expired")
)

type tokenExpiredFunc func(agentId int64) (*AccessToken, error)

type AccessTokenManager interface {
	Get(agentId int64) (string, error)
	Set(agentID int64, accessToken *AccessToken)
}

type accessTokenEntity struct {
	token     string    // token
	expiredAt time.Time // 到期时间
}

func (a *accessTokenEntity) expired() bool {
	return a.expiredAt.Unix() < time.Now().Unix()
}

var _ AccessTokenManager = (*DefaultAccessTokenManager)(nil)

// 默认基于内存的access token管理器
type DefaultAccessTokenManager struct {
	token sync.Map
}

func NewDefaultAccessTokenManager() *DefaultAccessTokenManager {
	return &DefaultAccessTokenManager{}
}

func (a *DefaultAccessTokenManager) Get(agentID int64) (string, error) {
	tokenRaw, ok := a.token.Load(agentID)
	if !ok {
		return "", ErrTokenExpired
	}

	token := tokenRaw.(*accessTokenEntity)
	if token.expired() {
		return "", ErrTokenExpired
	}
	return token.token, nil
}

func (a *DefaultAccessTokenManager) Set(agentID int64, accessToken *AccessToken) {
	expireAt := time.Now().Add(time.Duration(accessToken.ExpiresIn-10) * time.Second) // 减10秒误差

	a.token.Store(agentID, &accessTokenEntity{
		token:     accessToken.AccessToken,
		expiredAt: expireAt,
	})
}
