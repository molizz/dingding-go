package dingding

type Config struct {
	CorpId    string
	AgentId   int64
	AppKey    string
	AppSecret string

	// event subcription
	AesKey string
	Token  string
}

func NewConfig(corpId string, agentID int64, appKey, appSecret string) *Config {
	return &Config{
		CorpId:    corpId,
		AgentId:   agentID,
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

func (c *Config) SetEventSubcription(aesKey, token string) {
	c.AesKey = aesKey
	c.Token = token
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
		d.eventHub = NewEventHub(d.cfg.AppKey, d.cfg.AesKey, d.cfg.Token)
	}
	return d.eventHub
}
