package processing

import (
	"github.com/go-extras/colorcnt/pkg/image"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Result struct {
	Url       string
	TopColors []image.RGBColor
	Count     uint
	Ok        bool
}

type Worker struct {
	id        int
	input     <-chan string
	output    chan<- *Result
	maxColors uint
}

func NewWorker(id int, input <-chan string, output chan<- *Result, maxColors uint) *Worker {
	return &Worker{
		id:        id,
		input:     input,
		output:    output,
		maxColors: maxColors,
	}
}

func (w *Worker) process(url string) *Result {
	log.Infof("worker[%02d]: processing %s", w.id, url)

	result := &Result{
		Url: url,
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Warnf("worker[%02d]: error reading from url: %s", w.id, err.Error())
		return result
	}
	defer resp.Body.Close()
	img := image.NewImage(resp.Body)
	colors, err := img.GetRGBColorMap()
	if err != nil {
		log.Infof("worker[%02d]: error processing image: %s", w.id, err.Error())
		return result
	}
	result.Ok = true
	result.TopColors, result.Count = colors.MaxUsedColors(w.maxColors)

	return result
}

func (w *Worker) Run() {
	for c := range w.input {
		w.output <- w.process(c)
	}
}
