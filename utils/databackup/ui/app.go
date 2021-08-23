package ui

import(
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/flasherup/gradtage.de/utils/databackup/ui/mainwindow"
)

type Application struct {
	App fyne.App
	WMain *mainwindow.WindowMain
}

func NewApplication() *Application {
	a := app.New()
	w := mainwindow.NewWindowMain(a)
	return &Application{
		App: a,
		WMain: w,
	}
}