package util

import (
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFrame4Video(t *testing.T) {
	flag.Parse()
	videoPath := flag.Arg(0)
	fmt.Printf("输入视频文件：%s\n", videoPath)
	f, err := os.Open(videoPath)
	if err != nil {
		t.Fatal(err)
		return
	}
	r, err := Frame4Video(f)
	if err != nil {
		t.Fatal(err)
		return
	}

	f, err = os.CreateTemp(os.TempDir(), "douyin_test_*.jpeg")
	//defer func() {
	//	f.Close()
	//	os.Remove(f.Name())
	//}()
	if err != nil {
		t.Fatal(err)
		return
	}
	_, err = io.Copy(f, r)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("图片输出到：%s\n", f.Name())
}
