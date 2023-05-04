package upload

import (
	"blog_service/global"
	"blog_service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

func GetFileName(name string) string {
	ext := GetFileExt(name)

	//刪除後綴ext
	fileName := strings.TrimSuffix(name, ext)
	//檔案名稱加密
	fileName = util.EncodeMD5(fileName)

	//返回加密檔名+副檔名
	return fileName + ext
}

// 取得檔案副檔名
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 取得檔案存儲位址
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// 檢查目錄是否存在
func CheckSavePath(dst string) bool {
	//返回檔案描述
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// 檢查副檔名是否在約定的設定中
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	//副檔名轉換成全大寫
	ext = strings.ToUpper(ext)

	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

// 檢查檔案是否超出限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)

	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

// 檢查檔案許可權是否足夠
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)

	return err
}
