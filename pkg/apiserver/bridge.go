package apiserver

import (
	"openbce.io/kmds/pkg/storage"

	"go.etcd.io/etcd/api/v3/etcdserverpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Metadata Server Bridge
type MdsBridge struct {
	storage storage.Storage
}

type MdsBridgeConfig struct {
	Endpoint string
	Engine   string
	Backend  string
}

func NewMdsBridage(config *MdsBridgeConfig) (*MdsBridge, error) {
	storage := storage.New(config.Endpoint, config.Backend)

	return &MdsBridge{
		storage: storage,
	}, nil
}

func (m *MdsBridge) Run() error {

	m.registerServers(nil)

	return nil
}

func (m *MdsBridge) registerServers(server *grpc.Server) {
	etcdserverpb.RegisterLeaseServer(server, m)
	etcdserverpb.RegisterWatchServer(server, m)
	etcdserverpb.RegisterKVServer(server, m)
	etcdserverpb.RegisterClusterServer(server, m)
	etcdserverpb.RegisterMaintenanceServer(server, m)

	hsrv := health.NewServer()
	hsrv.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(server, hsrv)

	reflection.Register(server)
}
