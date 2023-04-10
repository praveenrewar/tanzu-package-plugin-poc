// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"github.com/spf13/cobra"

	"tanzu-package-plugin-poc/packageclients/pkg/packagedatamodel"
)

var repoOp = packagedatamodel.NewRepositoryOptions()

var repositoryCmd = &cobra.Command{
	Use:               "repository",
	Short:             "Repository operations",
	ValidArgs:         []string{"add", "list", "get", "delete", "update"},
	Args:              cobra.RangeArgs(1, 3),
	Long:              `Add, list, get or delete a package repository for Tanzu packages. A package repository is a collection of packages that are grouped together into an imgpkg bundle.`,
	PersistentPreRunE: packagingAvailabilityCheck,
}

func init() {
	repositoryCmd.PersistentFlags().StringVarP(&repoOp.Namespace, "namespace", "n", "default", "Namespace for repository command, optional")
}