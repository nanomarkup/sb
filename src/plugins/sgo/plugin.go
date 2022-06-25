package sgo

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/sapplications/sbuilder/src/plugins"
)

func (p *Plugin) Execute() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})
	builder := builder{
		logger:    logger,
		builder:   p.Builder,
		generator: p.Generator,
	}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"sgo": &plugins.BuilderPlugin{Impl: &builder},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
