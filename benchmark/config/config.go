package config

import (
	"github.com/base/base-bench/service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/urfave/cli/v2"
)

type CLIConfig struct {
	service.Config
}

func (c *CLIConfig) Check() error {
	return nil
}

func (c *CLIConfig) LogConfig() oplog.CLIConfig {
	return c.Config.LogConfig()
}

// NewCLIConfig parses the Config from the provided flags or environment variables.
func NewCLIConfig(ctx *cli.Context) *CLIConfig {
	return &CLIConfig{
		Config: service.NewConfig(ctx),
	}
}
