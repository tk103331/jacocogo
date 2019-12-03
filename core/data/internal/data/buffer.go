package data

import (
	"sync"
)

type Buffer struct {
	data []byte
	lock sync.Mutex
	cap  int
	size int
	pos  int
}

type DataBuffer struct {
	*Buffer
	*DataWriter
	*DataReader
}

func NewBuffer(cap int) *Buffer {
	return &Buffer{size: 0, pos: 0, cap: cap}
}

func NewDataBuffer(cap int) *DataBuffer {
	buffer := &Buffer{size: 0, pos: 0, cap: cap}
	return &DataBuffer{Buffer: buffer, DataWriter: NewWriter(buffer), DataReader: NewReader(buffer)}
}

func (b *Buffer) Read(data []byte) (int, error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	l := len(data)
	i := 0
	for ; i < l && i < b.size-b.pos; i++ {
		data[i] = b.data[b.pos+i]
	}
	b.pos = b.pos + i
	return i, nil
}

func (b *Buffer) Write(data []byte) (int, error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	l1 := len(b.data)
	l2 := len(data)

	if l1+l2 > b.cap {
		l := b.cap - l1
		b.data = append(b.data, data[0:b.cap-l1]...)
		b.size = b.size + l
		return l, nil
	} else {
		b.data = append(b.data, data...)
		b.size = b.size + l2
		return l2, nil
	}
}

func (b *Buffer) Data() []byte {
	b.lock.Lock()
	defer b.lock.Unlock()
	data := make([]byte, b.size)
	for i := 0; i < b.size; i++ {
		data[i] = b.data[i]
	}
	return data
}
