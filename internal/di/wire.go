// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/fordream2014/mykratos/internal/dao"
	"github.com/fordream2014/mykratos/internal/service"
	"github.com/fordream2014/mykratos/internal/server/grpc"
	"github.com/fordream2014/mykratos/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
