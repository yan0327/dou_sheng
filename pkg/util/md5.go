package util

import (
	"crypto/md5"
	"encoding/hex"
)

//对上传的视频名字进行格式化，避免直接暴露原始名称
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
