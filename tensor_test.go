package gotensor_test

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/helinwang/gotensor"
	"github.com/stretchr/testify/require"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func TestTensorGob(t *testing.T) {
	v := float64(1)
	tensor, err := tf.NewTensor(v)
	require.Nil(t, err)

	t0 := gotensor.Tensor{T: tensor}
	require.Equal(t, v, t0.T.Value())

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(t0)
	require.Nil(t, err)
	err = enc.Encode(&t0)
	require.Nil(t, err)

	var t1 gotensor.Tensor
	dec := gob.NewDecoder(bytes.NewReader(buf.Bytes()))
	err = dec.Decode(&t1)
	require.Nil(t, err)
	require.Equal(t, v, t1.T.Value())

	t1.T = nil
	err = dec.Decode(&t1)
	require.Nil(t, err)
	require.Equal(t, v, t1.T.Value())
}
