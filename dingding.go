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
