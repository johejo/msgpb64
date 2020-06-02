// Package msgpb64 provides encoder/decoder that combines msgpack and base64 to serialize/deserialize any structure as a string.
package msgpb64

import (
	"bytes"
	"encoding/base64"
	"io"

	"github.com/vmihailenco/msgpack/v4"
)

// Decoder describes decoder for msgpack and base64.
type Decoder struct {
	dec *msgpack.Decoder
}

// NewDecoder returns a new Decoder.
func NewDecoder(enc *base64.Encoding, r io.Reader) *Decoder {
	return &Decoder{dec: msgpack.NewDecoder(base64.NewDecoder(enc, r))}
}

// Decode decodes a stream into a value that encoded with msgpack and base64.
func (d *Decoder) Decode(v interface{}) error {
	return d.dec.Decode(&v)
}

// Encoder describes encoder for msgpack and base64.
type Encoder struct {
	enc *base64.Encoding
	w   io.Writer
}

// NewEncoder returns a new Encoder.
func NewEncoder(enc *base64.Encoding, w io.Writer) *Encoder {
	return &Encoder{w: w, enc: enc}
}

// Encode encodes a stream with msgpack and base64 to a value.
func (e *Encoder) Encode(v interface{}) error {
	var buf bytes.Buffer
	if err := msgpack.NewEncoder(&buf).Encode(v); err != nil {
		return err
	}
	b64e := base64.NewEncoder(e.enc, e.w)
	if _, err := io.Copy(b64e, &buf); err != nil {
		return err
	}
	return b64e.Close()
}
