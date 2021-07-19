package json

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/emurray647/run_log/core/format"
	"github.com/emurray647/run_log/core/format/fit"
	"github.com/emurray647/run_log/core/format/fit/generated"
)

type JsonWriter struct {
}

// Go is stupid
// type Activity fit.Activity

func (jw *JsonWriter) Write(w io.Writer, activity *fit.Activity, options ...format.FilterOption) error {

	var builder strings.Builder

	builder.WriteString("\"Messages\": ")
	builder.WriteByte('[')

	// TODO: there is probably a better way to filter out unknown messages
	// messages := make([]fit.DataMessage, 0)
	// for _, message := range activity.Messages {
	// 	if message.Data != nil {
	// 		messages = append(messages, message)
	// 	}
	// }
	messages := activity.Messages

	for i, message := range messages {

		if message.Data != nil {
			jsonValue, err := marshalMessage(message, options...)
			if err != nil {
				return err
			}
			builder.Write(jsonValue)
			if i != len(messages)-1 {
				builder.WriteByte(',')
			}
		}
	}
	builder.WriteByte(']')

	_, err := w.Write([]byte("{" + builder.String() + "}"))
	return err
}

func marshalMessage(message fit.DataMessage, options ...format.FilterOption) ([]byte, error) {

	m, _ := json.Marshal(message.Data)

	var a interface{}
	json.Unmarshal(m, &a)
	b := a.(map[string]interface{})

	for _, opt := range options {
		if opt.MessageID == message.GlobalMessageNum || opt.MessageID == format.FilterOptionAllIDs {
			if val, ok := b[opt.FieldName]; ok {
				b[opt.FieldName] = opt.ConvertFun(val)
			}
		}
	}

	out := map[string]interface{}{}
	out["MessageType"] = generated.MessageIDToString[message.GlobalMessageNum]
	out["ID"] = message.GlobalMessageNum
	out["Data"] = b

	return json.Marshal(out)
}

func MarshalJSON(activity *fit.Activity) ([]byte, error) {

	type Alias fit.Activity
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(activity),
	})

}
