package config

type Config struct {
	GatewayIP   string `env:"GATEWAY_IP"`
	GatewayPort string `env:"GATEWAY_PORT"`

	ClientID     string `env:"CLIENT_ID" mapstructure:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET" mapstructure:"CLIENT_SECRET"`

	ServerIP   string `env:"SERVER_IP"`
	ServerPort string `env:"SERVER_PORT"`

	DBAdapter  string `env:"DB_ADAPTER"`
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUsername string `env:"DB_USER" mapstructure:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD" mapstructure:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME" mapstructure:"DB_NAME"`

	TimeZone string `env:"TIMEZONE"`

	KakaoJWKs []KakaoJWK `json:"keys"`
}

type KakaoJWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

var config = &Config{
	GatewayIP:   "localhost",
	GatewayPort: "9091",
	ServerIP:    "localhost",
	ServerPort:  "9090",

	DBAdapter: "postgres",
	DBHost:    "192.168.64.1",
	DBPort:    "5432",

	TimeZone: "Asia/Seoul",
}

func GetConfig() *Config {
	return config
}
