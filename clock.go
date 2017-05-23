package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

type Clock struct {
	Bps   float64
	Ticks int64
	Beats int64
}

func (c *Clock) Encode(b io.Writer) {
	buf := new(bytes.Buffer)
	for _, v := range []interface{}{c.Bps, c.Ticks, c.Beats} {
		err := binary.Write(buf, binary.BigEndian, v)
		if err != nil {
			log.Fatal("binary.Write failed:", err)
		}
	}
	b.Write(buf.Bytes())
}

func (c *Clock) Decode(b []byte) {
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.BigEndian, c)
	if err != nil {
		log.Fatal("binary.Read failed:", err)
	}
}
