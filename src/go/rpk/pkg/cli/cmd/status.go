// Copyright 2020 Vectorized, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/vectorizedio/redpanda/src/go/rpk/pkg/cli/cmd/common"
	"github.com/vectorizedio/redpanda/src/go/rpk/pkg/cli/cmd/debug"
	"github.com/vectorizedio/redpanda/src/go/rpk/pkg/config"
)

func NewStatusCommand(fs afero.Fs, mgr config.Manager) *cobra.Command {
	command := debug.NewInfoCommand(fs, mgr)
	command.Deprecated = common.DeprecationMessage("rpk debug info")
	command.Hidden = true
	return command
}
