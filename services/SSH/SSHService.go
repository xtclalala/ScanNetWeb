package SSH

import (
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xtclalala/ScanNetWeb/global"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/model/SSH"
	"gorm.io/gorm/clause"
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
	err = global.Db.Save(dto).Error
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

func Run(id int) (err error) {
	return
}
