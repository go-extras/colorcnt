package processing

import (
	"bufio"
	"bytes"
	"context"
	"io"
)

type Reader struct {
	input io.Reader
	ctx   context.Context
	out   chan<- string
}

func NewReader(ctx context.Context, input io.Reader, out chan<- string) *Reader {
	return &Reader{
		input: input,
		ctx:   ctx,
		out:   out,
	}
}

func (p *Reader) Run() (lines int, err error) {
	// Start reading from the file with a reader.
	reader := bufio.NewReader(p.input)

loop:
	for {
		select {
		case <-p.ctx.Done():
			break loop // cancelled
		default:
		}

		var buffer bytes.Buffer

		var l []byte
		var isPrefix bool
		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				lines += 1
				break
			}

			// If we're just at the EOF, break
			if err != nil {
				break
			}
		}

		if err == io.EOF {
			err = nil
			break
		}

		p.out <- buffer.String()
	}

	close(p.out)

	return lines, err
}
