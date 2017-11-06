package gotensor_test

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/helinwang/gotensor"
	"github.com/stretchr/testify/require"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func TestGob(t *testing.T) {
	testCase := []interface{}{
		float32(2), float64(1),
		int8(3), int16(3), int32(4), int64(5),
		uint8(3), uint16(3),
		complex(100, 8),
		"string",

		// unsupported:
		// uint32, uint64
	}

	for _, v := range testCase {
		tensor, err := tf.NewTensor(v)
		require.Nil(t, err)

		t0 := gotensor.Tensor{tensor}
		require.Equal(t, v, t0.Value())

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
		require.Equal(t, v, t1.Value())

		var t2 gotensor.Tensor
		err = dec.Decode(&t2)
		require.Nil(t, err)
		require.Equal(t, v, t2.Value())
	}
}
