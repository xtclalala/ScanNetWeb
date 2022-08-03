package initServe

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/global"
	"log"
)

func InitConfig() *viper.Viper {
	temp := viper.New()
	temp.SetConfigFile(constant.ConfigPath)
	if err := temp.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件错误", err)
	}
	env := temp.GetString("env")

	v := viper.New()
	v.SetConfigFile(constant.ConfigName + constant.Rung + env + constant.Point + constant.ConfigType)
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("配置文件读取失败: %s \n", env))
	}
	if err := v.Unmarshal(&global.System); err != nil {
		panic(fmt.Errorf("结构化数据失败：%s \n", err))
	}
	v.OnConfigChange(func(event fsnotify.Event) {
		fmt.Println("配置文件数据更改：", event.Name)
		if err := v.Unmarshal(&global.System); err != nil {
		}
	})

	return v
}

func InitLinuxScanConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/config-linuxScan.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = viper.Unmarshal(global.LinuxScanConfig)
	global.LinuxScanConfig.OsConfig = viper.GetStringMapStringSlice("os")
	if err != nil {
		log.Fatalf(err.Error())
	}

}
