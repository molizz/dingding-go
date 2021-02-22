package dingding

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrTokenExpired = errors.New("token expired")
)

type tokenExpiredFunc func(agentId string) (*AccessToken, error)

type AccessTokenManager interface {
	Get(agentId string, f tokenExpiredFunc) (string, error)
}

type accessTokenEntity struct {
	token     string    // token
	expiredAt time.Time // 到期时间
}

func (a *accessTokenEntity) expired() bool {
	return a.expiredAt.Unix() < time.Now().Unix()
}

var _ AccessTokenManager = (*DefaultAccessTokenManager)(nil)

// 基于内存,默认的access token管理器
type DefaultAccessTokenManager struct {
	token sync.Map
}

func (a *DefaultAccessTokenManager) Get(agentId string, f tokenExpiredFunc) (string, error) {
	tokenRaw, ok := a.token.Load(agentId)
	if !ok {
		tk, err := a.store(agentId, f)
		if err != nil {
			return "", err
		}
		return tk, nil
	}
	token := tokenRaw.(*accessTokenEntity)
	if token.expired() {
		tk, err := a.store(agentId, f)
		if err != nil {
			return "", err
		}
		return tk, nil
	}
	return token.token, nil
}

func (a *DefaultAccessTokenManager) store(agentId string, f tokenExpiredFunc) (string, error) {
	at, err := f(agentId)
	if err != nil {
		return "", err
	}

	expireAt := time.Now().Add(time.Duration(at.ExpiresIn-10) * time.Second) // 减10秒误差

	a.token.Store(agentId, &accessTokenEntity{
		token:     at.AccessToken,
		expiredAt: expireAt,
	})
	return at.AccessToken, nil
}
