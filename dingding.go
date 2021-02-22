package dingding

type Config struct {
	agentId   string
	appKey    string
	appSecret string
}

func NewConfig(agentID, appKey, appSecret string) *Config {
	return &Config{
		agentId:   agentID,
		appKey:    appKey,
		appSecret: appSecret,
	}
}

type Dingding struct {
	*Api

	// Access Token manager
	atm AccessTokenManager
}

func New(cfg *Config, atm AccessTokenManager) *Dingding {
	return &Dingding{
		Api: &Api{
			client: getHttpClient(),
			cfg:    cfg,
		},
		atm: atm,
	}
}

func (d *Dingding) AccessToken() (string, error) {
	token, err := d.atm.Get(d.cfg.agentId, func(agentId string) (*AccessToken, error) {
		return d.Api.AccessToken()
	})
	if err != nil {
		return "", err
	}
	return token, nil
}
