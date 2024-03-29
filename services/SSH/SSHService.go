package SSH

import (
	"encoding/json"
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/global"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/model/SSH"
	"github.com/xtclalala/ScanNetWeb/model/file"
	fileService "github.com/xtclalala/ScanNetWeb/services/file"
	wsServicxe "github.com/xtclalala/ScanNetWeb/services/ws"
	"github.com/xtclalala/ScanNetWeb/tools"
	"gorm.io/gorm/clause"
	"io"
	"strings"
	"sync"
)

func Search(dto *SSH.SearchSSH) (list []SSH.BizSSH, total int64, err error) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&SSH.BizSSH{}).Preload("File").Where(&SSH.BizSSH{
		State: dto.State,
	})

	if dto.Name != "" {
		db.Where("name like ?", "%"+dto.Name+"%")
	}

	oc := clause.OrderByColumn{
		Column: clause.Column{Name: "create_time", Raw: true},
		Desc:   dto.Desc,
	}
	err = db.Count(&total).Limit(limit).Offset(offset).Order(oc).Find(&list).Error
	err = proError.WrapOrNil(err, "Search ssh task fail!")
	return
}

func Create(dto *SSH.BizSSH) (err error) {
	err = global.Db.Create(&dto).Error
	return proError.WrapOrNil(err, "Create ssh task: %s fail!", dto.Name)
}

func UpdateAll(dto *SSH.BizSSH) (err error) {
	var fileData []file.BizFile
	err = global.Db.Updates(dto).Association("File").Replace(fileData)
	return proError.WrapOrNil(err, "Update ssh task: %s fail!", dto.Id)
}

func UpdateState(id int, state constant.TaskState) (err error) {
	err = global.Db.Model(&SSH.BizSSH{}).Where("id = ?", id).Update("state", state).Error
	return proError.WrapOrNil(err, "Update ssh task: %s fail!", id)
}

func Delete(id int) (err error) {
	err = global.Db.Delete(&SSH.BizSSH{}, id).Error
	return proError.WrapOrNil(err, "Delete ssh task: %s fail!", id)
}

func FindById(id int) (task SSH.BizSSH, err error) {
	err = global.Db.First(&task, id).Error
	return task, proError.WrapOrNil(err, "find ssh task: %s fail!", id)
}

func Run(task *SSH.BizSSH) (err error) {

	// read check file
	workers := new([]*worker)
	// read file
	fileData, err := fileService.FindById(task.FileId)
	if err != nil {
		return err
	}
	if err = readFile(workers, tools.FileAbsPath(fileData.Path, fileData.Id.String()), task.Sheet, *task.Ip, *task.Port, *task.User, *task.Password); err != nil {
		return err
	}
	var fns []func()

	data := &sync.Map{}
	for _, item := range *workers {
		var fn = func(worker *worker) func() {
			return func() {
				s := tools.NewSsh(worker.ip, worker.port, worker.user, worker.password, task.Timeout)
				s.Connect()
				// you can do something, run diy cmd
				res := s.ScanOS()
				// todo
				data.Store(worker, res)
			}

		}(item)
		fns = append(fns, fn)

	}
	go func() {
		tools.Start(fns, task.Thread)
		var dataList []SSH.BizSSHResult
		data.Range(func(key, value any) bool {
			worker := key.(*worker)
			res := strings.Split(strings.ReplaceAll(value.(string), "nohup.out", "*"), "||")

			parse := SSH.BizSSHResult{
				TaskId:   task.Id,
				Addr:     worker.ip,
				User:     worker.user,
				Password: worker.password,
			}
			var items []SSH.BizSSHResultItem
			for _, re := range res {
				dec := json.NewDecoder(strings.NewReader(re))
				item := SSH.BizSSHResultItem{}
				for {
					if err := dec.Decode(&item); err == io.EOF {
						break
					} else if err != nil {
						// todo log err  not panic
						break
					}
				}
				items = append(items, item)
			}
			parse.Items = items
			dataList = append(dataList, parse)
			return true
		})
		// save data
		rErr := CreateResultPare(dataList)

		if err = UpdateState(task.Id, constant.Finish); err != nil || rErr != nil {
			wsServicxe.PushMessage(wsServicxe.NewMessage("扫描失败", "扫描结果存储失败或更新任务状态失败", constant.Error, constant.LinuxScan))
			return
		}
		wsServicxe.PushMessage(wsServicxe.NewMessage("扫描成功", task.Name+"任务完成", constant.Success, constant.LinuxScan))

	}()

	return
}

type worker struct {
	ip       string
	port     string
	user     string
	password string
}

func readFile(workers *[]*worker, inFilename, sheet string, ip, port, user, password int) error {

	rows, err := tools.Readfile(inFilename, sheet)
	if err != nil {
		return proError.Wrap(err, "read file %s is fail", inFilename)
	}
	for _, row := range rows {
		i := &worker{
			ip:       row[ip],
			port:     row[port],
			user:     row[user],
			password: row[password],
		}
		*workers = append(*workers, i)
	}
	return nil
}
