package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HomeScreen(onSend func(emotion string)) fyne.CanvasObject {

	title := widget.NewLabelWithStyle("PairEmotion", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	heartbtn := widget.NewButton("❤️", func() { onSend("heart") })
	smileBtn := widget.NewButton("😊", func() { onSend("smile") })
	pocerbtn := widget.NewButton("-_-", func() { onSend("pocer") })
	kissBtn := widget.NewButton("😘", func() { onSend("kiss") })

	grid := container.NewGridWithRows(2,
		container.NewHBox(heartbtn, smileBtn),
		container.NewHBox(pocerbtn, kissBtn),
	)

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewLabel("emotion"),
		grid,
	)

	return content
}
