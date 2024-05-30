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

	RedisHost     string `env:"REDIS_HOST" mapstructure:"REDIS_HOST"`
	RedisPort     string `env:"REDIS_PORT" mapstructure:"REDIS_PORT"`
	RedisPassword string `env:"REDIS_PASSWORD" mapstructure:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB" mapstructure:"REDIS_DB"`

	TimeZone string `env:"TIMEZONE"`

	KakaoJWKs []Key `json:"keys"`
}

type Key struct {
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
	DBHost:    "192.168.0.64",
	DBPort:    "5432",

	RedisHost: "192.168.0.64",
	RedisPort: "6379",

	TimeZone: "Asia/Seoul",

	KakaoJWKs: []Key{
		{
			Kid: "3f96980381e451efad0d2ddd30e3d3",
			Kty: "RSA",
			Alg: "RS256",
			Use: "sig",
			N:   "q8zZ0b_MNaLd6Ny8wd4cjFomilLfFIZcmhNSc1ttx_oQdJJZt5CDHB8WWwPGBUDUyY8AmfglS9Y1qA0_fxxs-ZUWdt45jSbUxghKNYgEwSutfM5sROh3srm5TiLW4YfOvKytGW1r9TQEdLe98ork8-rNRYPybRI3SKoqpci1m1QOcvUg4xEYRvbZIWku24DNMSeheytKUz6Ni4kKOVkzfGN11rUj1IrlRR-LNA9V9ZYmeoywy3k066rD5TaZHor5bM5gIzt1B4FmUuFITpXKGQZS5Hn_Ck8Bgc8kLWGAU8TzmOzLeROosqKE0eZJ4ESLMImTb2XSEZuN1wFyL0VtJw",
			E:   "AQAB",
		},
		{
			Kid: "9f252dadd5f233f93d2fa528d12fea",
			Kty: "RSA",
			Alg: "RS256",
			Use: "sig",
			N:   "qGWf6RVzV2pM8YqJ6by5exoixIlTvdXDfYj2v7E6xkoYmesAjp_1IYL7rzhpUYqIkWX0P4wOwAsg-Ud8PcMHggfwUNPOcqgSk1hAIHr63zSlG8xatQb17q9LrWny2HWkUVEU30PxxHsLcuzmfhbRx8kOrNfJEirIuqSyWF_OBHeEgBgYjydd_c8vPo7IiH-pijZn4ZouPsEg7wtdIX3-0ZcXXDbFkaDaqClfqmVCLNBhg3DKYDQOoyWXrpFKUXUFuk2FTCqWaQJ0GniO4p_ppkYIf4zhlwUYfXZEhm8cBo6H2EgukntDbTgnoha8kNunTPekxWTDhE5wGAt6YpT4Yw",
			E:   "AQAB",
		},
	},
}

func GetConfig() *Config {
	return config
}
