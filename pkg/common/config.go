package common

type Configuration struct {
	Mysql MysqlConfiguration `mapstructure:"mysql" json:"mysql"`
}

type MysqlConfiguration struct {
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Query    string `mapstructure:"query" json:"query"`
}
