package app

import (
	"context"
	"path/filepath"

	"github.com/itzmeanjan/f2d/app/config"
	"github.com/itzmeanjan/f2d/app/data"
)

// SetUp - Do basic ground set up work, required for
// running `f2d` on this machine
func SetUp(ctx context.Context) (*data.Resources, error) {

	path, err := filepath.Abs("./.env")
	if err != nil {
		return nil, err
	}

	if err := config.Read(path); err != nil {
		return nil, err
	}

	return data.Acquire(ctx)

}
