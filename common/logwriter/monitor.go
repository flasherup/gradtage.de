package logwriter

import (
	"fmt"
	"github.com/flasherup/gradtage.de/testsvc/client"
)

type MonitorWriter struct {
	client *client.MonitorClient
	url string
}

func NewMonitorWriter(url string) *MonitorWriter {
	return &MonitorWriter{
		client: client.NewMonitorClient(url),
		url:url,
	}
}

func (mw *MonitorWriter)Write(p []byte) (n int, err error) {
	fmt.Println("log", p)
	err = mw.client.Log(string(p))
	n = len(p)
	return
}