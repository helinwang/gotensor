package gotensor_test

import (
	"io/ioutil"
	"testing"

	"github.com/helinwang/gotensor"
	"github.com/stretchr/testify/require"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func TestFetch(t *testing.T) {
	// a graph of b = a + 1
	graph, err := ioutil.ReadFile("./test_data/a_plus_1.pb")
	require.Nil(t, err)

	s, err := gotensor.New(graph)
	require.Nil(t, err)

	var resp gotensor.Response
	tensor, err := tf.NewTensor(int32(2))
	require.Nil(t, err)

	err = s.Run(gotensor.Request{
		Feeds: []gotensor.Feed{
			gotensor.Feed{
				Edge:   gotensor.Edge{OpName: "a"},
				Tensor: gotensor.Tensor{tensor},
			},
		},
		Fetches: []gotensor.Edge{
			gotensor.Edge{OpName: "b"},
		},
	}, &resp)
	require.Nil(t, err)
	require.Equal(t, "", resp.Error)
	require.Equal(t, int32(3), resp.Outputs[0].Value())
}

func TestTarget(t *testing.T) {
	// a graph of b = a + 1
	graph, err := ioutil.ReadFile("./test_data/a_plus_1.pb")
	require.Nil(t, err)

	s, err := gotensor.New(graph)
	require.Nil(t, err)

	var resp gotensor.Response
	tensor, err := tf.NewTensor(int32(2))
	require.Nil(t, err)

	err = s.Run(gotensor.Request{
		Feeds: []gotensor.Feed{
			gotensor.Feed{
				Edge:   gotensor.Edge{OpName: "a"},
				Tensor: gotensor.Tensor{tensor},
			},
		},
		Targets: []string{"b"},
	}, &resp)
	require.Nil(t, err)
	require.Equal(t, "", resp.Error)
	// no fetch is required, so no output.
	require.Equal(t, 0, len(resp.Outputs))
}
