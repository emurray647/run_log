package json

import (
	"time"

	"github.com/emurray647/run_log/core/format"
	"github.com/emurray647/run_log/core/format/fit/generated"
)

var DefaultJSONFilterOptions []format.FilterOption = []format.FilterOption{
	{
		MessageID: generated.RecordDataMessageID,
		FieldName: "position_lat",
		ConvertFun: func(val interface{}) interface{} {
			return generated.Convert(val, generated.RECORD_POSITION_LAT_DEFAULT_UNIT, generated.UNIT_DEGREES)
		},
	},
	{
		MessageID: generated.RecordDataMessageID,
		FieldName: "position_long",
		ConvertFun: func(val interface{}) interface{} {
			return generated.ConvertEnum(val, generated.UNIT_SEMICIRCLES, generated.UNIT_DEGREES)
		},
	},
	{
		MessageID: format.FilterOptionAllIDs, //generated.RecordDataMessageID,
		FieldName: "timestamp",
		ConvertFun: func(val interface{}) interface{} {
			t, _ := time.Parse("2006-01-02 15:04:05", "1989-12-31 00:00:00")
			return uint32(t.Unix()) + uint32(val.(float64))
		},
	},
}
