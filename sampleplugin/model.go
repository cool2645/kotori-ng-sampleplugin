package sampleplugin

import (
	"github.com/cool2645/kotori-ng-sampleplugin/handler"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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

func (p *SamplePlugin) RegRouter(r *mux.Router) {
	r.Methods("GET").Path("/").HandlerFunc(handler.Pong)
}

func (p *SamplePlugin) InitDB(db *gorm.DB) {
	return
}
