package SSH

import (
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/model"
)

type BizSSH struct {
	model.BaseID
	Name   string             `json:"name" gorm:"comment:任务name;"`
	Desc   string             `json:"desc" gorm:"comment:任务相关信息;"`
	State  constant.TaskState `json:"state" gorm:"default:1;comment:任务状态"`
	Ip     string             `json:"ip" gorm:"not null;comment:ip"`
	Thread int                `json:"thread" gorm:"default:5;comment:thread"`
}

type SearchSSH struct {
	model.BasePage
	Name  string             `json:"name"`
	State constant.TaskState `json:"state" `
}

type CreateSSH struct {
	Name   string `json:"name" validate:"max=50,min=1,required" label:"任务名"`
	Desc   string `json:"desc" validate:"-" label:"简介"`
	Ip     string `json:"ip" validate:"required" label:"目标地址"`
	Thread int    `json:"thread" validate:"max=2000,min=1,required" label:"并发量"`
}

type UpdateSSH struct {
	Id     int    `json:"id" validate:"required" label:"任务Id"`
	Name   string `json:"name" validate:"max=50,min=1,required" label:"任务Id"`
	Desc   string `json:"desc" validate:"required" label:"任务Id"`
	Ip     string `json:"ip" validate:"required" label:"目标地址"`
	Thread int    `json:"thread" validate:"max=2000,min=1,required" label:"并发量"`
}

type DeleteSSH struct {
	Id int `json:"id" validate:"required" label:"任务Id"`
}

type RunSSH struct {
	Id int `json:"id" validate:"required" label:"任务Id"`
}
