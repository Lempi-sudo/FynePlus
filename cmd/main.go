package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	mywidget "github.com/Lempi-sudo/FynePlus/widget"
)

func main() {
	a := app.NewWithID("12")
	myWindow := a.NewWindow("Button Widget")
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.SetFixedSize(true)

	selectFileBtn := widget.NewButton("Загрузить файл ...", func() {
		second_w := a.NewWindow("Выбор файла")
		second_w.Resize(fyne.NewSize(float32(1000), float32(1000)))

		myWindow.Hide()
		second_w.Show()
		myWindow.SetMaster()

		dlg := mywidget.NewFileOpenPlus(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				log.Printf("UI ERROR: Can't choose file %e\n", err)
				myWindow.Show()
				return
			}
			if uc != nil {
				log.Printf("UI INFO: Choosed file %s\n", uc.URI().Path())
				second_w.Close()
				myWindow.Show()
			}

		}, second_w)
		dlg.SetFilter(storage.NewExtensionFileFilter([]string{".pst"}))
		dlg.Resize(fyne.NewSize(float32(2500), float32(2500)))

		dlg.SetOnCancel(func() {
			myWindow.Show()
			second_w.Close()
		})
		second_w.SetCloseIntercept(func() {
			myWindow.Show()
			second_w.Close()
		})
		dlg.Show()
	})

	myWindow.SetFixedSize(false)
	myWindow.SetContent(selectFileBtn)
	myWindow.ShowAndRun()
}
