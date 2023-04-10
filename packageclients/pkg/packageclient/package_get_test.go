// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package packageclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"tanzu-package-plugin-poc/packageclients/pkg/fakes"
	. "tanzu-package-plugin-poc/packageclients/pkg/packageclient"
	"tanzu-package-plugin-poc/packageclients/pkg/packagedatamodel"

	kappipkg "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packaging/v1alpha1"
)

var _ = Describe("Get Installed Package", func() {
	var (
		ctl     PackageClient
		kappCtl *fakes.KappClient
		err     error
		opts    = packagedatamodel.PackageOptions{
			PackageName: testPkgName,
			Namespace:   testNamespaceName,
		}
		options    = opts
		pkgInstall *kappipkg.PackageInstall
	)

	JustBeforeEach(func() {
		ctl, err = NewPackageClientWithKappClient(kappCtl)
		Expect(err).NotTo(HaveOccurred())

		pkgInstall, err = ctl.GetPackageInstall(&options)
	})

	Context("failure in getting installed packages due to GetPackageInstall API error", func() {
		BeforeEach(func() {
			kappCtl = &fakes.KappClient{}
			kappCtl.GetPackageInstallReturns(nil, errors.New("failure in GetPackageInstall"))
		})
		It(testFailureMsg, func() {
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failure in GetPackageInstall"))
			Expect(pkgInstall).To(BeNil())
		})
		AfterEach(func() { options = opts })
	})

	Context("success in getting installed packages", func() {
		BeforeEach(func() {
			kappCtl = &fakes.KappClient{}
			kappCtl.GetPackageInstallReturns(testPkgInstall, nil)
		})
		It(testSuccessMsg, func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(pkgInstall).NotTo(BeNil())
			Expect(pkgInstall).To(Equal(testPkgInstall))
		})
		AfterEach(func() { options = opts })
	})
})
