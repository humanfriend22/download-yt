package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsFFmpegAvailable() bool {
	err := exec.Command("ffmpeg").Run()
	return err != nil
}

func FormatVideo(formatI int, final string) string {
	argsTemplate := "-i temp.mp4 -%sn " + final + "%s"
	var args string
	if formatI == 0 {
		args = fmt.Sprintf(argsTemplate, "a", "4")
	} else {
		args = fmt.Sprintf(argsTemplate, "v", "3")
	}

	command := exec.Command("ffmpeg", strings.Split(args, " ")...)
	command.Run()

	return args[len(args)-1:]
}
