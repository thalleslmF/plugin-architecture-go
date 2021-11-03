package plugin

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"os"
)

func handShakeConfig(key, value string) plugin.HandshakeConfig {
return plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   key,
	MagicCookieValue: value,
}
}

func getLogger() hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})
}
