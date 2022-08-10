package SSH

import (
	"github.com/xtclalala/ScanNetWeb/global"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/model/SSH"
)

func SelectResult(dto *SSH.SearchSSHResult) (data *SSH.BizSSHResult, err error) {
	err = global.Db.Model(&SSH.BizSSHResult{ResultId: dto.ResultId}).Find(&data).Error
	err = proError.WrapOrNil(err, "Search ssh task result: %s fail!", dto.ResultId)
	return
}
