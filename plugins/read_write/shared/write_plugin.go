package shared

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type Writer interface {
	Update(data []byte) string
}

func (s *WriterRPCServer) Write(mapArgs map[string]interface{}, resp *string) error {
	data := mapArgs["data"].([]byte)
	*resp = s.Impl.Update(data)
	return nil
}

func (g *WriterRPCClient) Update(data []byte) string {
	var resp string
	err := g.client.Call("Plugin.Update", map[string]interface{}{
		"data":   data,
	}, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

type WriterRPCClient struct {
	client *rpc.Client
}

type WriterRPCServer struct {
	Impl Writer
}

type WriterPlugin struct {
	Impl Writer
}

func (r *WriterPlugin) Server(broker *plugin.MuxBroker) (interface{}, error){
	return &WriterRPCServer{Impl: r.Impl}, nil
}


func (r *WriterPlugin) Client (broker *plugin.MuxBroker, c *rpc.Client) (interface{}, error){
	return &WriterRPCClient{client: c}, nil
}