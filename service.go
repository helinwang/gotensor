package gotensor

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

type Service struct {
	Sess  *tf.Session
	Graph *tf.Graph
}

type Edge struct {
	OpName string
	Index  int
}

type Feed struct {
	Edge   Edge
	Tensor Tensor
}

type Request struct {
	Feeds   []Feed
	Fetches []Edge
	Targets []string
}

type Response struct {
	Error   string
	Outputs []Tensor
}

func New(graph []byte) (*Service, error) {
	g := tf.NewGraph()
	err := g.Import(graph, "")
	if err != nil {
		return nil, err
	}

	s, err := tf.NewSession(g, nil)
	if err != nil {
		return nil, err
	}

	return &Service{Sess: s, Graph: g}, nil
}

func (s *Service) Run(req Request, resp *Response) error {
	feeds := make(map[tf.Output]*tf.Tensor)
	for _, f := range req.Feeds {
		feeds[s.Graph.Operation(f.Edge.OpName).Output(f.Edge.Index)] = f.Tensor.Tensor
	}

	fetches := make([]tf.Output, len(req.Fetches))
	for i, f := range req.Fetches {
		fetches[i] = s.Graph.Operation(f.OpName).Output(f.Index)
	}

	targets := make([]*tf.Operation, len(req.Targets))
	for i, t := range req.Targets {
		targets[i] = s.Graph.Operation(t)
	}

	tensors, err := s.Sess.Run(feeds, fetches, targets)
	if err != nil {
		resp.Error = err.Error()
		return nil
	}

	if len(tensors) == 0 {
		return nil
	}

	resp.Outputs = make([]Tensor, len(tensors))
	for i, t := range tensors {
		resp.Outputs[i] = Tensor{t}
	}
	return nil
}
