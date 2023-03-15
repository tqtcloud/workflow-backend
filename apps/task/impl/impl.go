package impl

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/workflow-backend/conf"
	"google.golang.org/grpc"

	"github.com/tqtcloud/workflow-backend/apps/task"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	//col *mongo.Collection
	conf *conf.Config
	log  logger.Logger
	task.UnimplementedServiceServer
}

func (s *service) Config() error {

	s.conf = conf.C()

	//s.col = db.Collection(s.Name())

	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Name() string {
	return task.AppName
}

func (s *service) Registry(server *grpc.Server) {
	task.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
