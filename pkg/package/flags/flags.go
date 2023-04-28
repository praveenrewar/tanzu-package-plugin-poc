// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package flags

import "github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"

var PersistentFlagsDefault = PersistentFlags{}

type PersistentFlags struct {
	Kubeconfig string
	LogLevel   int32
}

func (f *PersistentFlags) Set(p *plugin.Plugin) {
	p.Cmd.PersistentFlags().Int32VarP(&f.LogLevel, "verbose", "", 0, "Number for the log level verbosity(0-9)")
	p.Cmd.PersistentFlags().StringVarP(&f.Kubeconfig, "kubeconfig", "", "", "The path to the kubeconfig file, optional")
}
