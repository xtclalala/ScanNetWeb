package SSH

import (
	"github.com/google/uuid"
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/model"
)

type BizSSH struct {
	model.BaseID
	Name     string             ` gorm:"not null;comment:任务name;"`
	Desc     string             ` gorm:"comment:任务相关信息;"`
	State    constant.TaskState ` gorm:"default:1;comment:任务状态"`
	FileId   uuid.UUID          ` gorm:"comment:文件"`
	Thread   int                ` gorm:"default:5;comment:thread"`
	Sheet    string             ` gorm:"default:Sheet1;comment:sheet"`
	Ip       *int               ` gorm:"comment:ip在文件的列数"`
	Port     *int               ` gorm:"comment:port在文件的列数"`
	User     *int               ` gorm:"comment:user在文件的列数"`
	Password *int               ` gorm:"comment:password在文件的列数"`
	Timeout  int                ` gorm:"default:5;comment:ssh连接超时事件"`

	BizSSHResult []BizSSHResult `gorm:"foreignKey:TaskId"`
}

type SearchSSH struct {
	model.BasePage
	Name  string             `json:"name"`
	State constant.TaskState `json:"state"`
}

type CreateSSH struct {
	Name     string    `json:"name" validate:"max=50,min=1,required" label:"任务名"`
	Desc     string    `json:"desc" validate:"-" label:"简介"`
	Thread   int       `json:"thread" validate:"omitempty,max=2000,min=1" label:"并发量"`
	FileId   uuid.UUID `json:"fileId" validate:"-" label:"文件"`
	Sheet    string    `json:"sheet" validate:"omitempty,max=50,min=1" label:"工作表"`
	Ip       *int      `json:"ip" validate:"-" label:"目标地址列"`
	Port     *int      `json:"port" validate:"-" label:"端口列"`
	User     *int      `json:"user" validate:"-" label:"用户列"`
	Password *int      `json:"password" validate:"-" label:"密码列"`
	Timeout  int       `json:"timeout" validate:"omitempty,max=20,min=1" label:"超时时间"`
}

type UpdateSSH struct {
	Id       int       `json:"id" validate:"required" label:"任务Id"`
	Name     string    `json:"name" validate:"max=50,min=1,required" label:"任务Id"`
	Desc     string    `json:"desc" validate:"-" label:"任务Id"`
	Thread   int       `json:"thread" validate:"omitempty,max=2000,min=1" label:"并发量"`
	FileId   uuid.UUID `json:"fileId" validate:"-" label:"文件"`
	Sheet    string    `json:"sheet" validate:"omitempty,max=50,min=1" label:"工作表"`
	Ip       *int      `json:"ip" validate:"-" label:"目标地址列"`
	Port     *int      `json:"port" validate:"-" label:"端口列"`
	User     *int      `json:"user" validate:"-" label:"用户列"`
	Password *int      `json:"password" validate:"-" label:"密码列"`
	Timeout  int       `json:"timeout" validate:"omitempty,max=20,min=1" label:"超时时间"`
}

type DeleteSSH struct {
	Id int `json:"id" validate:"required" label:"任务Id"`
}

type RunSSH struct {
	Id int `json:"id" validate:"required" label:"任务Id"`
}

type BizSSHResult struct {
	model.BaseUUID
	TaskId   int    `gorm:"not null;comment:任务Id;"`
	Addr     string `gorm:"comment:目标地址;"`
	User     string `gorm:"comment:账号;"`
	Password string `gorm:"comment:密码;"`
	Os       string `gorm:"comment:操作系统;"`
	Result   string `gorm:"comment:结果;"`
}

type SearchSSHResult struct {
	model.BasePage
	TaskId int    `json:"taskId" validator:"required" label:"任务"`
	Os     string `json:"os"`
}
