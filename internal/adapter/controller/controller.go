package controller

import "github.com/otakakot/2023-golang-project-layout/pkg/api"

var _ api.Handler = (*Controller)(nil)

type Controller struct {
	*Todo
}
