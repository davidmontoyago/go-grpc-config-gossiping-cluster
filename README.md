# go-grpc-config-gossiping-cluster

Example of eventually consistent cluster of gRPC services with [hashicorp/memberlist](https://github.com/hashicorp/memberlist). 

Uses Memberlist [gossiping layer](https://github.com/hashicorp/memberlist/blob/2288bf30e9c8d7b5f6549bf62e07120d72fd4b6c/delegate.go) to replicate config data across nodes. Data is sent over the wire as [gobs](https://blog.golang.org/gobs-of-data).

## See

[SWIM membership protocol](https://prakhar.me/articles/swim/)

## Run
```
# generate code if needed
make grpc

# watch your TCP/UDP connections
watch 'netstat -an | grep "790\|900"'

# run a gossiping cluster of gRPC services
make run-cluster

# send a put config request to any node; new config should be gossiped infection-style to the other nodes
make run-client
```