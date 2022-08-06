package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// BaseUUID uuid模板
type BaseUUID struct {
	Id         uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	CreateTime time.Time
	UpdateTime time.Time
	Deleted    gorm.DeletedAt
}

// BeforeCreate 创建时添加uuid
func (bUuid *BaseUUID) BeforeCreate(tx *gorm.DB) (err error) {
	bUuid.Id = uuid.New()
	bUuid.CreateTime = time.Now()
	bUuid.UpdateTime = time.Now()
	return
}

// BeforeUpdate 更新时添加更新时间
func (bUuid *BaseUUID) BeforeUpdate(tx *gorm.DB) (err error) {
	bUuid.UpdateTime = time.Now()
	return
}

// BaseID 自增id模板
type BaseID struct {
	Id         int `json:"id" gorm:"primary_key"`
	CreateTime time.Time
	UpdateTime time.Time
	Deleted    gorm.DeletedAt
}

// BeforeCreate 创建时添加uuid
func (bUuid *BaseID) BeforeCreate(tx *gorm.DB) (err error) {
	bUuid.CreateTime = time.Now()
	bUuid.UpdateTime = time.Now()
	return
}

// BeforeUpdate 更新时添加更新时间
func (bUuid *BaseID) BeforeUpdate(tx *gorm.DB) (err error) {
	bUuid.UpdateTime = time.Now()
	return
}

type BasePage struct {
	Page     int  `json:"page"      form:"page"  validate:"omitempty,min=0"   label:"页数"`
	PageSize int  `json:"pageSize"  form:"pageSize"  validate:"omitempty,lt=50"   label:"分页大小"`
	Desc     bool `json:"desc" form:"desc"`
}

func (s *BasePage) GetPage() (page int) {
	page = s.Page
	if page != 0 {
		page -= 1
	}
	return
}

func (s *BasePage) GetOffset() (offset int) {
	page := s.GetPage()
	offset = page * s.PageSize
	return
}
