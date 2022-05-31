package util

import (
	"bytes"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
)

func Frame4Video(r io.Reader) (io.Reader, error) {
	// 视频存到临时文件
	f, err := os.CreateTemp(os.TempDir(), "ffmpeg_frame4video_*")
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(f, r)
	if err != nil {
		return nil, err
	}

	// ffmpeg 转换，返回流
	buf := bytes.NewBuffer(nil)
	err = ffmpeg_go.Input(f.Name()).
		Filter("select", ffmpeg_go.Args{"gte(n,1)"}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}

	return buf, nil
}
