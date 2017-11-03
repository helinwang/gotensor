package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/helinwang/gotensor"
)

func main() {
	graph := flag.String("graph", "", "path to the graph file")
	port := flag.Int("port", 8080, "serve port")
	ip := flag.String("ip", "127.0.0.1", "serve ip")
	flag.Parse()

	b, err := ioutil.ReadFile(*graph)
	if err != nil {
		panic(err)
	}

	s, err := gotensor.New(b)
	if err != nil {
		panic(err)
	}

	rpc.Register(s)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", *ip+":"+strconv.Itoa(*port))
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
