package apiserver

import (
	"context"

	"go.etcd.io/etcd/api/v3/etcdserverpb"
	mvccpb "go.etcd.io/etcd/api/v3/mvccpb"

	"openbce.io/kmds/pkg/storage"
)

// Range gets the keys in the range from the key-value store.
func (m *MdsBridge) Range(cxt context.Context, req *etcdserverpb.RangeRequest) (*etcdserverpb.RangeResponse, error) {
	// TODO (k82cn): handle key range

	records, err := m.storage.Range([][]byte{req.Key})
	if err != nil {
		return nil, err
	}

	kvs := make([]*mvccpb.KeyValue, 0, len(records))

	for _, r := range records {
		kvs = append(kvs, &mvccpb.KeyValue{
			Key:            r.Key,
			Value:          r.Value,
			Lease:          r.Lease,
			CreateRevision: r.Revision,
			ModRevision:    r.Revision,
		})
	}

	// return rangeResponse, nil
	return &etcdserverpb.RangeResponse{
		Header: &etcdserverpb.ResponseHeader{
			Revision: req.Revision,
		},
		More:  false,
		Count: 1,
		Kvs:   kvs,
	}, nil
}

// Put puts the given key into the key-value store.
// A put request increments the revision of the key-value store
// and generates one event in the event history.
func (m *MdsBridge) Put(cxt context.Context, req *etcdserverpb.PutRequest) (*etcdserverpb.PutResponse, error) {
	// TODO (k82cn): handle PreKV and IgnoreValue

	r := &storage.Record{
		Key:   req.Key,
		Value: req.Value,
		Lease: req.Lease,
	}

	r, err := m.storage.CreateOrUpdate(r)

	if err != nil {
		return nil, err
	}

	// TODO (k82cn): add an event here for Put api

	return &etcdserverpb.PutResponse{
		Header: &etcdserverpb.ResponseHeader{
			Revision: r.Revision,
		},
	}, nil
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
