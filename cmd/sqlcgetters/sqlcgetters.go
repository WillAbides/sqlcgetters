package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"github.com/willabides/sqlcgetters"
)

type cmd struct {
	Path []string `kong:"arg,help='Path to sqlc output file. Can be glob. Default: \"*.go\"',default='*.go'"`
	Out  string   `kong:"short=o,help='Output file. Default: stdout'"`
}

func main() {
	var cli cmd
	kctx := kong.Parse(&cli)
	kctx.FatalIfErrorf(kctx.Run())
}

func (c *cmd) Run(kctx *kong.Context) (err error) {
	if len(c.Path) == 0 {
		c.Path = []string{"*.go"}
	}

	var buf bytes.Buffer
	_, err = sqlcgetters.GenerateGetters(&buf, os.DirFS("."), c.Path...)
	if err != nil {
		return err
	}

	// default to stdout
	if c.Out == "" {
		_, err = io.Copy(kctx.Stdout, &buf)
		return err
	}

	// write file with permissions from parent dir minus executable bit
	dirInfo, err := os.Stat(filepath.Dir(c.Out))
	if err != nil {
		return err
	}
	return os.WriteFile(c.Out, buf.Bytes(), dirInfo.Mode()&^0o111)
}
