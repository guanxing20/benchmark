package service

import (
	"errors"
	"fmt"
	"os"

	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/urfave/cli/v2"
)

type Config interface {
	Check() error
	LogConfig() oplog.CLIConfig
	ConfigPath() string
}

type config struct {
	logConfig  oplog.CLIConfig
	configPath string
}

func NewConfig(ctx *cli.Context) Config {
	return &config{
		logConfig:  oplog.ReadCLIConfig(ctx),
		configPath: ctx.String("config"),
	}
}

func (c *config) ConfigPath() string {
	return c.configPath
}

func (c *config) Check() error {
	if c.configPath == "" {
		return errors.New("config path is required")
	}

	// ensure file exists
	if _, err := os.Stat(c.configPath); err != nil {
		return fmt.Errorf("config file does not exist: %w", err)
	}
	return nil
}

func (c *config) LogConfig() oplog.CLIConfig {
	return c.logConfig
}
