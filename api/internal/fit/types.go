package fit

import "github.com/emurray647/run_log/core/format/fit/generated"

type Header struct {
	size            uint8
	protocolVersion uint8
	profileVersion  uint16
	dataSize        uint32
	// ".Fit"
	crc uint16
}

type RecordHeaderByte struct {
	timeCompressed bool // 0 for normal header
	messageType    bool // 0 for DataMessage, 1 for DefinitionMessage
	// messageTypeSpecific   bool
	developerFlagsEnabled bool
	localMessageType      uint8 // 4 bit field
}

func (r RecordHeaderByte) isDefinitionMessage() bool {
	return r.messageType
}

type RecordDefinitionMessage struct {
	header              RecordHeaderByte
	architecture        uint8 // 0 -> little endian
	globalMessageNumber uint16
	numFields           uint8
	fields              []FieldDefinition
	numDeveloperFields  uint8
	developerFields     []FieldDefinition
}

type FieldDefinition struct {
	fieldDefinitionNumber uint8
	size                  uint8
	baseType              uint8
}

type TypeMap map[uint8]*RecordDefinitionMessage

type DataMessage struct {
	header           RecordHeaderByte
	GlobalMessageNum uint
	Data             generated.DataMessageData
}
