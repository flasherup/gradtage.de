package mainwindow

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/asaskevich/EventBus"
)

type BackupBox struct {
	Status  *widget.Label
	BackupBtn *widget.Button
	Bus     EventBus.Bus
}

func NewBackupBox(ct *fyne.Container, bus EventBus.Bus) *BackupBox {
	bck := BackupBox{
		Bus: bus,
	}

	bck.Status = widget.NewLabel("Backup")
	bck.BackupBtn = widget.NewButton("Start", bck.onBackupStart)

	ct.Add(container.NewVBox(
		bck.Status,
		bck.BackupBtn,
	))
	return &bck
}

func (bck *BackupBox) onBackupStart() {
	bck.Bus.Publish(OnBackupStart)
}
