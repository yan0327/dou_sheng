package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	inputs := "storage/uploadsVideo/447d482077d6db1dc4afc22124a99ac5.mp4"
	imagedst := strings.Replace(inputs, ".mp4", ".jpg", 1)
	imagedst, _ = filepath.Abs(strings.Replace(imagedst, "uploadsVideo", "uploadsImage", 1))
	cmdArguments := []string{"-i", inputs, "-y", "-f", "image2", "-t", "1", "-s", "1364x900", imagedst}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	err := cmd.Run()
	fmt.Println(err)
}
