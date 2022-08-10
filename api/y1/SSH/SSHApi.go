package SSH

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/internal/net"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/internal/validator"
	"github.com/xtclalala/ScanNetWeb/model"
	"github.com/xtclalala/ScanNetWeb/model/SSH"
	service "github.com/xtclalala/ScanNetWeb/services/SSH"
	"github.com/xtclalala/ScanNetWeb/tools"
)

func Search(c *gin.Context) {
	var data SSH.SearchSSH
	if err := c.ShouldBindQuery(&data); err != nil {
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
	if err := validator.Validate(&data); err != nil {
		net.FailWithMessage(err.Error(), c)
		return
	}

	task, err := service.FindById(data.Id)
	if err != nil {
		net.FailWhitStatus(proError.FindTaskError, c)
		return
	}

	if err = tryUpdateTaskState(&task); err != nil {
		net.FailWhitStatus(proError.TaskCantRevise, c)
		return
	}
	dto := &SSH.BizSSH{
		BaseID: model.BaseID{
			Id: data.Id,
		},
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
	if err = service.UpdateAll(dto); err != nil {
		net.FailWhitStatus(proError.UpdateTaskError, c)
		return
	}
	net.Ok(c)
}

func Delete(c *gin.Context) {
	var data SSH.DeleteSSH
	if err := c.ShouldBindQuery(&data); err != nil {
		net.FailWhitStatus(proError.ParamResolveFault, c)
		return
	}
	if err := validator.Validate(&data); err != nil {
		net.FailWithMessage(err.Error(), c)
		return
	}
	task, err := service.FindById(data.Id)
	if err != nil {
		net.FailWhitStatus(proError.FindTaskError, c)
		return
	}

	if err = tryUpdateTaskState(&task); err != nil {
		net.FailWhitStatus(proError.TaskCantRevise, c)
		return
	}

	if err = service.Delete(data.Id); err != nil {
		net.FailWhitStatus(proError.DeleteTaskError, c)
		return
	}
	net.Ok(c)
}

func GetResult(c *gin.Context) {
	var data SSH.SearchSSHResult
	if err := c.ShouldBindQuery(&data); err != nil {
		net.FailWhitStatus(proError.ParamResolveFault, c)
		return
	}
	if err := validator.Validate(&data); err != nil {
		net.FailWithMessage(err.Error(), c)
		return
	}
	vo, err := service.SelectResult(&data)
	if err != nil {
		net.FailWhitStatus(proError.SearchTaskResultError, c)
		return
	}
	net.OkWithData(vo, c)
}

func Run(c *gin.Context) {
	var data SSH.RunSSH
	if err := c.ShouldBindJSON(&data); err != nil {
		net.FailWhitStatus(proError.ParamResolveFault, c)
		return
	}
	if err := validator.Validate(&data); err != nil {
		net.FailWithMessage(err.Error(), c)
		return
	}
	task, err := service.FindById(data.Id)
	if err != nil {
		net.FailWhitStatus(proError.FindTaskError, c)
		return
	}
	if err = updateTaskStateDoing(&task); err != nil {
		net.FailWhitStatus(proError.TaskCantRevise, c)
		return
	}
	if err = service.UpdateState(task.Id, constant.Doing); err != nil {
		net.FailWhitStatus(proError.UpdateTaskError, c)
		return
	}
	if err = service.Run(&task); err != nil {
		net.FailWhitStatus(proError.RunTaskError, c)
		return
	}

	net.Ok(c)
}

func checkTaskIsReady(data *SSH.BizSSH) constant.TaskState {
	state := constant.Build
	if !tools.IsAllNil(data.Ip, data.Port, data.User, data.Password) &&
		!tools.IsEmpty[uuid.UUID](data.FileId) &&
		!tools.IsEmpty[string](data.Sheet) &&
		!tools.IsEmpty[int](data.Timeout) &&
		!tools.IsEmpty[int](data.Thread) {
		state = constant.Ready
	}
	return state
}

// 	确认任务是否在运行中
//	运行中, 完成 不能修改任务信息
func tryUpdateTaskState(data *SSH.BizSSH) (err error) {
	if err = updateTaskStateDoing(data); err != nil {
		return
	}
	return nil
}

func updateTaskStateDoing(data *SSH.BizSSH) (err error) {
	if constant.Doing == data.State {
		return errors.New("task is running, you can't change state!")
	}
	return nil
}

func updateTaskStateFinish(data *SSH.BizSSH) (err error) {
	if constant.Finish == data.State {
		return errors.New("task is running, you can't change state!")
	}
	return nil
}
