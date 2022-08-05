package task

import (
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/model"
)

type BizTask struct {
	model.BaseID
	Name  string             `json:"name" gorm:"comment:任务name;"`
	Desc  string             `json:"desc" gorm:"comment:任务相关信息;"`
	State constant.TaskState `json:"state" gorm:"default:1;comment:任务状态"`

	// o2m
	SysPermissions []SysPermission `json:"permissions" gorm:"foreignKey:SysMenuId"`
	// m2m
	SysRoles []SysRole `json:"roles" gorm:"many2many:m2m_role_menu;"`
}
