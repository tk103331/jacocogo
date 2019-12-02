package data

import (
	"bufio"
	"encoding/binary"
	"io"
)

// DataReader is wrapper of bufio.Reader.
type DataReader struct {
	r *bufio.Reader
}

func NewReader(reader io.Reader) *DataReader {
	return &DataReader{bufio.NewReader(reader)}
}

// Read reads a byte.
func (dr *DataReader) ReadByte() (byte, error) {
	return dr.r.ReadByte()
}
func (dr *DataReader) ReadRune() (rune, error) {
	r, _, err := dr.r.ReadRune()
	return r, err
}

// Read reads a int16.
func (dr *DataReader) ReadInt16() (int16, error) {
	var value int16
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// Read reads a uint16.
func (dr *DataReader) ReadUint16() (uint16, error) {
	var value uint16
	err := binary.Read(dr.r, binary.BigEndian, &value)
	return value, err
}

// ReadChar reads a java char , it is a uint16.
func (dr *DataReader) ReadChar() (uint16, error) {
	return dr.ReadUint16()
}

// ReadUTF reads a utf string.
func (dr *DataReader) ReadUTF() (string, error) {
	return dr.ReadString()
}

// ReadInt64 reads a int64.
func (dr *DataReader) ReadInt64() (int64, error) {
	var value int64
	err := binary.Read(dr.r, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// ReadInt64 reads a java long, it is a int64.
func (dr *DataReader) ReadLong() (int64, error) {
	return dr.ReadInt64()
}

// ReadVarInt reads a variable length representation of an integer value.
func (dr *DataReader) ReadVarInt() (int, error) {
	b, e := dr.r.ReadByte()
	if e != nil {
		return 0, e
	}
	value := 0xFF & int(b)
	if value&0x80 == 0 {
		return value, nil
	}
	varInt, e := dr.ReadVarInt()
	if e != nil {
		return 0, e
	}
	return (value & 0x7F) | (varInt << 7), nil
}

// ReadBooleanArray reads a boolean array.
func (dr *DataReader) ReadBoolArray() ([]bool, error) {
	count, err := dr.ReadVarInt()
	if err != nil {
		return nil, err
	}

	probes := make([]bool, count)
	var buffer byte = 0x00
	for i := 0; i < len(probes); i++ {
		if (i % 8) == 0 {
			buffer, err = dr.r.ReadByte()
			if err != nil {
				return nil, err
			}
		}
		probes[i] = (buffer & 0x01) != 0
		buffer = buffer >> 1
	}

	return probes, nil
}

// ReadBooleanArray reads a string.
func (dr *DataReader) ReadString() (string, error) {
	var bytesNumber uint16
	err := binary.Read(dr.r, binary.BigEndian, &bytesNumber)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, bytesNumber)
	_, err = dr.r.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:]), nil
}
