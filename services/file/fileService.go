package file

import (
	"github.com/google/uuid"
	"github.com/xtclalala/ScanNetWeb/global"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/model/file"
)

func Create(dtos []*file.BizFile) (err error) {
	err = global.Db.Create(&dtos).Error
	return proError.WrapOrNil(err, "create file: %s is fail", dtos)
}

func FindById(id uuid.UUID) (data file.BizFile, err error) {
	err = global.Db.First(&data, id).Error
	return data, proError.WrapOrNil(err, "find file: %s is fail", id)
}
