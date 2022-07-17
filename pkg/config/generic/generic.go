package generic

// MysqlOptions mysql 数据链接配置信息
type MysqlOptions struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	DB       string `json:"db" yaml:"db"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
}

// RedisOptions redis 数据库连接配置信息
type RedisOptions struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	DB       int    `json:"db" yaml:"db"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Required bool   `json:"required" yaml:"required"`
	PoolSzie int    `json:"poolSize" yaml:"poolSize"`
}

// SystemOptions system config
type SystemOptions struct {
	ComponentName string `json:"componentName" yaml:"componentName"`
	Debug         bool   `json:"debug" yaml:"debug"`
	Address       string `json:"address" yaml:"address"`
	Port          int    `json:"port" yaml:"port"`
	Prefix        string `json:"prefix" yaml:"prefix"`
}
