package app

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/sapplications/sbuilder/src/plugins"
)

var langs = struct {
	Go string
}{
	"go",
}

var suppLangs = map[string]bool{
	langs.Go: true,
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SMART_PLUGIN",
	MagicCookieValue: "sbuilder",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"sgo": &plugins.BuilderPlugin{},
}

func newPlugin(name string) (client *plugin.Client, raw interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(ErrorMessageF, r)
			if client != nil {
				client.Kill()
			}
		}
	}()
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Error,
	})
	client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(fmt.Sprintf("%s.exe", name)),
		Logger:          logger,
	})
	rpcClient, err := client.Client()
	if err != nil {
		return nil, nil, err
	}
	raw, err = rpcClient.Dispense(name)
	if err != nil {
		return nil, nil, err
	}
	return
}

func handleError() {
	if r := recover(); r != nil {
		fmt.Printf(ErrorMessageF, r)
	}
}
