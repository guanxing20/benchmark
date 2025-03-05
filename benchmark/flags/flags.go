package flags

import (
	"github.com/urfave/cli/v2"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
)

const EnvVarPrefix = "BASE_BENCH"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(EnvVarPrefix, name)
}

var (
	ConfigFlag = &cli.StringFlag{
		Name:     "config",
		Usage:    "Config Path",
		EnvVars:  prefixEnvVars("CONFIG"),
		Required: true,
	}
)

// Flags contains the list of configuration options available to the binary.
var Flags = []cli.Flag{
	ConfigFlag,
}

func init() {
	Flags = append(Flags, oplog.CLIFlags(EnvVarPrefix)...)
}
