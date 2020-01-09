package processing

import (
	"github.com/go-extras/colorcnt/pkg/image"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
)

type WorkerTestSuite struct {
	suite.Suite
}

func TestWorker(t *testing.T) {
	suite.Run(t, new(WorkerTestSuite))
}

func (s *WorkerTestSuite) TestRun() {
	url := "http://i.imgur.com/FApqk3D.jpg"

	input := make(chan string)
	output := make(chan *Result)
	w := NewWorker(1, input, output, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		w.Run()
		s.True(true)
		close(output)
		wg.Done()
	}()
	input <- url
	res := <-output
	s.True(res.Ok)
	s.Equal(url, res.Url)
	s.Equal(uint(3), res.Count)
	s.Len(res.TopColors, 3)
	s.Equal([]image.RGBColor{{0xFF, 0xFF, 0xFF}, {0x00, 0x00, 0x00}, {0xF3, 0xC3, 0x00}}, res.TopColors)
	close(input)
	wg.Wait()
}
