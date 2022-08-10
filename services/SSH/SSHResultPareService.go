package SSH

import (
	"github.com/xtclalala/ScanNetWeb/global"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/model/SSH"
)

func CreateResultPare(dtos []*SSH.BizSSHResultParse) (err error) {
	err = global.Db.Create(&dtos).Error
	return proError.WrapOrNil(err, "Create ssh task result fail!")
}
