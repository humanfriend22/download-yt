package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/kkdai/youtube/v2"
)

func open(file string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer.exe", "/select,", file)
	case "darwin":
		cmd = exec.Command("open", "-R", file)
	case "linux":
		cmd = exec.Command("xdg-open", file)
	}

	err := cmd.Run()
	if err != nil {
		Throw(err)
	}
}

func DownloadVideo(url string, folder string, temp bool) string {

	// Get Video Metadata
	client := youtube.Client{}
	video, err := client.GetVideo(url)
	if err != nil {
		if strings.HasSuffix(err.Error(), "This video is unavailable") || strings.HasSuffix(err.Error(), "the video id must be at least 10 characters long") {
			Throw(errors.New("invalid url or id"))
		}
		Throw(err)
	}

	// Video Download Stream
	stream, _, err := client.GetStream(video, &video.Formats.WithAudioChannels()[0])
	if err != nil {
		Throw(err)
	}

	// Donwload Destination
	slashes := regexp.MustCompile(`[\/\\]`)
	finalFilename := strings.Join(
		strings.Split(strings.ToLower(
			strings.ReplaceAll(
				slashes.ReplaceAllString(video.Title, ""),
				"&",
				"",
			)), " "), "_") + ".mp"
	finalFilepath := path.Join(folder, finalFilename)

	filename := "temp.mp4"
	if !temp {
		filename = finalFilepath + "4"
	}

	// Create File
	file, err := os.Create(filename)
	if err != nil {
		Throw(err)
	}
	defer file.Close()

	// Download
	_, err = io.Copy(file, stream)
	if err != nil {
		Throw(err)
	}

	// fmt.Println(`Finished downloading "` + video.Title + `"!`)

	err = beeep.Alert("Download Completed!", "Your Youtube video/audio has been downloaded", "/Users/humanfriend22/dev/download-yt-go/youtube-logo.png")
	if err != nil {
		panic(err)
	}

	// notify := notificator.New(notificator.Options{
	// 	DefaultIcon: "youtube-logo.png",
	// 	AppName:     "Youtube Downloader",
	// })

	// notify.Push("Download Completed!", "Your Youtube video/audio has been downloaded.", "youtube-logo.png", notificator.UR_NORMAL)

	return finalFilepath
}
