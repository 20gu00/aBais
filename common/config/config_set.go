package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(AppConfig)

type AppConfig struct {
	Addr          string `mapstructure:"Addr"`
	WSAddr        string `mapstructure:"WsAddr"`
	KubeConfigs   string `mapstructure:"kubeConfigs"`
	AdminUser     string `mapstructure:"AdminUser"`
	AdminPassword string `mapstructure:"AdminPwd"`

	PodLogTailLine int  `mapstructure:"PodLogTailLine"`
	LogMode        bool `mapstructure:"LogMode"`

	UploadPath string `mapstructure:"UploadPath"`

	*DBConf    `mapstructure:"DB"`
	*LogConfig `mapstructure:"log"`
}

type DBConf struct {
	DBType       string `mapstructure:"DbType"`
	DBHost       string `mapstructure:"DbHost"`
	DBPort       int    `mapstructure:"DbPort"`
	DBName       string `mapstructure:"DbName"`
	DBUser       string `mapstructure:"DbUser"`
	DBPassword   string `mapstructure:"DbPwd"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
	MaxLifeTime  int    `mapstructure:"MaxLifeTime"`
}

type LogConfig struct {
	Level     string `mapstructure:"level"`
	Filename  string `mapstructure:"file_name"`
	MaxSize   int    `mapstructure:"max_size"`
	MaxAge    int    `mapstructure:"max_age"`
	MaxBackup int    `mapstructure:"max_backup"`
	Compress  bool   `mapstructure:"compress"`
}

func ConfigRead(configFile string) error {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml") // 针对运城配置存储使用
		viper.AddConfigPath("./conf")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("未找到配置文件", err)
		} else {
			fmt.Println("读取配置文件失败", err)
			return err
		}
	}

	if err := viper.Unmarshal(Config); err != nil {
		fmt.Println("将配置信息添加进结构体操作失败", err)
		return err
	}

	// 配置及文件热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已经修改")
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Println("配置文件已经修改,将配置写入结构体操作失败", err)
		}
	})
	return nil
}
