package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"log"
	"net"
	"os"
)

type client struct {
	raddr           string
	dataLen         int
	startOutputFlag bool
	verboseFlag     bool
}

func newClient(raddr string, verbose bool) *client {
	c := &client{
		raddr: raddr,
	}
	c.startOutputFlag = false
	c.verboseFlag = verbose
	return c
}

func (c *client) dial() net.Conn {
	conn, err := tls.Dial("tcp", c.raddr, &tls.Config{
		InsecureSkipVerify:       true,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	})
	if err != nil {
		log.Fatalf("[piper::client] %s", err)
	}
	log.Printf("[piper::client] connected to %s (TLS)", conn.RemoteAddr())
	return conn
}

func (c *client) connect() {

	conn := c.dial()
	defer conn.Close()

	var transferSizeBuffer = make([]byte, bufsize)
	var buf = make([]byte, bufsize)

	for {

		n, err := conn.Read(buf[:cap(buf)])
		if err != nil {
			log.Fatalf("[piper::client] %s - got %d bytes", err, c.dataLen)
		}

		//start of piped data
		if bytes.Equal(buf[:n], knockBytes) {
			c.startOutputFlag = true
			continue
		}

		if c.startOutputFlag {
			os.Stdout.Write(buf[:n])
		}

		c.dataLen += n

		//write to server the byte count
		binary.BigEndian.PutUint64(transferSizeBuffer, uint64(c.dataLen))
		conn.Write(transferSizeBuffer)

		if c.verboseFlag {
			log.Printf("[piper::client] got %d bytes", c.dataLen)
		}
	}

}
