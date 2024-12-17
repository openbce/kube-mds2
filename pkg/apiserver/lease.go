package apiserver

import (
	"context"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

var _ etcdserverpb.LeaseServer = (*MdsBridge)(nil)

// LeaseGrant creates a lease which expires if the server does not receive a keepAlive
// within a given time to live period. All keys attached to the lease will be expired and
// deleted if the lease expires. Each expired key generates a delete event in the event history.
func (m *MdsBridge) LeaseGrant(_ context.Context, req *etcdserverpb.LeaseGrantRequest) (*etcdserverpb.LeaseGrantResponse, error) {
	return &etcdserverpb.LeaseGrantResponse{
		Header: &etcdserverpb.ResponseHeader{},
		ID:     req.TTL,
		TTL:    req.TTL,
	}, nil
}

// LeaseRevoke revokes a lease. All keys attached to the lease will expire and be deleted.
func (m *MdsBridge) LeaseRevoke(context.Context, *etcdserverpb.LeaseRevokeRequest) (*etcdserverpb.LeaseRevokeResponse, error) {
	return nil, nil
}

// LeaseKeepAlive keeps the lease alive by streaming keep alive requests from the client
// to the server and streaming keep alive responses from the server to the client.
func (m *MdsBridge) LeaseKeepAlive(etcdserverpb.Lease_LeaseKeepAliveServer) error {
	return nil
}

// LeaseTimeToLive retrieves lease information.
func (m *MdsBridge) LeaseTimeToLive(context.Context, *etcdserverpb.LeaseTimeToLiveRequest) (*etcdserverpb.LeaseTimeToLiveResponse, error) {
	return nil, nil
}

// LeaseLeases lists all existing leases.
func (m *MdsBridge) LeaseLeases(context.Context, *etcdserverpb.LeaseLeasesRequest) (*etcdserverpb.LeaseLeasesResponse, error) {
	return nil, nil
}
