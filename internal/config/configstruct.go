package config

type Config struct {
	Server Server       `mapstructure:"server"`
	Logs1  Logs         `mapstructure:"logs1"`
	Logs2  Logs         `mapstructure:"logs2"`
	UserDB DatabaseConn `mapstructure:"userdb"`
	Cache  CacheParams  `mapstructure:"cache"`
}

type Server struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Certif    string `mapstructure:"certif"`
	CertifKey string `mapstructure:"certifkey"`
}

type Logs struct {
	Path  string `mapstructure:"path"`
	Level int    `mapstructure:"level"`
}

type DatabaseConn struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"dbname"`
}

type CacheParams struct {
	ExpirationTime int `mapstrcture:"expirationtime"`
	ExpurgeTime    int `mapstructure:"expurgetime"`
}
