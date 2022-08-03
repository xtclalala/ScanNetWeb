package conf

type Config struct {
	App    *App    `mapstructure:"app" yaml:"app"`
	Db     *Db     `mapstructure:"db"  yaml:"db"`
	Logger *Logger `mapstructure:"logger"  yaml:"logger"`
	Jwt    *Jwt    `mapstructure:"jwt"  yaml:"jwt"`
	File   *File   `mapstructure:"file"  yaml:"file"`
}

type App struct {
	Port string `mapstructure:"port" yaml:"port"`
	Auth string `mapstructure:"auth" yaml:"auth"`
}

type File struct {
	Path string `mapstructure:"path" yaml:"path"`
}

type Db struct {
	DbType string `mapstructure:"type" yaml:"type"`
	Host   string `mapstructure:"host" yaml:"host"`
	Port   string `mapstructure:"port" yaml:"port"`
	User   string `mapstructure:"user" yaml:"user"`
	Passwd string `mapstructure:"passwd" yaml:"passwd"`
	DbName string `mapstructure:"dbName" yaml:"dbName"`
}

type Logger struct {
	FilePath string `mapstructure:"filePath" yaml:"filePath"`
	FileName string `mapstructure:"fileName" yaml:"fileName"`
}

type Jwt struct {
	SignKey    string `mapstructure:"signKey" yaml:"signKey"`
	Timeout    int64  `mapstructure:"timeout" yaml:"timeout"`
	Iss        string `mapstructure:"iss" yaml:"iss"`
	BufferTime int64  `mapstructure:"bufferTime" yaml:"bufferTime"`
}

type LinuxScan struct {
	File     *file    `mapstructure:"file"`
	Burst    *burst   `mapstructure:"burst"`
	Connect  *connect `mapstructure:"connect"`
	OsConfig map[string][]string
}

type file struct {
	InFileName  string `mapstructure:"inFileName"`
	OutFileName string `mapstructure:"outFileName"`
	Sheet       string `mapstructure:"sheet"`
	Ip          int    `mapstructure:"ip"`
	Port        int    `mapstructure:"port"`
	User        int    `mapstructure:"user"`
	Password    int    `mapstructure:"password"`
}

type burst struct {
	BurstNum int `mapstructure:"burstNum"`
}

type connect struct {
	Timeout int64 `mapstructure:"timeout"`
}
