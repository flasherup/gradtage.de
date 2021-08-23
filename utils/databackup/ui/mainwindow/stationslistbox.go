package mainwindow

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/asaskevich/EventBus"
)

type StationsListBox struct {
	Status  *widget.Label
	LoadBtn *widget.Button
	Bus     EventBus.Bus
}

func NewStationListBox(ct *fyne.Container, bus EventBus.Bus) *StationsListBox {
	stl := StationsListBox{
		Bus: bus,
	}

	stl.Status = widget.NewLabel("Load Stations")
	stl.LoadBtn = widget.NewButton("Load", stl.onLoadStations)

	ct.Add(container.NewVBox(
		stl.Status,
		stl.LoadBtn,
	))
	return &stl
}

func (stl *StationsListBox) onLoadStations() {
	stl.Bus.Publish(OnStationsLoad)
}
