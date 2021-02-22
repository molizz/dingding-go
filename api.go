package dingding

const (
	URLAccessToken = "https://oapi.dingtalk.com/gettoken"
)

type Api struct {
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
