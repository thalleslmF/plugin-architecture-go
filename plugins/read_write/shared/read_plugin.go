package shared

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type Reader interface {
	Read(key string) string
}

func (s *ReaderRPCServer) Read(mapArgs map[string]interface{}, resp *string) error {
	args := mapArgs["data"].(string)
	*resp = s.Impl.Read(args)
	return nil
}

func (g *ReaderRPCClient) Read(key string) string {
	var resp string
	err := g.client.Call("Plugin.Read", map[string]interface{}{
		"data":   key,
	}, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

type ReaderRPCClient struct {
	client *rpc.Client
}

type ReaderRPCServer struct {
	Impl Reader
}

type ReaderPlugin struct {
	Impl Reader
}

func (r *ReaderPlugin) Server(broker *plugin.MuxBroker) (interface{}, error){
	return &ReaderRPCServer{Impl: r.Impl}, nil
}

func (r *ReaderPlugin) Client (broker *plugin.MuxBroker, c *rpc.Client) (interface{}, error){
	return &ReaderRPCClient{client: c}, nil
}