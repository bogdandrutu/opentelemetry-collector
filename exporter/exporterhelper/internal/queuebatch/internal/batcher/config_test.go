// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batcher

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfig_Validate(t *testing.T) {
	cfg := newTestConfig()
	require.NoError(t, cfg.Validate())

	cfg = newTestConfig()
	cfg.FlushTimeout = 0
	require.EqualError(t, cfg.Validate(), "`flush_timeout` must be positive")

	cfg = newTestConfig()
	cfg.MinSize = -1
	require.EqualError(t, cfg.Validate(), "`min_size` must be non-negative")

	cfg = newTestConfig()
	cfg.MaxSize = -1
	require.EqualError(t, cfg.Validate(), "`max_size` must be non-negative")

	cfg = newTestConfig()
	cfg.MinSize = 2048
	cfg.MaxSize = 1024
	require.EqualError(t, cfg.Validate(), "`max_size` must be greater or equal to `min_size`")
}

func newTestConfig() Config {
	return Config{
		FlushTimeout: 200 * time.Millisecond,
		MinSize:      2048,
		MaxSize:      0,
	}
}
