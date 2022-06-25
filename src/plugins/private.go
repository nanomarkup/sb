package plugins

import "net/rpc"

type builderClient struct {
	client *rpc.Client
}

type builderServer struct {
	Impl Builder
}
