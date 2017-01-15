package main

import (
	"bytes"
	"sync"
)

const bufsize = 1024

type storage struct {
	buffer *bytes.Buffer
	mtx    sync.Mutex
}

func newStorage() *storage {
	var buf = make([]byte, bufsize)
	return &storage{
		buffer:bytes.NewBuffer(buf),
		mtx:sync.Mutex{},
	}
}

func (s *storage)append(data []byte) (int, error) {
	s.mtx.Lock()
	n, err := s.buffer.Write(data)
	s.mtx.Unlock()
	return n, err
}

func (s *storage)next() []byte {
	s.mtx.Lock()
	data := s.buffer.Next(bufsize)
	s.mtx.Unlock()
	return data
}