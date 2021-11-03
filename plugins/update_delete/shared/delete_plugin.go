package shared

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type Deleter interface {
	Delete(key string) string
}

func (s *DeleterRPCServer) Delete(mapArgs map[string]interface{}, resp *string) error {
	args := mapArgs["data"].(string)
	*resp = s.Impl.Delete(args)
	return nil
}

func (g *DeleterRPCClient) Delete(key string) string {
	var resp string
	err := g.client.Call("Plugin.Delete", map[string]interface{}{
		"data":   key,
	}, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

type DeleterRPCClient struct {
	client *rpc.Client
}

type DeleterRPCServer struct {
	Impl Deleter
}

type DeleterPlugin struct {
	Impl Deleter
}

func (r *DeleterPlugin) Server(broker *plugin.MuxBroker) (interface{}, error){
	return &DeleterRPCServer{Impl: r.Impl}, nil
}

func (r *DeleterPlugin) Client (broker *plugin.MuxBroker, c *rpc.Client) (interface{}, error){
	return &DeleterRPCClient{client: c}, nil
}