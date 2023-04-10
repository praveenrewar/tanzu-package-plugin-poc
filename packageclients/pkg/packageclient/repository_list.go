// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package packageclient

import (
	"github.com/pkg/errors"

	"tanzu-package-plugin-poc/packageclients/pkg/packagedatamodel"

	kappipkg "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packaging/v1alpha1"
)

func (p *pkgClient) ListRepositories(o *packagedatamodel.RepositoryOptions) (*kappipkg.PackageRepositoryList, error) {
	packageRepositoryList, err := p.kappClient.ListPackageRepositories(o.Namespace)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list existing package repositories in the cluster")
	}

	return packageRepositoryList, nil
}
