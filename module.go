package fxconfig

import (
	"github.com/templatedop/config"
	"go.uber.org/fx"
)


const ModuleName = "config"

var FxConfigModule = fx.Module(
	ModuleName,
	fx.Provide(
		config.NewDefaultConfigFactory,
		NewFxConfig,
	),
)

// FxConfigParam allows injection of the required dependencies in [NewFxConfig].
type FxConfigParam struct {
	fx.In
	Factory config.ConfigFactory
}

// NewFxConfig returns a [config.Config].
func NewFxConfig(p FxConfigParam) (*config.Config, error) {
	return p.Factory.Create(
		config.WithFileName("config"),
		config.WithFilePaths(
			".",
			"./configs",
		),
	)
}
