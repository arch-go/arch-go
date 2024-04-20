package commands

import "io"

type BaseCommand struct {
	Output io.Writer
}
