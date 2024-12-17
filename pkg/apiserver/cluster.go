package apiserver

import (
	"context"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

// TODO (k82cn): support member related api for loadbalance.

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
	return &etcdserverpb.MemberListResponse{
		Header: &etcdserverpb.ResponseHeader{},
		Members: []*etcdserverpb.Member{
			{
				Name:       "kube-mds2",
				ClientURLs: []string{"0.0.0.0:2379"},
				PeerURLs:   []string{"0.0.0.0:2379"},
			},
		},
	}, nil
}

// MemberPromote promotes a member from raft learner (non-voting) to raft voting member.
func (m *MdsBridge) MemberPromote(context.Context, *etcdserverpb.MemberPromoteRequest) (*etcdserverpb.MemberPromoteResponse, error) {
	return nil, nil
}
