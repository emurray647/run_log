package format

import (
	"io"

	"github.com/emurray647/run_log/core/format/fit"
)

// type FilterOption map[uint]map[string]func(uint) uint
type FilterOption struct {
	MessageID  uint
	FieldName  string
	ConvertFun func(interface{}) interface{}
}

const FilterOptionAllIDs uint = 0xFFFF

type Writer interface {
	Write(io.Writer, *fit.Activity)
}
