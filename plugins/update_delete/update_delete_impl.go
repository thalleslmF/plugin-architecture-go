package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"io/ioutil"
	"log"
	"os"
	"plugin-hashcorp/newApp/plugins/update_delete/shared"
	"strings"
)
type UpdaterFile struct {
	logger hclog.Logger
}

type DeleterFile struct {
	logger hclog.Logger
}
func (*UpdaterFile) Update(key, value string) string {
	arrayUpdate := make([]string, 0)
	file, err := os.Open("/tmp/file.txt")
	if err != nil {
		f, err := os.Create("/tmp/file.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		return ""
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		arrayData := strings.Split(data, " ")
		if arrayData[0] == key {
			arrayData[1] = value
		}
		arrayUpdate = append(arrayUpdate, fmt.Sprintf("%s %s", arrayData[0],arrayData[1]))
	}
	fileData := strings.Join(arrayUpdate, "\n")

	err = ioutil.WriteFile("/tmp/file.txt", bytes.NewBufferString(fileData).Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return ""
}

func ( w *DeleterFile) Delete(key string) string {
	arrayUpdate := make([]string, 0)
	file, err := os.OpenFile("/tmp/file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	w.logger.Debug("file", err)
	if err != nil {
		_, err := os.Create("/tmp/file.txt")
		w.logger.Debug("file", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		arrayData := strings.Split(data, " ")
		if arrayData[0] == key {
			continue
		}
		arrayUpdate = append(arrayUpdate, fmt.Sprintf("%s %s", arrayData[0],arrayData[1]))
	}
	fileData := strings.Join(arrayUpdate, "\n")

	err = ioutil.WriteFile("/tmp/file.txt", bytes.NewBufferString(fileData).Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return ""

}
func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	deleter := &DeleterFile{
		logger: logger,
	}

	updater := &UpdaterFile{
		logger: logger,
	}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"reader": &shared.DeleterPlugin{Impl: deleter},
		"writer": &shared.UpdaterPlugin{Impl: updater},
	}

	logger.Debug("message from plugin", "foo", "bar")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}