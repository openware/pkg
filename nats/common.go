// Package nats is used for handling nats (or nats jetstream) pub/sub
package nats

import (
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

// InitNats initialize nats using connectionSting
func InitNats(connectionString string) (*nats.Conn, error) {
	nc, err := nats.Connect(connectionString)

	return nc, err
}

// InitEmbededNats initialize nats in memory
func InitEmbededNats() (*nats.Conn, error) {
	opts := &server.Options{}
	ns, err := server.NewServer(opts)
	if err != nil {
		panic("failed to initialize nats mock server")
	}

	ns.Start()
	nc, err := nats.Connect(ns.ClientURL())

	return nc, err
}
