package mainwindow

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/asaskevich/EventBus"
)

const (
	OnStationsLoad = "OnStationsLoad"
	OnBackupStart = "OnBackupStart"
)

type WindowMain struct {
	App    *fyne.App
	Window fyne.Window
	Bus    EventBus.Bus
	Slb    StationsListBox
	Bkb    BackupBox
}

func NewWindowMain(a fyne.App) *WindowMain {
	wm := WindowMain{
		App:           &a,
		Window:        a.NewWindow("Data Backup"),
		Bus:           EventBus.New(),
	}

	ct := container.New(layout.NewVBoxLayout())
	wm.Slb = *NewStationListBox(ct, wm.Bus)
	ct.Add(layout.NewSpacer())
	wm.Bkb = *NewBackupBox(ct, wm.Bus)
	wm.Window.SetContent(ct)
	wm.Window.Resize(fyne.NewSize(300, 500))
	return &wm
}

func (wm *WindowMain) SetStationsLength(l int) {
	wm.Slb.Status.SetText(fmt.Sprintf("%d stations loaded", l))
	wm.Slb.LoadBtn.SetText("Reload")
	wm.Window.Content().Refresh()
}
