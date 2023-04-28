// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-framework/packageclients/pkg/packageclient"
	"tanzu-package-plugin-poc/pkg/package/flags"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/component"
)

var repositoryGetCmd = &cobra.Command{
	Use:   "get REPOSITORY_NAME",
	Short: "Get details for a package repository",
	Args:  cobra.ExactArgs(1),
	Example: `
    # Get details for a repository in specified namespace 	
    tanzu package repository get repo --namespace test-ns`,
	RunE:         repositoryGet,
	SilenceUsage: true,
}

func init() {
	repositoryGetCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format (yaml|json|table), optional")
	PackageRepositoryCmd.AddCommand(repositoryGetCmd)
}

func repositoryGet(cmd *cobra.Command, args []string) error {
	if len(args) == 1 {
		repoOp.RepositoryName = args[0]
	} else {
		return errors.New("incorrect number of input parameters. Usage: tanzu package repository get REPOSITORY_NAME [FLAGS]")
	}

	pkgClient, err := packageclient.NewPackageClient(flags.PersistentFlagsDefault.Kubeconfig)
	if err != nil {
		return err
	}
	t, err := component.NewOutputWriterWithSpinner(cmd.OutOrStdout(), getOutputFormat(),
		fmt.Sprintf("Retrieving repository %s...", repoOp.RepositoryName), true)
	if err != nil {
		return err
	}

	packageRepository, err := pkgClient.GetRepository(repoOp)
	if err != nil || packageRepository == nil {
		t.StopSpinner()
		return err
	}

	repository, tag, err := packageclient.GetCurrentRepositoryAndTagInUse(packageRepository)
	if err != nil {
		t.StopSpinner()
		return err
	}

	t.SetKeys("name", "version", "repository", "tag", "status", "reason")
	t.AddRow(packageRepository.Name, packageRepository.ResourceVersion, repository, tag,
		packageRepository.Status.FriendlyDescription, packageRepository.Status.UsefulErrorMessage)

	t.RenderWithSpinner()
	return nil
}
