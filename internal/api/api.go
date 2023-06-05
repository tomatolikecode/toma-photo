package api

import (
	"github.com/toma-photo/internal/api/basic"
	"github.com/toma-photo/internal/api/system"
)

type ApiGroup struct {
	BasicApiGroup  basic.ApiGroup
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
