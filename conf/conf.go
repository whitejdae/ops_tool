package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(ConfigYaml)

type ConfigYaml struct {
	*System          `mapstructure:"system"`
	*JenkinsDingTalk `mapstructure:"jenkinsDingTalk"`
}

type JenkinsDingTalk struct {
	*Jenkins  `mapstructure:"jenkins"`
	*DingDing `mapstructure:"dingDing"`
}

type Jenkins struct {
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
}

type DingDing struct {
	DefaultNotifier string `mapstructure:"default_notifier"`
	Token           string `mapstructure:"token"`
}

type System struct {
	*App   `mapstructure:"app"`
	*Log   `mapstructure:"log"`
	*Mysql `mapstructure:"mysql"`
}

type App struct {
	Mode     string `mapstructure:"mode"`
	Host     string `mapstructure:"host"`
	Port     int64  `mapstructure:"port"`
	InitData bool   `mapstructure:"init-data"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	LogPath    string `mapstructure:"logPath"`
	Compress   bool   `mapstructure:"compress"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
}

type Mysql struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	MaxOpenCons int    `mapstructure:"max_open_cons"`
	MaxIdleCons int    `mapstructure:"max_idle_cons"`
}

func InitConfig() {
	viper.SetConfigFile("./conf/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 未读取配置文件
			panic(fmt.Errorf("未读取到配置文件:%s", err))
		} else {
			panic(fmt.Errorf("configFile load failed, err:%s", err))
		}
	}

	// 把文件反序列到结构体
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s", err))
	}

	// 监听配置文件发生变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Printf("Config file changed:%v", in.Name)
		// 把文件反序列到结构体
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s", err))
			return
		}
	})
}
