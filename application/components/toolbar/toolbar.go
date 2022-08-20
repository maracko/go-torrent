package toolbar

import (
	"log"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// type Toolbar struct {
// 	Window fyne.Window
// 	Run  func(fyne.Window) fyne.CanvasObject
// }

var (
	Toolbar = widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New torrent")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
)
