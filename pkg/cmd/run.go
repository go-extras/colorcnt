package cmd

import (
	"context"
	"encoding/csv"
	"github.com/go-extras/colorcnt/pkg/processing"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)

type RunCommand struct {
	InputFile    string `long:"input-file" required:"yes" description:"Input file name (must consist of URL)"`
	OutputFile   string `long:"output-file" required:"yes" description:"Output file name"`
	MaxWorkers   uint   `long:"max-workers" default:"10" description:"Max workers to run concurrently"`
	MaxTopColors uint   `long:"max-top-colors" default:"3" description:"Max number of the top colors to choose"`
	WriteBuffer  uint   `long:"write-buffer" default:"10000" description:"Size of CSVWriter buffer (in records)"`
}

func RegisterRunCommand(parser *flags.Parser) *RunCommand {
	cmd := &RunCommand{}
	parser.AddCommand("run", "runs main processor", "", cmd)
	return cmd
}

func (cmd *RunCommand) Execute(args []string) error {
	log.Info("Starting the processor")
	inputFile, err := os.Open(cmd.InputFile)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	ctx := context.Background()
	inputCh := make(chan string)
	outputCh := make(chan *processing.Result, cmd.WriteBuffer)
	reader := processing.NewReader(ctx, inputFile, inputCh)

	var wg, writerWG sync.WaitGroup
	wg.Add(int(cmd.MaxWorkers))
	writerWG.Add(1)

	go func() {
		lines, err := reader.Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("Reader stopped, %d lines read", lines)
	}()

	outputFile, err := os.Create(cmd.OutputFile)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(outputFile)
	writer := processing.NewWriter(outputCh, csvWriter, cmd.MaxTopColors)
	go func() {
		err = writer.Run()
		if err != nil {
			log.Fatal(err)
		}
		writerWG.Done()
	}()

	for n := 1; n <= int(cmd.MaxWorkers); n++ {
		go func(n int) {
			w := processing.NewWorker(n, inputCh, outputCh, cmd.MaxTopColors)
			w.Run()
			wg.Done()
		}(n)
	}
	wg.Wait()
	log.Debug("All workers stopped")
	close(outputCh)
	writerWG.Wait()
	log.Debug("CSV writer stopped")
	csvWriter.Flush()
	outputFile.Close()

	return nil
}
