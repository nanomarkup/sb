package app

import (
	"fmt"

	"github.com/hashicorp/go-plugin"
)

type builder interface {
	Build(app string, sources *map[string]map[string]string) error
	Clean(app string, sources *map[string]map[string]string) error
	Generate(app string, sources *map[string]map[string]string) error
}

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

func handleError() {
	if r := recover(); r != nil {
		fmt.Printf(ErrorMessageF, r)
	}
}
