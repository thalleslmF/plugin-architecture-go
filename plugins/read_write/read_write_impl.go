package main

import (
	"bufio"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"io/ioutil"
	"log"
	"os"
	shared2 "plugin-hashcorp/newApp/plugins/read_write/shared"
	"strings"
)
type ReaderFile struct {
	logger hclog.Logger
}

type WriterFile struct {
	logger hclog.Logger
}
func (*ReaderFile) Read(key string) string {
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
			return arrayData[1]
		}
	}
	return ""
}

func ( w *WriterFile) Update(data []byte) string {

	file, err := os.OpenFile("/tmp/file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	w.logger.Debug("file", err)
	if err != nil {
		_, err := os.Create("/tmp/file.txt")
		w.logger.Debug("file", err.Error())
		if err != nil {
			log.Fatalln(err)
		}
	}
	_, err = file.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	byte, err := ioutil.ReadFile("/tmp/file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return string(byte)
}
func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	reader := &ReaderFile{
		logger: logger,
	}

	writer := &WriterFile{
		logger: logger,
	}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"reader": &shared2.ReaderPlugin{Impl: reader},
		"writer": &shared2.WriterPlugin{Impl: writer},
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