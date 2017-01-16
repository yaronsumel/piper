package main

import (
	"encoding/binary"
	"github.com/yaronsumel/pipe"
	"github.com/yaronsumel/ttls"
	"log"
	"net"
	"os"
)

type server struct {
	laddr   string
	storage *storage
	ttls    *ttls.TTLS
	pipeEOF bool
	len     uint64
	verbose bool
}

var knockBytes = []byte{83, 79, 83}

func newServer(laddr string, verbose bool) *server {
	s, err := ttls.NewTTLSListener(laddr, nil)
	if err != nil {
		panic(1)
	}
	return &server{
		laddr:   laddr,
		storage: newStorage(),
		ttls:    s,
		pipeEOF: false,
		verbose: verbose,
	}
}

// listen is limited by design to server one client and die
func (s *server) listen() {
	go s.pipe()
	log.Printf("[piper::server] listening on %s", s.laddr)
	conn, err := s.ttls.Listener.Accept()
	if err != nil {
		if s.verbose {
			log.Printf("[piper::server] accept: %s", err)
		}
	}
	s.handle(conn)
}

func (s *server) handle(c net.Conn) {

	log.Printf("[piper::server] connection accepted from %s", c.RemoteAddr())
	defer c.Close()

	var headerFlag = false
	var wroteLen int
	var buf = make([]byte, bufsize)

	// cmp client len of received bytes to the len of the data
	go func() {
		for {
			_, err := c.Read(buf[:cap(buf)])
			size := binary.BigEndian.Uint64(buf)
			if err != nil || s.pipeEOF && size == s.len || size > s.len {
				if s.verbose {
					log.Printf("[piper::server] client has reached EOF")
				}
				os.Exit(0)
			}
		}
	}()

	for {
		data := s.storage.next()
		if !headerFlag {
			headerFlag = true
			c.Write(knockBytes)
			continue
		}
		c.Write(data)
		if s.verbose {
			wroteLen += len(data)
			log.Printf("[piper::server] client got %d bytes", wroteLen)
		}
	}
}

func (s *server) pipe() {
	StdinChannel := make(pipe.StdDataChannel)
	go pipe.AsyncRead(pipe.Stdin, bufsize, StdinChannel)
	for {
		select {
		case stdin := <-StdinChannel:
			if stdin.Err != nil {
				s.pipeEOF = true
				log.Printf("[piper::server] EOF reached (%d bytes)", s.len)
				return
			}
			s.storage.append(stdin.Data)
			s.len += uint64(len(stdin.Data))
			if s.verbose {
				log.Printf("[piper::server] got stdin data (%d bytes)", s.len)
			}
		}
	}
}
