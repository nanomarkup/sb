package sgo

import (
	"github.com/hashicorp/go-plugin"
	"github.com/sapplications/sbuilder/src/plugins"
)

func (p *Plugin) Execute() {
	builder := builder{
		builder:   p.Builder,
		generator: p.Generator,
	}
	builder.builder.SetLogger(p.Logger)
	builder.generator.SetLogger(p.Logger)
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		AppName: &plugins.BuilderPlugin{Impl: &builder},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: p.Handshake,
		Plugins:         pluginMap,
	})
}
