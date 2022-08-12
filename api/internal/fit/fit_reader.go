package fit

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/emurray647/run_log/internal/fit/generated"
	"github.com/emurray647/run_log/internal/model"
	// "github.com/emurray647/run_log/internal/activities"
)

//go:generate go run generator/generate.go

type FitReader struct {
	r         *io.Reader
	bytesRead uint
}

func (fr *FitReader) Read(r io.Reader, opts ...func(DataMessage) bool) []model.Activity {

	fr.r = &r
	fr.bytesRead = 0

	var messages []DataMessage

	header, _ := fr.readHeader()

	typeMap := TypeMap{}

	countMap := make(map[uint]uint)

	for fr.bytesRead < uint(header.dataSize-uint32(header.size)-2) {
		headerByte, err := fr.readRecordHeaderByte()
		if err != nil {
			// log.Fatalf(err.Error())
			panic(err.Error())
		}

		if headerByte.isDefinitionMessage() {
			recordDefinitionMessage, _ := fr.readRecordDefinitionMessage(headerByte)
			typeMap[recordDefinitionMessage.header.localMessageType] = &recordDefinitionMessage

		} else {
			dataMessage, _ := fr.readRecordDataMessage(headerByte, typeMap)

			if generated.IsValidMessageID(dataMessage.GlobalMessageNum) {
				accept := true
				for _, opt := range opts {
					accept = accept && opt(dataMessage)
				}
				if accept {
					messages = append(messages, dataMessage)
				}
			}

			countMap[dataMessage.GlobalMessageNum] += 1

		}
	}

	activity := model.Activity{}

	for _, msg := range messages {
		if msg.GlobalMessageNum == generated.SessionDataMessageID {
			session := msg.Data.(*generated.SessionDataMessage)
			t, _ := time.Parse("2006-01-02 15:04:05", "1989-12-31 00:00:00")
			startTime := uint32(t.Unix()) + uint32(*session.StartTime)
			activity.ActivitySummary = model.ActivitySummary{
				StartTime:     int(startTime),
				TotalTime:     float32(generated.Convert(*session.TotalTimerTime, generated.SESSION_TOTAL_TIMER_TIME_DEFAULT_UNIT, generated.UNIT_SECOND)),
				TotalDistance: float32(generated.Convert(*session.TotalDistance, generated.SESSION_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER)),
				AvgHeartRate:  uint8(*session.AvgHeartRate),
				AvgCadence:    uint8(*session.AvgCadence),
			}
		}

	}

	for _, msg := range messages {
		if msg.GlobalMessageNum == generated.RecordDataMessageID {
			activity.Records = append(activity.Records, *convertToActivityRecord(msg.Data.(*generated.RecordDataMessage), &activity.ActivitySummary))
		}

		if msg.GlobalMessageNum == generated.LapDataMessageID {
			activity.Laps = append(activity.Laps, *convertToActivityLap(msg.Data.(*generated.LapDataMessage), &activity.ActivitySummary))
		}
	}

	// todo: it is technically possible for multiple activities to come out of this
	return []model.Activity{activity}
}

func (r *FitReader) readHeader() (Header, error) {
	header := Header{}
	headerBuffer, err := r.readBytes(1)
	if err != nil {
		return header, err
	}
	header.size = headerBuffer[0]
	headerBuffer, err = r.readBytes(uint(header.size - 1))
	if err != nil {
		return header, err
	}

	header.protocolVersion = headerBuffer[0]
	header.profileVersion = binary.LittleEndian.Uint16(headerBuffer[1:3])
	header.dataSize = binary.LittleEndian.Uint32(headerBuffer[3:7])

	if header.size == 16 {
		header.crc = binary.LittleEndian.Uint16(headerBuffer[11:])
	}

	return header, err
}

func (r *FitReader) readRecordHeaderByte() (RecordHeaderByte, error) {
	recordHeaderByte := RecordHeaderByte{}

	buffer, err := r.readBytes(1)
	if err != nil {
		return recordHeaderByte, err
	}

	byteValue := buffer[0]
	recordHeaderByte.timeCompressed = (byteValue & 0x80) != 0
	if !recordHeaderByte.timeCompressed {
		recordHeaderByte.messageType = (byteValue & 0x40) != 0
		if recordHeaderByte.isDefinitionMessage() {
			recordHeaderByte.developerFlagsEnabled = (byteValue & 0x20) != 0
		}
		recordHeaderByte.localMessageType = (byteValue & 0x0F)
	} else {
		recordHeaderByte.localMessageType = (byteValue & 0x60) >> 5
		recordHeaderByte.messageType = false
	}

	return recordHeaderByte, nil
}

func (r *FitReader) readRecordDefinitionMessage(headerByte RecordHeaderByte) (RecordDefinitionMessage, error) {
	definitionMessage := RecordDefinitionMessage{header: headerByte}

	buffer, err := r.readBytes(4)
	if err != nil {
		return definitionMessage, err
	}

	definitionMessage.architecture = buffer[1]

	var byteOrder binary.ByteOrder
	if definitionMessage.architecture == 0 {
		byteOrder = binary.LittleEndian
	} else {
		byteOrder = binary.BigEndian
	}

	definitionMessage.globalMessageNumber = byteOrder.Uint16(buffer[2:])

	buffer, err = r.readBytes(1)
	if err != nil {
		return definitionMessage, err
	}
	definitionMessage.numFields = buffer[0]

	definitionMessage.fields = make([]FieldDefinition, definitionMessage.numFields)
	buffer, err = r.readBytes(uint(3 * definitionMessage.numFields))
	if err != nil {
		return definitionMessage, err
	}

	for i := uint8(0); i < definitionMessage.numFields; i++ {
		bufferIdx := i * 3
		definitionMessage.fields[i].fieldDefinitionNumber = buffer[bufferIdx]
		definitionMessage.fields[i].size = buffer[bufferIdx+1]
		definitionMessage.fields[i].baseType = buffer[bufferIdx+2]
	}

	// TODO: developer fields
	if definitionMessage.header.developerFlagsEnabled {
		buffer, err = r.readBytes(1)
		if err != nil {
			return definitionMessage, err
		}
		definitionMessage.numDeveloperFields = buffer[0]
	}

	return definitionMessage, nil
}

func (r *FitReader) readRecordDataMessage(header RecordHeaderByte, types TypeMap) (DataMessage, error) {
	record := DataMessage{header: header}

	// use the localMessageType to lookup the global message num
	definitionMessage := types[header.localMessageType]
	record.GlobalMessageNum = uint(definitionMessage.globalMessageNumber)
	// use the global message number to create the right type
	messageData, _ := generated.CreateMessageData(record.GlobalMessageNum)

	for _, field := range definitionMessage.fields {
		buffer, _ := r.readBytes(uint(field.size))
		if messageData != nil {
			messageData.Read(uint(field.fieldDefinitionNumber), buffer, binary.LittleEndian)
		}

	}
	record.Data = messageData

	return record, nil
}

func (fr *FitReader) readBytes(n uint) ([]byte, error) {
	buffer := make([]byte, n)
	bytesRead, err := (*fr.r).Read(buffer)
	fr.bytesRead = fr.bytesRead + uint(bytesRead)
	if err != nil {
		return buffer, err
	} else if uint(bytesRead) != n {
		return buffer, fmt.Errorf("could not read %d bytes", n)
	}
	return buffer, nil
}
