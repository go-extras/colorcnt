package main

import (
	"fmt"
	"github.com/go-extras/colorcnt/pkg/cmd"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	parser := flags.NewParser(nil, flags.Default)
	cmd.RegisterRunCommand(parser)

	parser.CommandHandler = func(command flags.Commander, args []string) error {
		err := command.Execute(args)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		return nil
	}
	_, err := parser.Parse()
	if err != nil {
		os.Exit(-1)
	}
}
