package file

import (
	"github.com/google/uuid"
	"github.com/xtclalala/ScanNetWeb/model"
	"github.com/xtclalala/ScanNetWeb/tools"
	"gorm.io/gorm"
	"time"
)

type BizFile struct {
	model.BaseUUID
	Name string `json:"name" gorm:"not null;comment:文件名称;"`
	Type string `json:"type" gorm:"not null;comment:文件类型;"`
	// 程序运行的当前路径为根路径
	Path string `json:"path" gorm:"not null;comment:文件存储路径;"`
}

func (u *BizFile) BeforeCreate(tx *gorm.DB) (err error) {
	if tools.IsEmpty(u.BaseUUID.Id) {
		u.BaseUUID.Id = uuid.New()
	}
	u.BaseUUID.CreateTime = time.Now()
	u.BaseUUID.UpdateTime = time.Now()
	return
}

type DownloadFile struct {
	Id uuid.UUID `json:"id" validate:"required" label:"文件Id"`
}
