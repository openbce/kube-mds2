package apiserver

import (
	"context"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

// Alarm activates, deactivates, and queries alarms regarding cluster health.
func (m *MdsBridge) Alarm(context.Context, *etcdserverpb.AlarmRequest) (*etcdserverpb.AlarmResponse, error) {
	return nil, nil
}

// Status gets the status of the member.
func (m *MdsBridge) Status(cxt context.Context, req *etcdserverpb.StatusRequest) (*etcdserverpb.StatusResponse, error) {
	size, err := m.storage.Size()
	if err != nil {
		return nil, err
	}
	return &etcdserverpb.StatusResponse{
		Header:  &etcdserverpb.ResponseHeader{},
		DbSize:  size,
		Version: "v3",
	}, nil
}

// Defragment defragments a member's backend database to recover storage space.
func (m *MdsBridge) Defragment(context.Context, *etcdserverpb.DefragmentRequest) (*etcdserverpb.DefragmentResponse, error) {
	return nil, nil
}

// Hash computes the hash of whole backend keyspace,
// including key, lease, and other buckets in storage.
// This is designed for testing ONLY!
// Do not rely on this in production with ongoing transactions,
// since Hash operation does not hold MVCC locks.
// Use "HashKV" API instead for "key" bucket consistency checks.
func (m *MdsBridge) Hash(context.Context, *etcdserverpb.HashRequest) (*etcdserverpb.HashResponse, error) {
	return nil, nil
}

// HashKV computes the hash of all MVCC keys up to a given revision.
// It only iterates "key" bucket in backend storage.
func (m *MdsBridge) HashKV(context.Context, *etcdserverpb.HashKVRequest) (*etcdserverpb.HashKVResponse, error) {
	return nil, nil
}

// Snapshot sends a snapshot of the entire backend from a member over a stream to a client.
func (m *MdsBridge) Snapshot(*etcdserverpb.SnapshotRequest, etcdserverpb.Maintenance_SnapshotServer) error {
	return nil
}

// MoveLeader requests current leader node to transfer its leadership to transferee.
func (m *MdsBridge) MoveLeader(context.Context, *etcdserverpb.MoveLeaderRequest) (*etcdserverpb.MoveLeaderResponse, error) {
	return nil, nil
}

// Downgrade requests downgrades, verifies feasibility or cancels downgrade
// on the cluster version.
// Supported since etcd 3.5.
func (m *MdsBridge) Downgrade(context.Context, *etcdserverpb.DowngradeRequest) (*etcdserverpb.DowngradeResponse, error) {
	return nil, nil
}
