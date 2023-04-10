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
)

const (
	testSecretName = "test-secret"
	testPassword   = "test-password"
	testRegistry   = "test-registry"
	testUsername   = "test-username"
	testServer     = "us-east4-docker.pkg.dev"
)

var _ = Describe("Add Secret", func() {
	var (
		ctl     PackageClient
		crtCtl  *fakes.CrtClient
		kappCtl *fakes.KappClient
		err     error
		opts    = packagedatamodel.RegistrySecretOptions{
			ExportToAllNamespaces: false,
			Namespace:             testNamespaceName,
			Password:              testPassword,
			Server:                testRegistry,
			SecretName:            testSecretName,
			Username:              testUsername,
		}
		options = opts
	)

	JustBeforeEach(func() {
		ctl, err = NewPackageClientWithKappClient(kappCtl)
		Expect(err).NotTo(HaveOccurred())
		err = ctl.AddRegistrySecret(&options)
	})

	Context("failure in creating Secret due to Create API error", func() {
		BeforeEach(func() {
			kappCtl = &fakes.KappClient{}
			crtCtl = &fakes.CrtClient{}
			kappCtl.GetClientReturns(crtCtl)
			crtCtl.CreateReturnsOnCall(0, errors.New("failure in create Secret"))
		})
		It(testFailureMsg, func() {
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failure in create Secret"))
		})
		AfterEach(func() { options = opts })
	})

	Context("failure in creating SecretExport due to Create API error", func() {
		BeforeEach(func() {
			options.ExportToAllNamespaces = true
			kappCtl = &fakes.KappClient{}
			crtCtl = &fakes.CrtClient{}
			kappCtl.GetClientReturns(crtCtl)
			crtCtl.CreateReturnsOnCall(0, nil)
			crtCtl.CreateReturnsOnCall(1, errors.New("failure in create SecretExport"))
		})
		It(testFailureMsg, func() {
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failure in create SecretExport"))
		})
		AfterEach(func() { options = opts })
	})

	Context("success in creating Secret and SecretExport", func() {
		BeforeEach(func() {
			options.ExportToAllNamespaces = true
			kappCtl = &fakes.KappClient{}
			crtCtl = &fakes.CrtClient{}
			kappCtl.GetClientReturns(crtCtl)
			crtCtl.CreateReturnsOnCall(0, nil)
			crtCtl.CreateReturnsOnCall(1, nil)
		})
		It(testSuccessMsg, func() {
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() { options = opts })
	})
})