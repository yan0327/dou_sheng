package upload

/*
用于对上传视频的保存
*/

import (
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"simple-demo/global"
	"simple-demo/pkg/util"
	"strconv"
	"strings"
	"time"
)

type FileType int

func GetFileName(name string, username string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	rand.Seed(time.Now().Unix())
	strn := strconv.Itoa(rand.Intn(math.MaxInt8))
	fileName += username + strn
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func GetServerUrl() string {
	return global.AppSetting.UploadServerUrl
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsNotExist(err)
}

func CheckContainExt(name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch ext {
	case ".image":
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	case ".MP4":
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(name string, f multipart.File) bool {
	ext := strings.ToUpper(GetFileExt(name))
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch ext {
	case ".MP4":
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

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
