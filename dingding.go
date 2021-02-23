package dingding

type Config struct {
	agentId   int64
	appKey    string
	appSecret string
}

func NewConfig(agentID int64, appKey, appSecret string) *Config {
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
			atm:    atm,
			client: getHttpClient(),
			cfg:    cfg,
		},
		atm: atm,
	}
}

func (d *Dingding) AccessToken() (string, error) {
	token, err := d.atm.Get(d.cfg.agentId)
	if err != nil {
		if err == ErrTokenExpired {
			at, err := d.Api.AccessToken()
			if err != nil {
				return "", err
			}
			d.atm.Set(d.cfg.agentId, at)
			return at.AccessToken, nil
		}
		return "", err
	}

	return token, nil
}
