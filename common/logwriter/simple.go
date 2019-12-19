package logwriter

import "fmt"

type LogWriter struct {

}

func (d LogWriter)Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	n = len(p)
	return
}