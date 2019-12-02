package data

import "sync"

type Buffer struct {
	data []byte
	lock sync.Mutex
	cap  int
	size int
}

func NewBuffer(cap int) *Buffer {
	return &Buffer{size: 0, cap: cap}
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
