package global

type ConfigManger struct {
	Mysql Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Log Log        `mapstructure:"log" json:"log" yaml:"log"`
	WebSocket WebSocket  `mapstructure:"websocket" json:"websocket" yaml:"websocket"`
}

type WebSocket struct {
	Port string      `mapstructure:"port" json:"port" yaml:"port"`
	MaxClient int `mapstructure:"maxclient" json:"maxclient" yaml:"maxclient"`
	InterValReq int `mapstructure:"intervalReq" json:"intervalReq" yaml:"intervalReq"`
}

type Mysql struct {
    Host string       `mapstructure:"host" json:"host" yaml:"host"`
    Port int          `mapstructure:"port" json:"port" yaml:"port"`
    DB   string       `mapstructure:"db" json:"db" yaml:"db"`
    UserName string   `mapstructure:"username" json:"username" yaml:"username"`
    PassWord string   `mapstructure:"password" json:"password" yaml:"password"`
}
type Redis struct {
	HostAndPort string `mapstructure:"host_port" json:"host_port" yaml:"host_port"`
	PassWord string    `mapstructure:"password" json:"password" yaml:"password"`
    MaxIdle int        `mapstructure:"maxidle" json:"maxidle" yaml:"maxidle"`
	MaxActive int      `mapstructure:"maxactive" json:"maxactive" yaml:"maxactive"`
	DB  int			   `mapstructure:"db" json:"db" yaml:"db"`
	IdleTimeout int    `mapstructure:"idletimeout" json:"idletimeout" yaml:"idletimeout"`
}

type Log struct {
	Mode int         `mapstructure:"mode" json:"mode" yaml:"mode"`
	LocalPath string `mapstructure:"localpath" json:"localpath" yaml:"localpath"`
	Level string     `mapstructure:"level" json:"level" yaml:"level"`
}