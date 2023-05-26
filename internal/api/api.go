package api

import "github.com/toma-photo/internal/api/basic"

type ApiGroup struct {
	BasicApiGroup basic.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
