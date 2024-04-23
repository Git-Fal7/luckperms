package friendforgate

import (
	"github.com/git-fal7/luckperms/internal/plugin"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

var Plugin = proxy.Plugin{
	Name: "Luckperms",
	Init: plugin.InitPlugin,
}
