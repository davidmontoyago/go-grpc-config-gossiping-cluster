# go-grpc-config-gossiping-cluster

Example of eventually consistent cluster of gRPC services with [hashicorp/memberlist](https://github.com/hashicorp/memberlist).

```
# generate code if needed
make grpc

# watch your TCP/UDP connections
watch 'netstat -an | grep "790\|900"'

# run a gossiping cluster of gRPC services
make run-cluster

# send a request to any node; request should be gossiped to the other nodes one by one
make run-client
```