package plugin

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/go-plugin"
	"log"
	"os/exec"
	"plugin-hashcorp/newApp/plugins/read_write/shared"
)

type PluginRunner struct {
	Name string
	Run  func(args []string)
	Args []string
	Short string
}

func (r *PluginRunner) SetRunArgs(args []string) {
	r.Args = args
}

var mapPluginsRun = map[string]*PluginRunner{
	"add": &PluginRunner{
		Name: "add",
		Short: "Insert a key and a value to a file",
		Run: func(args []string) {

			client := plugin.NewClient(&plugin.ClientConfig{

				HandshakeConfig: handShakeConfig("BASIC_PLUGIN", "hello"),
				Plugins: map[string]plugin.Plugin{
					"writer": &shared.WriterPlugin{},
				},
				Cmd:    exec.Command("./plugins/read_write/read_write"),
				Logger: getLogger(),
			})
			rpcClient, err := client.Client()
			if err != nil {
				log.Fatal(err)
			}
			raw, err := rpcClient.Dispense("writer")
			if err != nil {
				log.Fatal(err)
			}
			writer := raw.(shared.Writer)
			runArgs := fmt.Sprintf("%s %s \n", args[0], args[1])
			writer.Update(bytes.NewBufferString(runArgs).Bytes())
		},
	},
	//"update": PluginRun{},
	//"delete": PluginRun{},
	"get": &PluginRunner{
		Name: "get",
		Short: "Get a value from a key in a file",
		Run: func(args []string) {

			client := plugin.NewClient(&plugin.ClientConfig{

				HandshakeConfig: handShakeConfig("BASIC_PLUGIN", "hello"),
				Plugins: map[string]plugin.Plugin{
					"reader": &shared.ReaderPlugin{},
				},
				Cmd:    exec.Command("./plugins/read_write/read_write"),
				Logger: getLogger(),
			})
			rpcClient, err := client.Client()
			if err != nil {
				log.Fatal(err)
			}
			raw, err := rpcClient.Dispense("reader")
			if err != nil {
				log.Fatal(err)
			}
			reader := raw.(shared.Reader)
			value := reader.Read(args[0])
			fmt.Println(value)
		},
	},
}