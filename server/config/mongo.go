package config

type Mongo struct {
	Host       string `mapstructure:"host" json:"host" yaml:"host"`
	DataBase   string `mapstructure:"database" json:"database" yaml:"database"`
	Username   string `mapstructure:"username" json:"username" yaml:"username"`
	Password   string `mapstructure:"password" json:"password" yaml:"password"`
}
