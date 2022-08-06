package SSH

import (
	"github.com/gin-gonic/gin"
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/internal/net"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/internal/validator"
	"github.com/xtclalala/ScanNetWeb/model/SSH"
	service "github.com/xtclalala/ScanNetWeb/services/SSH"
	"github.com/xtclalala/ScanNetWeb/tools"
)

func Search(c *gin.Context) {
	var data SSH.SearchSSH
	if err := c.ShouldBindJSON(&data); err != nil {
		net.FailWhitStatus(proError.ParamResolveFault, c)
		return
	}
	if err := validator.Validate(&data); err != nil {
		net.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service.Search(&data)
	if err != nil {
		net.FailWhitStatus(proError.SearchTaskError, c)
		return
	}
	net.OkWithData(net.PageVO{
		Items: list,
		Total: total,
	}, c)
}

func Create(c *gin.Context) {
	var data SSH.CreateSSH
	if err := c.ShouldBindJSON(&data); err != nil {
		net.FailWhitStatus(proError.ParamResolveFault, c)
		return
	}
	if err := validator.Validate(&data); err != nil {
		net.FailWithMessage(err.Error(), c)
		return
	}
	state := constant.Build

	dto := &SSH.BizSSH{
		Name:     data.Name,
		Desc:     data.Desc,
		State:    state,
		FileId:   data.FileId,
		Thread:   data.Thread,
		Sheet:    data.Sheet,
		Ip:       data.Ip,
		Port:     data.Port,
		User:     data.User,
		Password: data.Password,
		Timeout:  data.Timeout,
	}

	dto.State = checkTaskIsReady(dto)

	if err := service.Create(dto); err != nil {
		net.FailWhitStatus(proError.CreateTaskError, c)
		return
	}
	net.Ok(c)
}

func Update(c *gin.Context) {
	var data SSH.UpdateSSH
	if err := c.ShouldBindJSON(&data); err != nil {
		net.FailWhitStatus(proError.ParamResolveFault, c)
		return
	}
	/** todo 确认任务是否在运行中
	 	运行中 不能修改任务信息
		完成   可以修改内容，进行第二次运行，需要一个保存记录的地方
	**/

	if err := validator.Validate(&data); err != nil {
		net.FailWithMessage(err.Error(), c)
		return
	}
	dto := &SSH.BizSSH{
		Name:     data.Name,
		Desc:     data.Desc,
		State:    constant.Build,
		FileId:   data.FileId,
		Thread:   data.Thread,
		Sheet:    data.Sheet,
		Ip:       data.Ip,
		Port:     data.Port,
		User:     data.User,
		Password: data.Password,
		Timeout:  data.Timeout,
	}

	dto.State = checkTaskIsReady(dto)
	if err := service.UpdateAll(dto); err != nil {
		net.FailWhitStatus(proError.UpdateTaskError, c)
		return
	}
	net.Ok(c)
}

func Delete(c *gin.Context) {
	var data SSH.DeleteSSH
	if err := c.ShouldBindJSON(&data); err != nil {
		net.FailWhitStatus(proError.ParamResolveFault, c)
		return
	}

	if err := validator.Validate(&data); err != nil {
		net.FailWithMessage(err.Error(), c)
		return
	}
	err := service.Delete(data.Id)
	if err != nil {
		net.FailWhitStatus(proError.DeleteTaskError, c)
		return
	}
	net.Ok(c)
}

func GetResult(c *gin.Context) {

}

func Run(c *gin.Context) {

}

func checkTaskIsReady(data *SSH.BizSSH) constant.TaskState {
	state := data.State
	if !tools.IsAllNil(data.Ip, data.Port, data.User, data.Password) &&
		!tools.IsEmpty[string](data.FileId) &&
		!tools.IsEmpty[string](data.Sheet) &&
		!tools.IsEmpty[int](data.Timeout) &&
		!tools.IsEmpty[int](data.Thread) {
		state = constant.Ready
	}
	return state
}
