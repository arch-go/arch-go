package commands

import "io"

type BaseCommand struct {
	Output io.Writer
}

type Command interface {
	Run()
}
