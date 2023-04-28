// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-framework/packageclients/pkg/packageclient"
	"github.com/vmware-tanzu/tanzu-framework/packageclients/pkg/packagedatamodel"
	"tanzu-package-plugin-poc/pkg/package/flags"
)

var packageInstallOp = packagedatamodel.NewPackageOptions()

var PackageInstallCmd = &cobra.Command{
	Use:   "install INSTALLED_PACKAGE_NAME --package-name PACKAGE_NAME --version VERSION",
	Short: "Install a package",
	Args:  cobra.ExactArgs(1),
	Example: `
    # Install package contour with installed package name as 'contour-pkg' with specified version and without waiting for package reconciliation to complete 	
    tanzu package install contour-pkg --package-name contour.tanzu.vmware.com --namespace test-ns --version 1.15.1-tkg.1-vmware1 --wait=false
	
    # Install package contour with kubeconfig flag and waiting for package reconciliation to complete	
    tanzu package install contour-pkg --package-name contour.tanzu.vmware.com --namespace test-ns --version 1.15.1-tkg.1-vmware1 --kubeconfig path/to/kubeconfig`,
	RunE:              packageInstall,
	SilenceUsage:      true,
	PersistentPreRunE: packagingAvailabilityCheck,
}

func init() {
	PackageInstallCmd.Flags().StringVarP(&packageInstallOp.PackageName, "package-name", "p", "", "Name of the package to be installed")
	PackageInstallCmd.Flags().StringVarP(&packageInstallOp.Version, "version", "v", "", "Version of the package to be installed")
	PackageInstallCmd.Flags().BoolVarP(&packageInstallOp.CreateNamespace, "create-namespace", "", false, "Create namespace if the target namespace does not exist, optional")
	PackageInstallCmd.Flags().StringVarP(&packageInstallOp.Namespace, "namespace", "n", "default", "Namespace indicates the location of the repository from which the package is retrieved")
	PackageInstallCmd.Flags().StringVarP(&packageInstallOp.ServiceAccountName, "service-account-name", "", "", "Name of an existing service account used to install underlying package contents, optional")
	PackageInstallCmd.Flags().StringVarP(&packageInstallOp.ValuesFile, "values-file", "f", "", "The path to the configuration values file, optional")
	PackageInstallCmd.Flags().BoolVarP(&packageInstallOp.Wait, "wait", "", true, "Wait for the package reconciliation to complete, optional. To disable wait, specify --wait=false")
	PackageInstallCmd.Flags().DurationVarP(&packageInstallOp.PollInterval, "poll-interval", "", packagedatamodel.DefaultPollInterval, "Time interval between subsequent polls of package reconciliation status, optional")
	PackageInstallCmd.Flags().DurationVarP(&packageInstallOp.PollTimeout, "poll-timeout", "", packagedatamodel.DefaultPollTimeout, "Timeout value for polls of package reconciliation status, optional")
	PackageInstallCmd.MarkFlagRequired("package-name") //nolint
	PackageInstallCmd.MarkFlagRequired("version")      //nolint
}

func packageInstall(cmd *cobra.Command, args []string) error {
	packageInstallOp.PkgInstallName = args[0]

	pkgClient, err := packageclient.NewPackageClient(flags.PersistentFlagsDefault.Kubeconfig)
	if err != nil {
		return err
	}

	return pkgClient.InstallPackageSync(packageInstallOp, packagedatamodel.OperationTypeInstall)
}
