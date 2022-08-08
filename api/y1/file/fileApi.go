package file

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xtclalala/ScanNetWeb/global"
	"github.com/xtclalala/ScanNetWeb/internal/net"
	"github.com/xtclalala/ScanNetWeb/internal/proError"
	"github.com/xtclalala/ScanNetWeb/model"
	"github.com/xtclalala/ScanNetWeb/model/file"
	service "github.com/xtclalala/ScanNetWeb/services/file"
	"github.com/xtclalala/ScanNetWeb/tools"
	"path"
)

// Upload 上传文件
func Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	var dataList []*file.BizFile
	dataMap := map[string]string{}

	uploadPath := global.System.File.Path
	for _, f := range files {
		id := uuid.New()
		idStr := id.String()
		b, _ := f.Open()
		defer b.Close()
		fileType, err := tools.GetFileType(b)
		if err != nil {
			net.FailWhitStatus(proError.FileReadType, c)
			return
		}
		data := &file.BizFile{
			BaseUUID: model.BaseUUID{
				Id: id,
			},
			Name: f.Filename,
			Type: fileType,
			Path: uploadPath,
		}
		dataList = append(dataList, data)
		dataMap[f.Filename] = idStr
		f.Filename = idStr
		if err = c.SaveUploadedFile(f, path.Join(uploadPath, f.Filename)); err != nil {
			net.FailWhitStatus(proError.UploadFileError, c)
			return
		}
	}
	if err := service.Create(dataList); err != nil {
		net.FailWhitStatus(proError.UploadFileError, c)
		return
	}
	net.OkWithData(dataMap, c)
}

// Download 获取文件信息
func Download(c *gin.Context) {
	fileId := c.Param("fileId")
	if len(fileId) <= 0 {
		net.FailWithMessage("文件id错误", c)
		return
	}
	fileData, err := service.FindById(uuid.MustParse(fileId))
	if err != nil {
		net.FailWhitStatus(proError.FindFileError, c)
		return
	}
	fileContentDisposition := "attachment;filename=\"" + fileData.Name + "\""
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fileContentDisposition)
	c.File(tools.FileAbsPath(fileData.Path, fileData.Id.String()))

}
