package tools

import (
	"github.com/xtclalala/ScanNetWeb/constant"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"net/http"
)

func Readfile(fileName, sheet string) (rows [][]string, err error) {
	file, err := excelize.OpenFile(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	// 获取 Sheet1 上所有单元格
	rows, err = file.GetRows(sheet)
	return
}

func GetFileType(out multipart.File) (fileType string, err error) {
	buffer := make([]byte, 512)
	if _, err = out.Read(buffer); err != nil {
		return "", err
	}
	fileType = http.DetectContentType(buffer)
	return
}

func FileAbsPath(path, name string) string {
	return path + constant.Slash + name
}
