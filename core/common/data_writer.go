package common

import (
	"bufio"
	"encoding/binary"
	"io"
)

// DataWriter is wrapper of bufio.Writer.
type DataWriter struct {
	w *bufio.Writer
}

func NewWriter(writer io.Writer) *DataWriter {
	return &DataWriter{w: bufio.NewWriter(writer)}
}

func (dw *DataWriter) WriteByte(value byte) error {
	return dw.w.WriteByte(value)
}

// WriteVarInt writes a variable length representation of an integer value that reduces the number of written bytes for small positive values.
func (dw *DataWriter) WriteVarInt(value int) error {
	if (value & 0xFFFFFF80) == 0 {
		err := dw.w.WriteByte(byte(value))
		if err != nil {
			return err
		}
	} else {
		err := dw.w.WriteByte(byte(0x80 | (value & 0x7F)))
		if err != nil {
			return err
		}
		value = value >> 7
		err = dw.WriteVarInt(value)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteString writes a string.
func (dw *DataWriter) WriteString(value string) error {
	bytes := []rune(value)
	bytesNumber := uint16(len(bytes))
	err := binary.Write(dw.w, binary.BigEndian, bytesNumber)
	if err != nil {
		return err
	}
	runes := []rune(value)
	for _, r := range runes {
		_, err = dw.w.WriteRune(r)
	}
	if err != nil {
		return err
	}
	return nil
}

// WriteString writes a utf string.
func (dw *DataWriter) WriteUTF(value string) error {
	return dw.WriteString(value)
}

// WriteBoolArray writes a boolean array.
func (dw *DataWriter) WriteBoolArray(array []bool) error {
	arrayLength := len(array)
	err := dw.WriteVarInt(arrayLength)
	if err != nil {
		return err
	}

	var buffer byte = 0
	var bufferSize uint = 0
	for _, b := range array {
		if b {
			buffer |= 0x01 << bufferSize
		}
		bufferSize++
		if bufferSize == 8 {
			err = dw.w.WriteByte(buffer)
			if err != nil {
				return err
			}
			buffer = 0
			bufferSize = 0
		}
	}
	if bufferSize > 0 {
		err = dw.w.WriteByte(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dw *DataWriter) WriteInt16(value int16) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteUint16(value uint16) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteChar(value uint16) error {
	return dw.WriteUint16(value)
}
func (dw *DataWriter) WriteInt64(value int64) error {
	return binary.Write(dw.w, binary.BigEndian, value)
}
func (dw *DataWriter) WriteLong(value int64) error {
	return dw.WriteInt64(value)
}
func (dw *DataWriter) WriteBool(value bool) error {
	if value {
		return dw.WriteByte(1)
	} else {
		return dw.WriteByte(0)
	}
}
func (dw *DataWriter) Flush() error {
	return dw.w.Flush()
}
