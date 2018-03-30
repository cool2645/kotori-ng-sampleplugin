package main

import (
	. "github.com/cool2645/kotori-ng/kotori_plugin"
	. "github.com/cool2645/kotori-ng-sampleplugin/sampleplugin"
)

var PluginInstance Plugin = &SamplePlugin{
	Name:    "Sample Plugin",
	Version: "0.0.1",
}
