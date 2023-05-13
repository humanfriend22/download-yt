package main

import (
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	fyneDialog "fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

func main() {
	a := app.New()
	w = a.NewWindow("Download YT")

	// Variables
	pathLabel := widget.NewLabel("")
	status := canvas.NewText("", color.RGBA{255, 165, 0, 255})

	// Input Data
	formats, format := []string{"Video Only", "Audio Only", "Video & Audio (Default)"}, 2
	// mediaQualities, videoQuality, audioQuality := []string{"Low", "Medium", "High"}, 2, 2

	// Inputs
	url := widget.NewEntry()

	// videoQualitySelect := &widget.Select{
	// 	Selected: mediaQualities[2],
	// 	Options:  mediaQualities,
	// 	OnChanged: func(value string) {
	// 		for index, option := range mediaQualities {
	// 			if option == value {
	// 				videoQuality = index
	// 			}
	// 		}
	// 	},
	// }

	// audioQualitySelect := &widget.Select{
	// 	Selected: mediaQualities[2],
	// 	Options:  mediaQualities,
	// 	OnChanged: func(value string) {
	// 		for index, option := range mediaQualities {
	// 			if option == value {
	// 				audioQuality = index
	// 			}
	// 		}
	// 	},
	// }

	formatSelect := &widget.Select{
		Selected: formats[2],
		Options:  formats,
		OnChanged: func(value string) {
			for index, option := range formats {
				if option == value {
					// if videoQualitySelect.Disabled() && audioQualitySelect.Disabled() {
					// 	switch value {
					// 	case mediaQualities[0]:
					// 		{
					// 			videoQualitySelect.Enable()
					// 			audioQualitySelect.Disable()
					// 		}

					// 	case mediaQualities[1]:
					// 		{
					// 			audioQualitySelect.Enable()
					// 			videoQualitySelect.Disable()
					// 		}

					// 	default:
					// 		{
					// 			videoQualitySelect.Enable()
					// 			audioQualitySelect.Enable()
					// 		}
					// 	}
					// }

					format = index
				}
			}
		},
	}

	// Submit Button
	var button *widget.Button
	button = &widget.Button{
		Text: "Download",
		OnTapped: func() {
			status.Color = color.RGBA{200, 100, 0, 255}
			status.Text = "Verifying..."
			if url.Text != "" {
				status.Text = "Choosing Directory..."
				dir, err := dialog.Directory().Title("Select a folder").Browse()

				if err != nil {
					if err.Error() != "Cancelled" {
						fyneDialog.ShowError(err, w)
					}

					return
				}

				button.Disable()
				pathLabel.SetText(dir)

				status.Text = "Downloading..."
				status.Color = color.RGBA{0, 0, 255, 200}

				final := DownloadVideo(url.Text, dir, format != 2)
				n := FormatVideo(format, final)
				open(final + n)
				//format, videoQuality, audioQuality
				// if err != nil && err.Error() == "invalid url or id" {
				// 	status.Text = "Invalid url or id"
				// 	status.Color = color.RGBA{255, 0, 0, 200}
				// 	return
				// }

				status.Text = ""
				button.Enable()
			}
		},
	}

	// Form
	var form *widget.Form
	form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Youtube URL", Widget: url},
			{Text: "Format", Widget: formatSelect},
			// {Text: "Video Quality", Widget: videoQualitySelect},
			// {Text: "Audio Quality", Widget: audioQualitySelect},
			{Widget: button},
		},

		OnSubmit: nil,
	}

	// Check for FFmpeg support
	if !IsFFmpegAvailable() {
		formatSelect.Disable()
		// videoQualitySelect.Disable()
		// audioQualitySelect.Disable()
	}

	// Render
	w.SetContent(
		container.NewBorder(
			container.NewVBox(
				container.NewPadded(form),
				// button,
			),
			container.NewVBox(
				container.NewPadded(
					container.NewCenter(
						status,
					),
				),
			),
			nil,
			nil,
		),
	)

	w.SetOnClosed(func() {
		os.Remove("temp.mp4")
		a.Quit()
	})

	// Run
	w.Resize(fyne.NewSize(600, 500))
	w.ShowAndRun()
}
