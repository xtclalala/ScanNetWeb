package global

import (
	"github.com/spf13/viper"
	"github.com/xtclalala/ScanNetWeb/conf"
	"github.com/xtclalala/ScanNetWeb/internal/proLog"
	"gorm.io/gorm"
)

var (
	Viper  *viper.Viper
	System conf.Config
	Db     *gorm.DB
	Log    *proLog.Log
)
