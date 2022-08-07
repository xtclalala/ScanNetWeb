package global

import (
	"github.com/spf13/viper"
	"github.com/xtclalala/ScanNetWeb/conf"
	"gorm.io/gorm"
)

var (
	Viper  *viper.Viper
	System conf.Config
	Db     *gorm.DB
)
