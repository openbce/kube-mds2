package apiserver

import (
	"context"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

// MemberAdd adds a member into the cluster.
func (m *MdsBridge) MemberAdd(context.Context, *etcdserverpb.MemberAddRequest) (*etcdserverpb.MemberAddResponse, error) {
	return nil, nil
}

// MemberRemove removes an existing member from the cluster.
func (m *MdsBridge) MemberRemove(context.Context, *etcdserverpb.MemberRemoveRequest) (*etcdserverpb.MemberRemoveResponse, error) {
	return nil, nil
}

// MemberUpdate updates the member configuration.
func (m *MdsBridge) MemberUpdate(context.Context, *etcdserverpb.MemberUpdateRequest) (*etcdserverpb.MemberUpdateResponse, error) {
	return nil, nil
}

// MemberList lists all the members in the cluster.
func (m *MdsBridge) MemberList(context.Context, *etcdserverpb.MemberListRequest) (*etcdserverpb.MemberListResponse, error) {
	return nil, nil
}

// MemberPromote promotes a member from raft learner (non-voting) to raft voting member.
func (m *MdsBridge) MemberPromote(context.Context, *etcdserverpb.MemberPromoteRequest) (*etcdserverpb.MemberPromoteResponse, error) {
	return nil, nil
}
