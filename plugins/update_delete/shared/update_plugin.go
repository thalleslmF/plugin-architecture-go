package shared

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type Updater interface {
	Update(key, value string) string
}

func (s *UpdaterRPCServer) Update(mapArgs map[string]interface{}, resp *string) error {
	key := mapArgs["key"].(string)
	value := mapArgs["value"].(string)
	*resp = s.Impl.Update(key, value)
	return nil
}

func (g *UpdaterRPCClient) Update(data  map[string]interface{}) string {
	var resp string
	err := g.client.Call("Plugin.Update", data, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

type UpdaterRPCClient struct {
	client *rpc.Client
}

type UpdaterRPCServer struct {
	Impl Updater
}

type UpdaterPlugin struct {
	Impl Updater
}

func (r *UpdaterPlugin) Server(broker *plugin.MuxBroker) (interface{}, error){
	return &UpdaterRPCServer{Impl: r.Impl}, nil
}


func (r *UpdaterPlugin) Client (broker *plugin.MuxBroker, c *rpc.Client) (interface{}, error){
	return &UpdaterRPCClient{client: c}, nil
}