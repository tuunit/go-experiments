package main

import (
	"io/ioutil"
	"log"

	"github.com/hashicorp/go-plugin"
	"tuunit.com/goexp/plugins/shared"
)

type KV struct{}

func (KV) Put(key string, value []byte) error {
	return ioutil.WriteFile("kv_"+key, value, 0644)
}

func (KV) Get(key string) ([]byte, error) {
	return ioutil.ReadFile("kv_" + key)
}

func main() {
	log.SetOutput(ioutil.Discard)

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"kv": &shared.KVPlugin{Impl: &KV{}},
		},
	})
}
