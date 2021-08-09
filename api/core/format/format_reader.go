package format

import (
	"io"

	"github.com/emurray647/run_log/core/format/fit"
	// "github.com/emurray647/run_log/core/models"
)

type ReadFilter struct {
	
}

type Reader interface {
	Read(io.Reader) *fit.Activity
}
