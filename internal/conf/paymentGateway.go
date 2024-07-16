package conf

type PaymentGateway struct {
	Midtrans struct {
		IDMerchant string `json:"ID_MERCHANT" mapstructure:"ID_MERCHANT"`
		ServerKey  string `json:"SERVER_KEY" mapstructure:"SERVER_KEY"`
		ClientKey  string `json:"CLIENT_KEY" mapstructure:"CLIENT_KEY"`
	} `json:"MIDTRANS" mapstructure:"MIDTRANS"`
}
