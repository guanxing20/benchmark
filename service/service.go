package service

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"

	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
)

var ErrAlreadyStopped = errors.New("already stopped")

type Service interface {
	cliapp.Lifecycle
	Kill() error
}

type service struct {
	config  Config
	version string
	log     log.Logger

	stopped atomic.Bool
}

func NewService(version string, cfg Config, log log.Logger) Service {
	return &service{
		config:  cfg,
		version: version,
		log:     log,
	}
}

func readBenchmarkConfig(path string) ([]BenchmarkConfig, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	var config []BenchmarkConfig
	err = yaml.NewDecoder(file).Decode(&config)
	return config, err
}

func (s *service) Start(ctx context.Context) error {
	s.log.Info("Starting")

	config, err := readBenchmarkConfig(s.config.ConfigPath())
	if err != nil {
		return errors.Wrap(err, "failed to read benchmark config")
	}

	fmt.Printf("Started %+#v\n", config)

	return nil
}

// Stopped returns if the service as a whole is stopped.
func (s *service) Stopped() bool {
	return s.stopped.Load()
}

// Kill is a convenience method to forcefully, non-gracefully, stop the Service.
func (s *service) Kill() error {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return s.Stop(ctx)
}

// Stop fully stops the batch-submitter and all its resources gracefully. After stopping, it cannot be restarted.
// See driver.StopBatchSubmitting to temporarily stop the batch submitter.
// If the provided ctx is cancelled, the stopping is forced, i.e. the batching work is killed non-gracefully.
func (s *service) Stop(ctx context.Context) error {
	if s.stopped.Load() {
		return ErrAlreadyStopped
	}
	s.log.Info("Service stopping")

	// var result error

	// if result == nil {
	// 	s.stopped.Store(true)
	// 	s.log.Info("Service stopped")
	// }
	return nil
}
