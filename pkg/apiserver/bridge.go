package apiserver

import (
	"net"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"go.etcd.io/etcd/server/v3/embed"
	"k8s.io/klog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"openbce.io/kmds/pkg/storage"
)

// Metadata Server Bridge
type MdsBridge struct {
	storage  storage.Storage
	endpoint string
}

type MdsBridgeConfig struct {
	Endpoint string
	Engine   string
	Backend  string
}

func NewMdsBridage(config *MdsBridgeConfig) (*MdsBridge, error) {
	klog.V(2).Info(config)
	storage := storage.New(config.Engine, config.Backend)

	return &MdsBridge{
		storage:  storage,
		endpoint: config.Endpoint,
	}, nil
}

func (m *MdsBridge) Run() error {
	gopts := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             embed.DefaultGRPCKeepAliveMinTime,
			PermitWithoutStream: false,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    embed.DefaultGRPCKeepAliveInterval,
			Timeout: embed.DefaultGRPCKeepAliveTimeout,
		}),
	}

	// if config.ServerTLSConfig.CertFile != "" && config.ServerTLSConfig.KeyFile != "" {
	// 	creds, err := credentials.NewServerTLSFromFile(config.ServerTLSConfig.CertFile, config.ServerTLSConfig.KeyFile)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	gopts = append(gopts, grpc.Creds(creds))
	// }

	grpcServer := grpc.NewServer(gopts...)
	m.registerServers(grpcServer)

	listener, err := net.Listen("tcp", m.endpoint)
	if err != nil {
		return err
	}

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

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
