package handler

import (
	"net/http"
	. "github.com/cool2645/kotori-ng/httputils"
	. "github.com/cool2645/kotori-ng-sampleplugin/config"
)

func Pong(w http.ResponseWriter, req *http.Request) {
	res := map[string]interface{}{
		"code":   http.StatusOK,
		"result": true,
		"msg":    GlobCfg.Prompt,
	}
	ResponseJson(w, res, http.StatusOK)
	return
}
