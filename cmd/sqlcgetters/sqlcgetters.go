package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/willabides/sqlcgetters"
)

type cmd struct {
	Path []string `kong:"arg,help='Path to sqlc output file. Can be glob. Default: \"*.go\"',default='*.go'"`
}

func main() {
	var cli cmd
	kctx := kong.Parse(&cli)
	kctx.FatalIfErrorf(kctx.Run())
}

func (c *cmd) Run(kctx *kong.Context) error {
	_, err := sqlcgetters.GenerateGetters(kctx.Stdout, os.DirFS("."), c.Path...)
	return err
}
