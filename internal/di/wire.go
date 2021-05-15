// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"git.code.oa.com/trpcprotocol/ctp_core/core_trpc_go/mykratos/internal/dao"
	"git.code.oa.com/trpcprotocol/ctp_core/core_trpc_go/mykratos/internal/service"
	"git.code.oa.com/trpcprotocol/ctp_core/core_trpc_go/mykratos/internal/server/grpc"
	"git.code.oa.com/trpcprotocol/ctp_core/core_trpc_go/mykratos/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
