package processing

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
)

type Writer struct {
	input     <-chan *Result
	output    *csv.Writer
	maxColors uint
}

func NewWriter(input <-chan *Result, output *csv.Writer, maxColors uint) *Writer {
	return &Writer{
		input:     input,
		output:    output,
		maxColors: maxColors,
	}
}

func (w *Writer) Run() error {
	for item := range w.input {
		if !item.Ok {
			log.Warnf("[csvwriter] Skipping failed url %s", item.Url)
			continue
		}
		strs := make([]string, w.maxColors+1, w.maxColors+1)
		strs[0] = item.Url
		for i := uint(0); i < w.maxColors; i++ {
			strs[i+1] = item.TopColors[i].String()
		}
		err := w.output.Write(strs)
		if err != nil {
			return err
		}
	}

	return nil
}
