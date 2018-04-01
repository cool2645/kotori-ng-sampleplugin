package main

import (
	. "github.com/cool2645/kotori-ng-sampleplugin/sampleplugin"
	. "github.com/cool2645/kotori-ng/kotoriplugin"
)

var PluginInstance Plugin = &SamplePlugin{
	Name:    "Sample Plugin",
	Version: "0.0.1",
}
