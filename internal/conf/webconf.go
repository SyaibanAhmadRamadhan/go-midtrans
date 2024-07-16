package conf

type WebConf struct {
	RestfullApiPort int    `json:"RESTFULL_API_PORT" mapstructure:"RESTFULL_API_PORT"`
	RestfullApiHost string `json:"RESTFULL_API_HOST" mapstructure:"RESTFULL_API_HOST"`
}
