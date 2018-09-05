package main

import (
	"fmt"
	"html"
	"log"
	"net"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

type Memcached struct {
	Conn net.Conn
}

type Option func(*Memcached) error

func New(network string, host string, port string, options ...Option) (*Memcached, error) {
	if network != "tcp" && network != "" {
		return nil, errors.New("invaild network protocol")
	}

	conn, err := net.Dial(network, fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, errors.Wrap(err, "failed gomemcached.New():")
	}

	mc := &Memcached{
		Conn: conn,
	}

	return mc, nil
}

func main() {
	mc, err := New("tcp", "localhost", "11211")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(mc.Conn, "<!--\n"+html.EscapeString(spew.Sdump(mc.Conn))+"\n-->")
}
