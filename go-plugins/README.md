# Simple RPC go-plugin implementation
Inspired by https://github.com/hashicorp/go-plugin/tree/main/examples/grpc

But instead of offering grpc and net/rpc implementations this directory only contains a single net/rpc implementation.

```bash
# Build the main application
$ go build -o kv

# Build the plugin
$ go build -o kv-plugin ./plugin

# Export the path to your build plugin
$ export KV_PLUGIN="$(pwd)/kv-plugin"

# Write
$ ./kv put hello world
# Or without building the main application
$ go run put hello world

# Read
$ ./kv get hello
world

# Or without building the main application
$ go run get hello
world
