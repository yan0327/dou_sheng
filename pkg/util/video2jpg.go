package util

import "os/exec"

func Video2JPG(input string, output string) {
	cmdArguments := []string{"-i", input, "-y", "-f",
		"image2", "-t", "1", "-s", "1364x900", output}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	_ = cmd.Run()
}
