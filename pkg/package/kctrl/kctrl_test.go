// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package kctrl

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
)

func TestKctrlInvoke(t *testing.T) {
	t.Run("Invoke", func(t *testing.T) {
		t.Run("add package commands from kctrl", func(t *testing.T) {
			testPlugin, err := plugin.NewPlugin(&plugin.PluginDescriptor{
				Name:        "test",
				Description: "test",
				Version:     "v1.0.0",
				Group:       "test",
			})
			assert.NoError(t, err)
			assert.Equal(t, false, testPlugin.Cmd.HasAvailableSubCommands())

			Invoke(testPlugin)

			assert.Equal(t, true, testPlugin.Cmd.HasAvailableSubCommands())
			assert.Contains(t, testPlugin.Cmd.UsageString(), "available     Manage available packages")
			assert.Contains(t, testPlugin.Cmd.UsageString(), "installed     Manage installed packages")
			assert.Contains(t, testPlugin.Cmd.UsageString(), "repository    Manage package repositories")
		})
	})
}
