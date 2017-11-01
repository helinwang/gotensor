package gotensor

import (
	"bytes"
	"encoding/gob"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

// Tensor is a TensorFlow tensor that satisfies the GobDecoder and the
// GobEncode interface.
type Tensor struct {
	T *tf.Tensor
}

// GobDecode overwrites the receiver, which must be a pointer,
// with the value represented by the byte slice, which was written
// by GobEncode, usually for the same concrete type.
func (t *Tensor) GobDecode(b []byte) error {
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)

	var dt tf.DataType
	err := dec.Decode(&dt)
	if err != nil {
		return err
	}

	var shape []int64
	err = dec.Decode(&shape)
	if err != nil {
		return err
	}

	tensor, err := tf.ReadTensor(dt, shape, r)
	if err != nil {
		return err
	}

	t.T = tensor
	return nil
}

// GobEncode returns a byte slice representing the encoding of the
// receiver for transmission to a GobDecoder, usually of the same
// concrete type.
func (t Tensor) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(t.T.DataType())
	if err != nil {
		return nil, err
	}

	err = enc.Encode(t.T.Shape())
	if err != nil {
		return nil, err
	}

	_, err = t.T.WriteContentsTo(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
