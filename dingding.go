package dingding

type Config struct {
	corpId    string
	agentId   int64
	appKey    string
	appSecret string

	// event subcription
	aesKey string
	token  string
}

func NewConfig(corpId string, agentID int64, appKey, appSecret string) *Config {
	return &Config{
		corpId:    corpId,
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

	// event hub
	eventHub *EventHub
}

func New(cfg *Config, atm AccessTokenManager) *Dingding {
	if atm == nil {
		atm = NewDefaultAccessTokenManager()
	}
	return &Dingding{
		Api: &Api{
			atm:    atm,
			client: getHttpClient(),
			cfg:    cfg,
		},
		atm: atm,
	}
}

func (d *Dingding) EventHub() *EventHub {
	if d.eventHub == nil {
		d.eventHub = NewEventHub(d.cfg.appKey, d.cfg.aesKey, d.cfg.token)
	}
	return d.eventHub
}
