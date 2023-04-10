package main

import (
	"os"
	"tanzu-package-plugin-poc/pkg/package/flags"
	"tanzu-package-plugin-poc/pkg/package/kctrl"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/config"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
)

var descriptor = plugin.PluginDescriptor{
	Name:        "package",
	Description: "package (available, init, install, installed, release, repository)",
	Target:      types.TargetK8s, // <<<FIXME! set the Target of the plugin to one of {TargetGlobal,TargetK8s,TargetTMC}
	Version:     buildinfo.Version,
	BuildSHA:    buildinfo.SHA,
	Group:       plugin.RunCmdGroup, // set group
}

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "")
	}

	if config.IsFeatureActivated(flags.FeatureFlagPackagePluginKctrlCommandTree) {
		kctrl.Invoke(p)
		if err := p.Execute(); err != nil {
			os.Exit(1)
		}
		return
	}

	p.AddCommands(
	// Add commands
	)
	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
