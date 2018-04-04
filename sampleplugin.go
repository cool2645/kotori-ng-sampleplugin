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
	// Here you load your plugin config file
	// A recommended place for your config file to locate is the conf.d folder
	_, err := toml.DecodeFile("conf.d/kotori-ng-sampleplugin.toml", &GlobCfg)
	return err
}

func (p *SamplePlugin) RegRouter(r *mux.Router) error {
	// Here you can configure your route
	// your "" will be /api/kotori-ng-sampleplugin (<-This is the plugin name)
	// while "/" will be /api/kotori-ng-sampleplugin/ (and so on)
	r.Methods("GET").Path("").HandlerFunc(handler.Pong)
	return nil
}

func (p *SamplePlugin) InitDB(db *gorm.DB) error {
	// Here you can do some migrations, such as
	// db.AutoMigrate(SomeModel{})
	return nil
}


var PluginInstance Plugin = &SamplePlugin{
	Name:    "Sample Plugin",
	Version: "0.0.1",
}
