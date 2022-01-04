package codec

import (
	"bufio"
	"encoding/json"
	"io"
)

type JsonCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *json.Decoder
	enc  *json.Encoder
}

var _ Codec = (*JsonCodec)(nil)

func NewJsonCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &JsonCodec{
		conn: conn,
		buf:  buf,
		dec:  json.NewDecoder(conn),
		enc:  json.NewEncoder(conn),
	}
}

func (c *JsonCodec) ReadHeader(h *Header) error {
	return nil
}

func (c *JsonCodec) ReadBody(body interface{}) error {
	return nil
}

func (c *JsonCodec) Write(h *Header, body interface{}) error {
	return nil
}

func (c *JsonCodec) Close() error {
	return c.conn.Close()
}
