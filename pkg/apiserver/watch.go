package apiserver

import (
	"go.etcd.io/etcd/api/v3/etcdserverpb"
)

// Watch watches for events happening or that have happened. Both input and output
// are streams; the input stream is for creating and canceling watchers and the output
// stream sends events. One watch RPC can watch on multiple key ranges, streaming events
// for several watches at once. The entire event history can be watched starting from the
// last compaction revision.
func (m *MdsBridge) Watch(ws etcdserverpb.Watch_WatchServer) error {

	return nil
}
