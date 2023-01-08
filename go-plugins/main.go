package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"tuunit.com/goexp/plugins/shared"
)

func main() {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"kv": &shared.KVPlugin{},
		},
		Cmd: exec.Command(os.Getenv("KV_PLUGIN")),
		Logger: hclog.New(&hclog.LoggerOptions{
			Level:      hclog.Error,
			JSONFormat: true,
		}),
	})

	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense("kv")
	if err != nil {
		log.Fatal(err)
	}

	kv := raw.(shared.KV)

	os.Args = os.Args[1:]

	switch os.Args[0] {
	case "get":
		res, err := kv.Get(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(res))

	case "put":
		err := kv.Put(os.Args[1], []byte(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Please only use 'get' or 'put', given %q", os.Args[0])
	}
}
