package dingding

type Config struct {
	agentId   int64
	appKey    string
	appSecret string

	// event subcription
	aesKey string
	token  string
}

func NewConfig(agentID int64, appKey, appSecret string) *Config {
	return &Config{
		agentId:   agentID,
		appKey:    appKey,
		appSecret: appSecret,
	}
}

func (c *Config) SetEventSubcription(aesKey, token string) {
	c.aesKey = aesKey
	c.token = token
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
