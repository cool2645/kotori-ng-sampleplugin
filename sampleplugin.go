package main

import (
	"github.com/cool2645/kotori-ng-sampleplugin/handler"
	. "github.com/cool2645/kotori-ng/kotoriplugin"
	. "github.com/cool2645/kotori-ng-sampleplugin/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/BurntSushi/toml"
)

type SamplePlugin struct {
	Name    string
	Version string
}

func (p *SamplePlugin) GetName() string {
	return p.Name
}

func (p *SamplePlugin) GetVersion() string {
	return p.Version
}

func (p *SamplePlugin) LoadConfig() error  {
	_, err := toml.DecodeFile("conf.d/kotori-ng-sampleplugin.toml", &GlobCfg)
	return err
}

func (p *SamplePlugin) RegRouter(r *mux.Router) error {
	r.Methods("GET").Path("").HandlerFunc(handler.Pong)
	return nil
}

func (p *SamplePlugin) InitDB(db *gorm.DB) error {
	return nil
}


var PluginInstance Plugin = &SamplePlugin{
	Name:    "Sample Plugin",
	Version: "0.0.1",
}
