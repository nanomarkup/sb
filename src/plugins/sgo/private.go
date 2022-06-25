package sgo

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type builder struct {
	logger    hclog.Logger
	builder   Builder
	generator Generator
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
