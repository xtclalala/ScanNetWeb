package SSH

import (
	"github.com/xtclalala/ScanNetWeb/global"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/model/SSH"
	"gorm.io/gorm/clause"
)

func CreateResultPare(dtos []*SSH.BizSSHResultParse) (err error) {
	err = global.Db.Create(&dtos).Error
	return proError.WrapOrNil(err, "Create ssh task result fail!")
}

func SearchResultParse(dto *SSH.SearchSSHResultParse) (list []SSH.BizSSHResultParse, total int64, err error) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&SSH.BizSSHResultParse{}).Where(&SSH.BizSSHResultParse{
		TaskId: dto.TaskId,
	})

	if dto.Os != "" {
		db.Where("os like ?", "%"+dto.Os+"%")
	}

	oc := clause.OrderByColumn{
		Column: clause.Column{Name: "create_time", Raw: true},
		Desc:   dto.Desc,
	}
	err = db.Count(&total).Limit(limit).Offset(offset).Order(oc).Find(&list).Error
	err = proError.WrapOrNil(err, "Search ssh task result fail!")
	return
}
