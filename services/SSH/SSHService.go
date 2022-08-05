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
	db := global.Db.Model(&SSH.BizSSH{})

	if dto.State != constant.State {
		db = db.Where("state = ?", dto.State)
	}
	if dto.Name != "" {
		db = db.Where("name like ?", "%"+dto.Name+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Limit(limit).Offset(offset)

	oc := clause.OrderByColumn{
		Column: clause.Column{Name: "sort", Raw: true},
		Desc:   dto.Desc,
	}

	err = db.Order(oc).Find(&list).Error
	return
}

func Create(dto *SSH.BizSSH) (err error) {
	err = global.Db.Create(&dto).Error
	return proError.WrapOrNil(err, "Create ssh task: %s fail!", dto.Name)
}

func Update(dto *SSH.BizSSH) (err error) {
	var old SSH.BizSSH
	err = global.Db.First(&old, dto.Id).Error
	if err != nil {
		return proError.Wrap(err, "Update ssh task: %s fail!", dto.Id)
	}

	err = global.Db.Save(dto).Error
	return
}

func Delete(id int) (err error) {
	err = global.Db.Delete(&SSH.BizSSH{}, id).Error
	return proError.WrapOrNil(err, "Delete ssh task: %s fail!", id)
}

func Run() {

}
