package apiserver

import (
	"context"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
	"openbce.io/kmds/pkg/storage"
)

// Range gets the keys in the range from the key-value store.
func (m *MdsBridge) Range(context.Context, *etcdserverpb.RangeRequest) (*etcdserverpb.RangeResponse, error) {
	return nil, nil
}

// Put puts the given key into the key-value store.
// A put request increments the revision of the key-value store
// and generates one event in the event history.
func (m *MdsBridge) Put(cxt context.Context, req *etcdserverpb.PutRequest) (*etcdserverpb.PutResponse, error) {
	r := &storage.Record{
		Key:   req.Key,
		Value: req.Value,
	}

	_, err := m.storage.Create(r)

	resp := &etcdserverpb.PutResponse{}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteRange deletes the given range from the key-value store.
// A delete request increments the revision of the key-value store
// and generates a delete event in the event history for every deleted key.
func (m *MdsBridge) DeleteRange(context.Context, *etcdserverpb.DeleteRangeRequest) (*etcdserverpb.DeleteRangeResponse, error) {
	return nil, nil
}

// Txn processes multiple requests in a single transaction.
// A txn request increments the revision of the key-value store
// and generates events with the same revision for every completed request.
// It is not allowed to modify the same key several times within one txn.
func (m *MdsBridge) Txn(context.Context, *etcdserverpb.TxnRequest) (*etcdserverpb.TxnResponse, error) {
	return nil, nil
}

// Compact compacts the event history in the etcd key-value store. The key-value
// store should be periodically compacted or the event history will continue to grow
// indefinitely.
func (m *MdsBridge) Compact(context.Context, *etcdserverpb.CompactionRequest) (*etcdserverpb.CompactionResponse, error) {
	return nil, nil
}
