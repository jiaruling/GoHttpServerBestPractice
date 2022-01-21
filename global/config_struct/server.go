package config_struct

/*
   功能说明: server.yaml配置文件对应的结构体
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 10:54
*/

type ServerConfig struct {
	Name                 string `mapstructure:"name"`
	Env                  string `mapstructure:"env"`
	RunMode              string `mapstructure:"runMode"`
	RunPort              string `mapstructure:"runPort"`
	StaticServiceRunPort string `mapstructure:"staticServiceRunPort"`
	MySQL                Mysql  `mapstructure:"mysql"`
}

type Mysql struct {
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Ip        string `mapstructure:"ip"`
	Port      int    `mapstructure:"port"`
	Db        string `mapstructure:"db"`
	Parameter string `mapstructure:"parameter"`
}
