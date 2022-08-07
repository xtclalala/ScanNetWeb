package SSH

import (
	"fmt"
	"github.com/xtclalala/ScanNetWeb/conf"
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/global"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/model/SSH"
	"github.com/xtclalala/ScanNetWeb/services/file"
	"github.com/xtclalala/ScanNetWeb/tools"
	"gorm.io/gorm/clause"
	"sync"
)

func Search(dto *SSH.SearchSSH) (list []SSH.BizSSH, total int64, err error) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&SSH.BizSSH{}).Where(&SSH.BizSSH{
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
	err = proError.WrapOrNil(err, "Search ssh task: %s fail!")
	return
}

func Create(dto *SSH.BizSSH) (err error) {
	err = global.Db.Create(&dto).Error
	return proError.WrapOrNil(err, "Create ssh task: %s fail!", dto.Name)
}

func UpdateAll(dto *SSH.BizSSH) (err error) {
	err = global.Db.Updates(dto).Error
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

func Run(task *SSH.BizSSH) (err error) {

	// read check file
	workers := new([]*worker)
	// read file
	fileData, err := file.FindById(task.FileId)
	if err != nil {
		return err
	}
	if err = readFile(workers, tools.FileAbsPath(fileData.Path, fileData.Id.String()), task.Sheet, *task.Ip, *task.Port, *task.User, *task.Password); err != nil {
		return err
	}
	osConfig := conf.LinuxScanConfig()
	var fns []func()

	data := &sync.Map{}
	for _, item := range *workers {
		var fn = func(worker *worker) func() {
			return func() {
				// connect
				s := tools.NewSsh(worker.ip, worker.port, worker.user, worker.password, task.Timeout, osConfig)
				s.Connect()
				s.GetOS()
				values := s.Save()
				// you can do something, run diy cmd
				res := s.ScanOS()
				values = append(values, res...)
				data.Store(worker.ip, values)
			}

		}(item)
		fns = append(fns, fn)

	}
	go func() {
		tools.Start(fns, task.Thread)
		data.Range(func(key, value any) bool {
			fmt.Printf("key: %s, value: %s\n", key, value)
			return true
		})
		// todo save data
		UpdateState(task.Id, constant.Finish)
		// todo ws notify success or fail

	}()

	return
}
