// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batcher // import "go.opentelemetry.io/collector/exporter/exporterhelper/internal/queuebatch"

import (
	"errors"
	"time"
)

// Config defines a configuration for batching requests based on a timeout and a minimum number of items.
type Config struct {
	// FlushTimeout sets the time after which a batch will be sent regardless of its size.
	FlushTimeout time.Duration `mapstructure:"flush_timeout"`

	// MinSize defines the configuration for the minimum size of a batch.
	MinSize int64 `mapstructure:"min_size"`

	// MaxSize defines the configuration for the maximum size of a batch.
	MaxSize int64 `mapstructure:"max_size"`
}

func (cfg *Config) Validate() error {
	if cfg == nil {
		return nil
	}

	if cfg.FlushTimeout <= 0 {
		return errors.New("`flush_timeout` must be positive")
	}

	if cfg.MinSize < 0 {
		return errors.New("`min_size` must be non-negative")
	}

	if cfg.MaxSize < 0 {
		return errors.New("`max_size` must be non-negative")
	}

	if cfg.MaxSize > 0 && cfg.MaxSize < cfg.MinSize {
		return errors.New("`max_size` must be greater or equal to `min_size`")
	}

	return nil
}
